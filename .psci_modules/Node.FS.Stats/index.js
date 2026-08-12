import * as $foreign from "./foreign.js";
import * as Data_JSDate from "../Data.JSDate/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Partial_Unsafe from "../Partial.Unsafe/index.js";

// | The numeric user identifier of the user that owns the file (POSIX).
var uid = function (s) {
    return $foreign.uidImpl(s);
};
var statusChangedTimeMs = function (s) {
    return $foreign.statusChangedTimeMsImpl(s);
};
var statusChangedTime = function (s) {
    var v = Data_JSDate.toDateTime($foreign.statusChangedTimeImpl(s));
    if (v instanceof Data_Maybe.Just) {
        return v.value0;
    };
    if (v instanceof Data_Maybe.Nothing) {
        return Partial_Unsafe.unsafeCrashWith("Impossible: `statusChangedTime` returned invalid DateTime value.");
    };
    throw new Error("Failed pattern match at Node.FS.Stats (line 181, column 23 - line 183, column 98): " + [ v.constructor.name ]);
};

// | The size of the file in bytes.
// | If the underlying file system does not support getting the size of the file, this will be 0.
var size = function (s) {
    return $foreign.sizeImpl(s);
};
var showStats = {
    show: function (s) {
        return "Stats " + $foreign.showStatsObj(s);
    }
};

// | A numeric device identifier if the file represents a device.
var rdev = function (s) {
    return $foreign.rdevImpl(s);
};

// | The number of hard-links that exist for the file.
var nlink = function (s) {
    return $foreign.nlinkImpl(s);
};
var modifiedTimeMs = function (s) {
    return $foreign.modifiedTimeMsImpl(s);
};
var modifiedTime = function (s) {
    var v = Data_JSDate.toDateTime($foreign.modifiedTimeImpl(s));
    if (v instanceof Data_Maybe.Just) {
        return v.value0;
    };
    if (v instanceof Data_Maybe.Nothing) {
        return Partial_Unsafe.unsafeCrashWith("Impossible: `modifiedTime` returned invalid DateTime value.");
    };
    throw new Error("Failed pattern match at Node.FS.Stats (line 174, column 18 - line 176, column 93): " + [ v.constructor.name ]);
};

// | A bit-field describing the file type and mode.
var mode = function (s) {
    return $foreign.modeImpl(s);
};
var isSymbolicLink = function (s) {
    return $foreign.isSymbolicLinkImpl(s);
};
var isSocket = function (s) {
    return $foreign.isSocketImpl(s);
};
var isFile = function (s) {
    return $foreign.isFileImpl(s);
};
var isFIFO = function (s) {
    return $foreign.isFIFOImpl(s);
};

// | Returns true if the <fs.Stats> object describes a file system directory.
// | If the `fs.Stats`> object was obtained from `fs.lstat()`, 
// | this method will always return `false``. This is because `fs.lstat()` 
// | returns information about a symbolic link itself and not the path to which it resolves.
var isDirectory = function (s) {
    return $foreign.isDirectoryImpl(s);
};
var isCharacterDevice = function (s) {
    return $foreign.isCharacterDeviceImpl(s);
};
var isBlockDevice = function (s) {
    return $foreign.isBlockDeviceImpl(s);
};

// | The file system specific "Inode" number for the file.
var inode = function (s) {
    return $foreign.inodeImpl(s);
};

// | The numeric group identifier of the group that owns the file (POSIX).
var gid = function (s) {
    return $foreign.gidImpl(s);
};

// | The numeric identifier of the device containing the file.
var dev = function (s) {
    return $foreign.devImpl(s);
};

// | The number of blocks allocated for this file.
var blocks = function (s) {
    return $foreign.blocksImpl(s);
};

// | The file system block size for i/o operations.
var blkSize = function (s) {
    return $foreign.blkSizeImpl(s);
};
var birthtimeMs = function (s) {
    return $foreign.birthtimeMsImpl(s);
};
var birthTime = function (s) {
    var v = Data_JSDate.toDateTime($foreign.birthTimeImpl(s));
    if (v instanceof Data_Maybe.Just) {
        return v.value0;
    };
    if (v instanceof Data_Maybe.Nothing) {
        return Partial_Unsafe.unsafeCrashWith("Impossible: `birthTime` returned invalid DateTime value.");
    };
    throw new Error("Failed pattern match at Node.FS.Stats (line 188, column 15 - line 190, column 90): " + [ v.constructor.name ]);
};
var accessedTimeMs = function (s) {
    return $foreign.accessedTimeMsImpl(s);
};
var accessedTime = function (s) {
    var v = Data_JSDate.toDateTime($foreign.accessedTimeImpl(s));
    if (v instanceof Data_Maybe.Just) {
        return v.value0;
    };
    if (v instanceof Data_Maybe.Nothing) {
        return Partial_Unsafe.unsafeCrashWith("Impossible: `accessedTime` returned invalid DateTime value.");
    };
    throw new Error("Failed pattern match at Node.FS.Stats (line 167, column 18 - line 169, column 93): " + [ v.constructor.name ]);
};
export {
    isFile,
    isDirectory,
    isBlockDevice,
    isCharacterDevice,
    isFIFO,
    isSocket,
    isSymbolicLink,
    dev,
    inode,
    mode,
    nlink,
    uid,
    gid,
    rdev,
    size,
    blkSize,
    blocks,
    accessedTimeMs,
    modifiedTimeMs,
    statusChangedTimeMs,
    birthtimeMs,
    accessedTime,
    modifiedTime,
    statusChangedTime,
    birthTime,
    showStats
};
