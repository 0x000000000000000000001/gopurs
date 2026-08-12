import * as $foreign from "./foreign.js";
import * as Data_DateTime_Instant from "../Data.DateTime.Instant/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_EuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data_Function from "../Data.Function/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Nullable from "../Data.Nullable/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Effect from "../Effect/index.js";
import * as Effect_Exception from "../Effect.Exception/index.js";
import * as Node_Buffer from "../Node.Buffer/index.js";
import * as Node_Encoding from "../Node.Encoding/index.js";
import * as Node_FS from "../Node.FS/index.js";
import * as Node_FS_Constants from "../Node.FS.Constants/index.js";
import * as Node_FS_Perms from "../Node.FS.Perms/index.js";

// | Writes text to a file using the specified encoding.
var writeTextFile = function (encoding) {
    return function (file) {
        return function (text) {
            return function () {
                return $foreign.writeFileSyncImpl(file, text, {
                    encoding: Data_Show.show(Node_Encoding.showEncoding)(encoding)
                });
            };
        };
    };
};

// | Writes a buffer to a file.
var writeFile = function (file) {
    return function (buff) {
        return function () {
            return $foreign.writeFileSyncImpl(file, buff, {});
        };
    };
};

// | Sets the accessed and modified times for the specified file.
var utimes = function (file) {
    return function (atime) {
        return function (mtime) {
            var toEpochMilliseconds = function ($7) {
                return Data_DateTime_Instant.unInstant(Data_DateTime_Instant.fromDateTime($7));
            };
            var ms = function (v) {
                return Data_Int.round(v);
            };
            var fromDate = function (date) {
                return Data_EuclideanRing.div(Data_EuclideanRing.euclideanRingInt)(ms(toEpochMilliseconds(date)))(1000);
            };
            return function () {
                return $foreign.utimesSyncImpl(file, fromDate(atime), fromDate(mtime));
            };
        };
    };
};

// | Deletes a file.
var unlink = function (file) {
    return function () {
        return $foreign.unlinkSyncImpl(file);
    };
};

// | Truncates a file to the specified length.
var truncate = function (file) {
    return function (len) {
        return function () {
            return $foreign.truncateSyncImpl(file, len);
        };
    };
};

// | Creates a symlink.
var symlink = function (src) {
    return function (dst) {
        return function (ty) {
            return function () {
                return $foreign.symlinkSyncImpl(src, dst, Node_FS.symlinkTypeToNode(ty));
            };
        };
    };
};

// | Gets file statistics.
var stat = function (file) {
    return function () {
        return $foreign.statSyncImpl(file);
    };
};

// | Deletes a directory with options.
var rmdir$prime = function (path) {
    return function (opts) {
        return function () {
            return $foreign.rmdirSyncImpl(path, opts);
        };
    };
};

// | Deletes a directory.
var rmdir = function (path) {
    return rmdir$prime(path)({
        maxRetries: 0,
        retryDelay: 100
    });
};

// | Deletes a file or directory with options.
var rm$prime = function (path) {
    return function (opts) {
        return function () {
            return $foreign.rmSyncImpl(path, opts);
        };
    };
};

// | Deletes a file or directory.
var rm = function (path) {
    return rm$prime(path)({
        force: false,
        maxRetries: 100,
        recursive: false,
        retryDelay: 1000
    });
};

// | Renames a file.
var rename = function (oldFile) {
    return function (newFile) {
        return function () {
            return $foreign.renameSyncImpl(oldFile, newFile);
        };
    };
};

// | Find the canonicalized absolute location for a path using a cache object for
// | already resolved paths.
var realpath$prime = function (path) {
    return function (cache) {
        return function () {
            return $foreign.realpathSyncImpl(path, cache);
        };
    };
};

// | Find the canonicalized absolute location for a path.
var realpath = function (path) {
    return function () {
        return $foreign.realpathSyncImpl(path, {});
    };
};

// | Reads the value of a symlink.
var readlink = function (path) {
    return function () {
        return $foreign.readlinkSyncImpl(path);
    };
};

// | Reads the contents of a directory.
var readdir = function (file) {
    return function () {
        return $foreign.readdirSyncImpl(file);
    };
};

// | Reads the entire contents of a text file with the specified encoding.
var readTextFile = function (encoding) {
    return function (file) {
        return function () {
            return $foreign.readFileSyncImpl(file, {
                encoding: Data_Show.show(Node_Encoding.showEncoding)(encoding)
            });
        };
    };
};

// | Reads the entire contents of a file returning the result as a raw buffer.
var readFile = function (file) {
    return function () {
        return $foreign.readFileSyncImpl(file, {});
    };
};
var mkdtemp$prime = function (prefix) {
    return function (encoding) {
        return function () {
            return $foreign.mkdtempImpl(prefix, Node_Encoding.encodingToNode(encoding));
        };
    };
};
var mkdtemp = function (prefix) {
    return mkdtemp$prime(prefix)(Node_Encoding.UTF8.value);
};

// | Makes a new directory with the specified permissions.
var mkdir$prime = function (file) {
    return function (v) {
        return function () {
            return $foreign.mkdirSyncImpl(file, {
                recursive: v.recursive,
                mode: Node_FS_Perms.permsToString(v.mode)
            });
        };
    };
};

// | Makes a new directory.
var mkdir = function (path) {
    return mkdir$prime(path)({
        recursive: false,
        mode: Node_FS_Perms.mkPerms(Node_FS_Perms.all)(Node_FS_Perms.all)(Node_FS_Perms.all)
    });
};

