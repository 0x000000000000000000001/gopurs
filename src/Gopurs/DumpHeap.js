import v8 from 'node:v8';
import fs from 'node:fs';

export const dumpHeapIfLarge = function() {
  const rss = process.memoryUsage().rss / 1024 / 1024;
  if (rss > 8000 && !globalThis.heapDumped) {
    globalThis.heapDumped = true;
    console.log("Memory exceeded 8GB, dumping heap...");
    const filename = `heapdump-${Date.now()}.heapsnapshot`;
    v8.writeHeapSnapshot(filename);
    console.log("Heap dumped to " + filename);
  }
};
