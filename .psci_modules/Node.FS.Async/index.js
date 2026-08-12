import * as $foreign from "./foreign.js";
import * as Data_DateTime_Instant from "../Data.DateTime.Instant/index.js";
import * as Data_Either from "../Data.Either/index.js";
import * as Data_EuclideanRing from "../Data.EuclideanRing/index.js";
import * as Data_Int from "../Data.Int/index.js";
import * as Data_Maybe from "../Data.Maybe/index.js";
import * as Data_Nullable from "../Data.Nullable/index.js";
import * as Data_Show from "../Data.Show/index.js";
import * as Node_Buffer from "../Node.Buffer/index.js";
import * as Node_Encoding from "../Node.Encoding/index.js";
import * as Node_FS from "../Node.FS/index.js";
import * as Node_FS_Constants from "../Node.FS.Constants/index.js";
import * as Node_FS_Perms from "../Node.FS.Perms/index.js";
var handleCallback = function (cb) {
    return function (err, a) {
        var v = Data_Nullable.toMaybe(err);
        if (v instanceof Data_Maybe.Nothing) {
            return cb(new Data_Either.Right(a))();
        };
        if (v instanceof Data_Maybe.Just) {
            return cb(new Data_Either.Left(v.value0))();
        };
        throw new Error("Failed pattern match at Node.FS.Async (line 66, column 43 - line 68, column 30): " + [ v.constructor.name ]);
    };
};

// | Creates a link to an existing file.
var link = function (src) {
    return function (dst) {
        return function (cb) {
            return function () {
                return $foreign.linkImpl(src, dst, handleCallback(cb));
            };
        };
    };
};

// | Gets file or symlink statistics. `lstat` is identical to `stat`, except
// | that if the `FilePath` is a symbolic link, then the link itself is stat-ed,
// | not the file that it refers to.
var lstat = function (file) {
    return function (cb) {
        return function () {
            return $foreign.lstatImpl(file, handleCallback(cb));
        };
    };
};