// | Gets file or symlink statistics. `lstat` is identical to `stat`, except
// | that if the `FilePath` is a symbolic link, then the link itself is stat-ed,
// | not the file that it refers to.
var lstat = function (file) {
    return function () {
        return $foreign.lstatSyncImpl(file);
    };
};

// | Creates a link to an existing file.
var link = function (src) {
    return function (dst) {
        return function () {
            return $foreign.linkSyncImpl(src, dst);
        };
    };
};

// | Write to a file synchronously. See the [Node documentation](http://nodejs.org/api/fs.html#fs_fs_writesync_fd_buffer_offset_length_position)
// | for details.
var fdWrite = function (fd) {
    return function (buff) {
        return function (off) {
            return function (len) {
                return function (pos) {
                    return function () {
                        return $foreign.writeSyncImpl(fd, buff, off, len, Data_Nullable.toNullable(pos));
                    };
                };
            };
        };
    };
};

// | Read from a file synchronously. See the [Node documentation](http://nodejs.org/api/fs.html#fs_fs_readsync_fd_buffer_offset_length_position)
// | for details.
var fdRead = function (fd) {
    return function (buff) {
        return function (off) {
            return function (len) {
                return function (pos) {
                    return function () {
                        return $foreign.readSyncImpl(fd, buff, off, len, Data_Nullable.toNullable(pos));
                    };
                };
            };
        };
    };
};

// | Open a file synchronously. See the [Node documentation](http://nodejs.org/api/fs.html#fs_fs_opensync_path_flags_mode)
// | for details.
var fdOpen = function (file) {
    return function (flags) {
        return function (mode) {
            return function () {
                return $foreign.openSyncImpl(file, Node_FS_Constants.fileFlagsToNode(flags), Data_Nullable.toNullable(mode));
            };
        };
    };
};

// | Convenience function to fill the whole buffer from the current
// | file position.
var fdNext = function (fd) {
    return function (buff) {
        return function __do() {
            var sz = Node_Buffer.size(buff)();
            return fdRead(fd)(buff)(0)(sz)(Data_Maybe.Nothing.value)();
        };
    };
};

// | Flush a file synchronously.  See the [Node documentation](http://nodejs.org/api/fs.html#fs_fs_fsyncsync_fd)
// | for details.
var fdFlush = function (fd) {
    return function () {
        return $foreign.fsyncSyncImpl(fd);
    };
};

// | Close a file synchronously. See the [Node documentation](http://nodejs.org/api/fs.html#fs_fs_closesync_fd)
// | for details.
var fdClose = function (fd) {
    return function () {
        return $foreign.closeSyncImpl(fd);
    };
};

// | Convenience function to append the whole buffer to the current
// | file position.
var fdAppend = function (fd) {
    return function (buff) {
        return function __do() {
            var sz = Node_Buffer.size(buff)();
            return fdWrite(fd)(buff)(0)(sz)(Data_Maybe.Nothing.value)();
        };
    };
};

// | Check if the path exists.
var exists = function (file) {
    return function () {
        return $foreign.existsSyncImpl(file);
    };
};
var copyFile$prime = function (src) {
    return function (dest) {
        return function (mode) {
            return function () {
                return $foreign.copyFileImpl(src, dest, mode);
            };
        };
    };
};
var copyFile = function (src) {
    return function (dest) {
        return function () {
            return $foreign.copyFileImpl(src, dest, Node_FS_Constants.defaultCopyMode);
        };
    };
};

// | Changes the ownership of a file.
var chown = function (file) {
    return function (uid) {
        return function (gid) {
            return function () {
                return $foreign.chownSyncImpl(file, uid, gid);
            };
        };
    };
};

// | Changes the permissions of a file.
var chmod = function (file) {
    return function (perms) {
        return function () {
            return $foreign.chmodSyncImpl(file, Node_FS_Perms.permsToString(perms));
        };
    };
};

// | Appends text to a file using the specified encoding.
var appendTextFile = function (encoding) {
    return function (file) {
        return function (buff) {
            return function () {
                return $foreign.appendFileSyncImpl(file, buff, {
                    encoding: Data_Show.show(Node_Encoding.showEncoding)(encoding)
                });
            };
        };
    };
};

// | Appends the contents of a buffer to a file.
var appendFile = function (file) {
    return function (buff) {
        return function () {
            return $foreign.appendFileSyncImpl(file, buff, {});
        };
    };
};
var access$prime = function (path) {
    return function (mode) {
        return Data_Functor.map(Effect.functorEffect)(Data_Either.blush)(Effect_Exception["try"](function () {
            return $foreign.accessImpl(path, mode);
        }));
    };
};
var access = /* #__PURE__ */ Data_Function.flip(access$prime)(Node_FS_Constants.defaultAccessMode);
export {
    access,
    access$prime,
    copyFile,
    copyFile$prime,
    mkdtemp,
    mkdtemp$prime,
    rename,
    truncate,
    chown,
    chmod,
    stat,
    lstat,
    link,
    symlink,
    readlink,
    realpath,
    realpath$prime,
    unlink,
    rmdir,
    rmdir$prime,
    rm,
    rm$prime,
    mkdir,
    mkdir$prime,
    readdir,
    utimes,
    readFile,
    readTextFile,
    writeFile,
    writeTextFile,
    appendFile,
    appendTextFile,
    exists,
    fdOpen,
    fdRead,
    fdNext,
    fdWrite,
    fdAppend,
    fdFlush,
    fdClose
};
