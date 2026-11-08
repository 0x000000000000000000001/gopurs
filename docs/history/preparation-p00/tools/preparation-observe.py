#!/usr/bin/env python3
"""Record resource contention around one private benchmark command."""
import datetime
import json
import os
from pathlib import Path
import signal
import subprocess
import sys
import threading
import time


def processes():
    result = subprocess.run(["ps", "-Ao", "pid,ppid,%cpu,rss,comm"],
                            capture_output=True, text=True, check=True)
    return [dict(pid=int(p[0]), parent=int(p[1]), cpu=float(p[2]),
                 rssBytes=int(p[3]) * 1024, command=p[4])
            for line in result.stdout.splitlines()[1:]
            if len(p := line.strip().split(None, 4)) == 5]


def owned(rows, current):
    parents = {row["pid"]: row["parent"] for row in rows}
    excluded = {current}
    parent = parents.get(current)
    while parent and parent not in excluded:
        excluded.add(parent)
        parent = parents.get(parent)
    descendants = {current}
    while True:
        children = {pid for pid, parent in parents.items() if parent in descendants}
        if children <= descendants:
            return excluded | descendants
        descendants |= children


def competing(rows, own):
    compilers = {"go", "compile", "link", "purs", "gopurs-native", "gopurs-rust",
                 "ghc", "cabal", "stack", "cargo", "rustc", "node",
                 "go-packages-helper", "build-cache-helper"}
    return [row for row in rows if row["pid"] not in own
            and Path(row["command"]).name in compilers
            and (row["cpu"] >= 20 or row["rssBytes"] >= 1024**3)]


def main():
    if sys.argv[1:] == ["--self-check"]:
        rows = [dict(pid=1, parent=0), dict(pid=2, parent=1),
                dict(pid=3, parent=2), dict(pid=4, parent=3), dict(pid=5, parent=2)]
        assert owned(rows, 3) == {1, 2, 3, 4}
        assert competing([dict(pid=5, command="/private/gopurs-native", cpu=400,
                               rssBytes=10 * 1024**3)], {1, 2, 3, 4})
        print("Resource attribution self-check passed")
        return 0
    assert len(sys.argv) >= 5 and sys.argv[2] == "--"
    destination, command = Path(sys.argv[1]), sys.argv[3:]
    assert not destination.exists(), "Preserve prior resource traces"
    destination.parent.mkdir(parents=True, exist_ok=True)
    stop = threading.Event()
    contaminated = []
    failures = []
    start = time.monotonic()
    with destination.open("x") as log:
        def sample():
            rows = processes()
            other = competing(rows, owned(rows, os.getpid()))
            contaminated.extend(other)
            record = dict(time=datetime.datetime.now().astimezone().isoformat(),
                          elapsed=time.monotonic() - start, competing=other,
                          busiest=sorted(rows, key=lambda r: r["cpu"], reverse=True)[:12])
            log.write(json.dumps(record) + "\n")
            log.flush()

        sample()
        if contaminated:
            print("P00: another heavy compiler is active; command deferred", file=sys.stderr)
            return 75
        child = subprocess.Popen(command, start_new_session=True)

        def forward(number, _frame):
            try:
                os.killpg(child.pid, number)
            except ProcessLookupError:
                pass

        for number in [signal.SIGINT, signal.SIGTERM, signal.SIGHUP]:
            signal.signal(number, forward)

        def monitor():
            while not stop.wait(2):
                try:
                    sample()
                except Exception as failure:
                    failures.append(str(failure))
                    break

        thread = threading.Thread(target=monitor)
        thread.start()
        status = child.wait()
        stop.set()
        thread.join()
        sample()
        log.write(json.dumps(dict(completed=True, status=status,
                                  contaminated=bool(contaminated), errors=failures)) + "\n")
        if contaminated or failures:
            print("P00: resource trace invalidates this timing; see " + str(destination), file=sys.stderr)
            return 75
        return status if status >= 0 else 128 - status


if __name__ == "__main__":
    sys.exit(main())