// | Makes a new directory with the specified permissions.
var mkdir$prime = function (file) {
    return function (v) {
        return function (cb) {
            return function () {
                return $foreign.mkdirImpl(file, {
                    recursive: v.recursive,
                    mode: Node_FS_Perms.permsToString(v.mode)
                }, handleCallback(cb));
            };
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
var mkdtemp$prime = function (prefix) {
    return function (encoding) {
        return function (cb) {
            return function () {
                return $foreign.mkdtempImpl(prefix, Node_Encoding.encodingToNode(encoding), handleCallback(cb));
            };
        };
    };
};
var mkdtemp = function (prefix) {
    return mkdtemp$prime(prefix)(Node_Encoding.UTF8.value);
};

// | Reads the entire contents of a file returning the result as a raw buffer.
var readFile = function (file) {
    return function (cb) {
        return function () {
            return $foreign.readFileImpl(file, {}, handleCallback(cb));
        };
    };
};

// | Reads the entire contents of a text file with the specified encoding.
var readTextFile = function (encoding) {
    return function (file) {
        return function (cb) {
            return function () {
                return $foreign.readFileImpl(file, {
                    encoding: Data_Show.show(Node_Encoding.showEncoding)(encoding)
                }, handleCallback(cb));
            };
        };
    };
};

// | Reads the contents of a directory.
var readdir = function (file) {
    return function (cb) {
        return function () {
            return $foreign.readdirImpl(file, handleCallback(cb));
        };
    };
};

// | Reads the value of a symlink.
var readlink = function (path) {
    return function (cb) {
        return function () {
            return $foreign.readlinkImpl(path, handleCallback(cb));
        };
    };
};

// | Find the canonicalized absolute location for a path.
var realpath = function (path) {
    return function (cb) {
        return function () {
            return $foreign.realpathImpl(path, {}, handleCallback(cb));
        };
    };
};

// | Find the canonicalized absolute location for a path using a cache object
// | for already resolved paths.
var realpath$prime = function (path) {
    return function (cache) {
        return function (cb) {
            return function () {
                return $foreign.realpathImpl(path, cache, handleCallback(cb));
            };
        };
    };
};

// | Renames a file.
var rename = function (oldFile) {
    return function (newFile) {
        return function (cb) {
            return function () {
                return $foreign.renameImpl(oldFile, newFile, handleCallback(cb));
            };
        };
    };
};

// | Deletes a file or directory with options.
var rm$prime = function (path) {
    return function (opts) {
        return function (cb) {
            return function () {
                return $foreign.rmImpl(path, opts, handleCallback(cb));
            };
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

// | Deletes a directory with options.
var rmdir$prime = function (path) {
    return function (opts) {
        return function (cb) {
            return function () {
                return $foreign.rmdirImpl(path, opts, handleCallback(cb));
            };
        };
    };
};

// | Deletes a directory.
var rmdir = function (path) {
    return function (cb) {
        return rmdir$prime(path)({
            maxRetries: 0,
            retryDelay: 100
        })(cb);
    };
};

// | Gets file statistics.
var stat = function (file) {
    return function (cb) {
        return function () {
            return $foreign.statImpl(file, handleCallback(cb));
        };
    };
};

// | Creates a symlink.
var symlink = function (src) {
    return function (dest) {
        return function (ty) {
            return function (cb) {
                return function () {
                    return $foreign.symlinkImpl(src, dest, Node_FS.symlinkTypeToNode(ty), handleCallback(cb));
                };
            };
        };
    };
};

// | Truncates a file to the specified length.
var truncate = function (file) {
    return function (len) {
        return function (cb) {
            return function () {
                return $foreign.truncateImpl(file, len, handleCallback(cb));
            };
        };
    };
};

// | Deletes a file.
var unlink = function (file) {
    return function (cb) {
        return function () {
            return $foreign.unlinkImpl(file, handleCallback(cb));
        };
    };
};

// | Sets the accessed and modified times for the specified file.
var utimes = function (file) {
    return function (atime) {
        return function (mtime) {
            return function (cb) {
                var toEpochMilliseconds = function ($11) {
                    return Data_DateTime_Instant.unInstant(Data_DateTime_Instant.fromDateTime($11));
                };
                var ms = function (v) {
                    return Data_Int.round(v);
                };
                var fromDate = function (date) {
                    return Data_EuclideanRing.div(Data_EuclideanRing.euclideanRingInt)(ms(toEpochMilliseconds(date)))(1000);
                };
                return function () {
                    return $foreign.utimesImpl(file, fromDate(atime), fromDate(mtime), handleCallback(cb));
                };
            };
        };
    };
};

// | Writes a buffer to a file.
var writeFile = function (file) {
    return function (buff) {
        return function (cb) {
            return function () {
                return $foreign.writeFileImpl(file, buff, {}, handleCallback(cb));
            };
        };
    };
};

// | Writes text to a file using the specified encoding.
var writeTextFile = function (encoding) {
    return function (file) {
        return function (buff) {
            return function (cb) {
                return function () {
                    return $foreign.writeFileImpl(file, buff, {
                        encoding: Data_Show.show(Node_Encoding.showEncoding)(encoding)
                    }, handleCallback(cb));
                };
            };
        };
    };
};

// | Write to a file asynchronously. See the [Node Documentation](https://nodejs.org/api/fs.html#fs_fs_write_fd_buffer_offset_length_position_callback)
// | for details.
var fdWrite = function (fd) {
    return function (buff) {
        return function (off) {
            return function (len) {
                return function (pos) {
                    return function (cb) {
                        return function () {
                            return $foreign.writeImpl(fd, buff, off, len, Data_Nullable.toNullable(pos), handleCallback(cb));
                        };
                    };
                };
            };
        };
    };
};

// | Read from a file asynchronously. See the [Node Documentation](https://nodejs.org/api/fs.html#fs_fs_read_fd_buffer_offset_length_position_callback)
// | for details.
var fdRead = function (fd) {
    return function (buff) {
        return function (off) {
            return function (len) {
                return function (pos) {
                    return function (cb) {
                        return function () {
                            return $foreign.readImpl(fd, buff, off, len, Data_Nullable.toNullable(pos), handleCallback(cb));
                        };
                    };
                };
            };
        };
    };
};

// | Open a file asynchronously. See the [Node Documentation](https://nodejs.org/api/fs.html#fs_fs_open_path_flags_mode_callback)
// | for details.
var fdOpen = function (file) {
    return function (flags) {
        return function (mode) {
            return function (cb) {
                return function () {
                    return $foreign.openImpl(file, Node_FS_Constants.fileFlagsToNode(flags), Data_Nullable.toNullable(mode), handleCallback(cb));
                };
            };
        };
    };
};

// | Convenience function to fill the whole buffer from the current
// | file position.
var fdNext = function (fd) {
    return function (buff) {
        return function (cb) {
            return function __do() {
                var sz = Node_Buffer.size(buff)();
                return fdRead(fd)(buff)(0)(sz)(Data_Maybe.Nothing.value)(cb)();
            };
        };
    };
};

// | Close a file asynchronously. See the [Node Documentation](https://nodejs.org/api/fs.html#fs_fs_close_fd_callback)
// | for details.
var fdClose = function (fd) {
    return function (cb) {
        return function () {
            return $foreign.closeImpl(fd, handleCallback(cb));
        };
    };
};

// | Convenience function to append the whole buffer to the current
// | file position.
var fdAppend = function (fd) {
    return function (buff) {
        return function (cb) {
            return function __do() {
                var sz = Node_Buffer.size(buff)();
                return fdWrite(fd)(buff)(0)(sz)(Data_Maybe.Nothing.value)(cb)();
            };
        };
    };
};
var copyFile$prime = function (src) {
    return function (dest) {
        return function (mode) {
            return function (cb) {
                return function () {
                    return $foreign.copyFileImpl(src, dest, mode, handleCallback(cb));
                };
            };
        };
    };
};
var copyFile = function (src) {
    return function (dest) {
        return copyFile$prime(src)(dest)(Node_FS_Constants.defaultCopyMode);
    };
};

// | Changes the ownership of a file.
var chown = function (file) {
    return function (uid) {
        return function (gid) {
            return function (cb) {
                return function () {
                    return $foreign.chownImpl(file, uid, gid, handleCallback(cb));
                };
            };
        };
    };
};

// | Changes the permissions of a file.
var chmod = function (file) {
    return function (perms) {
        return function (cb) {
            return function () {
                return $foreign.chmodImpl(file, Node_FS_Perms.permsToString(perms), handleCallback(cb));
            };
        };
    };
};

// | Appends text to a file using the specified encoding.
var appendTextFile = function (encoding) {
    return function (file) {
        return function (buff) {
            return function (cb) {
                return function () {
                    return $foreign.appendFileImpl(file, buff, {
                        encoding: Data_Show.show(Node_Encoding.showEncoding)(encoding)
                    }, handleCallback(cb));
                };
            };
        };
    };
};

// | Appends the contents of a buffer to a file.
var appendFile = function (file) {
    return function (buff) {
        return function (cb) {
            return function () {
                return $foreign.appendFileImpl(file, buff, {}, handleCallback(cb));
            };
        };
    };
};
var access$prime = function (path) {
    return function (mode) {
        return function (cb) {
            return function () {
                return $foreign.accessImpl(path, mode, function (err) {
                    return cb(Data_Nullable.toMaybe(err))();
                });
            };
        };
    };
};
var access = function (path) {
    return access$prime(path)(Node_FS_Constants.defaultAccessMode);
};
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
    lstat,
    stat,
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
    fdOpen,
    fdRead,
    fdNext,
    fdWrite,
    fdAppend,
    fdClose
};
