import * as Data_Either from "../Data.Either/index.js";
import * as Data_Functor from "../Data.Functor/index.js";
import * as Effect from "../Effect/index.js";
import * as Effect_Aff from "../Effect.Aff/index.js";
import * as Node_FS_Async from "../Node.FS.Async/index.js";
var toAff = function (p) {
    return Effect_Aff.makeAff(function (k) {
        return Data_Functor.voidLeft(Effect.functorEffect)(p(k))(Effect_Aff.nonCanceler);
    });
};
var toAff1 = function (f) {
    return function (a) {
        return toAff(f(a));
    };
};

// |
// | Deletes a file.
// |
var unlink = /* #__PURE__ */ toAff1(Node_FS_Async.unlink);
var toAff2 = function (f) {
    return function (a) {
        return function (b) {
            return toAff(f(a)(b));
        };
    };
};

// |
// | Truncates a file to the specified length.
// |
var truncate = /* #__PURE__ */ toAff2(Node_FS_Async.truncate);

// |
// | Writes a buffer to a file.
// |
var writeFile = /* #__PURE__ */ toAff2(Node_FS_Async.writeFile);
var toAff3 = function (f) {
    return function (a) {
        return function (b) {
            return function (c) {
                return toAff(f(a)(b)(c));
            };
        };
    };
};

// |
// | Sets the accessed and modified times for the specified file.
// |
var utimes = /* #__PURE__ */ toAff3(Node_FS_Async.utimes);

// |
// | Writes text to a file using the specified encoding.
// |
var writeTextFile = /* #__PURE__ */ toAff3(Node_FS_Async.writeTextFile);
var toAff5 = function (f) {
    return function (a) {
        return function (b) {
            return function (c) {
                return function (d) {
                    return function (e) {
                        return toAff(f(a)(b)(c)(d)(e));
                    };
                };
            };
        };
    };
};

// |
// | Creates a symlink.
// |
var symlink = /* #__PURE__ */ toAff3(Node_FS_Async.symlink);

// |
// | Gets file statistics.
// |
var stat = /* #__PURE__ */ toAff1(Node_FS_Async.stat);

// |
// | Deletes a directory with options.
// |
var rmdir$prime = /* #__PURE__ */ toAff2(Node_FS_Async["rmdir$prime"]);

// |
// | Deletes a directory.
// |
var rmdir = /* #__PURE__ */ toAff1(Node_FS_Async.rmdir);

// |
// | Deletes a file or directory with options.
// |
var rm$prime = /* #__PURE__ */ toAff2(Node_FS_Async["rm$prime"]);

// |
// | Deletes a file or directory.
// |
var rm = /* #__PURE__ */ toAff1(Node_FS_Async.rm);

// |
// | Rename a file.
// |
var rename = /* #__PURE__ */ toAff2(Node_FS_Async.rename);

// |
// | Find the canonicalized absolute location for a path using a cache object
// | for already resolved paths.
// |
var realpath$prime = /* #__PURE__ */ toAff2(Node_FS_Async["realpath$prime"]);

// |
// | Find the canonicalized absolute location for a path.
// |
var realpath = /* #__PURE__ */ toAff1(Node_FS_Async.realpath);

// |
// | Reads the value of a symlink.
// |
var readlink = /* #__PURE__ */ toAff1(Node_FS_Async.readlink);

// |
// | Reads the contents of a directory.
// |
var readdir = /* #__PURE__ */ toAff1(Node_FS_Async.readdir);

// |
// | Reads the entire contents of a text file with the specified encoding.
// |
var readTextFile = /* #__PURE__ */ toAff2(Node_FS_Async.readTextFile);

// |
// | Reads the entire contents of a file returning the result as a raw buffer.
// |
var readFile = /* #__PURE__ */ toAff1(Node_FS_Async.readFile);
var mkdtemp$prime = /* #__PURE__ */ toAff2(Node_FS_Async["mkdtemp$prime"]);
var mkdtemp = /* #__PURE__ */ toAff1(Node_FS_Async.mkdtemp);

// |
// | Makes a new directory with all of its options.
// |
var mkdir$prime = /* #__PURE__ */ toAff2(Node_FS_Async["mkdir$prime"]);

// |
// | Makes a new directory.
// |
var mkdir = /* #__PURE__ */ toAff1(Node_FS_Async.mkdir);

// | Gets file or symlink statistics. `lstat` is identical to `stat`, except
// | that if the `FilePath` is a symbolic link, then the link itself is stat-ed,
// | not the file that it refers to.
var lstat = /* #__PURE__ */ toAff1(Node_FS_Async.lstat);

// |
// | Creates a link to an existing file.
// |
var link = /* #__PURE__ */ toAff2(Node_FS_Async.link);

// | Write to a file asynchronously. See the [Node Documentation](https://nodejs.org/api/fs.html#fs_fs_write_fd_buffer_offset_length_position_callback)
// | for details.
var fdWrite = /* #__PURE__ */ toAff5(Node_FS_Async.fdWrite);

// | Read from a file asynchronously. See the [Node Documentation](https://nodejs.org/api/fs.html#fs_fs_read_fd_buffer_offset_length_position_callback)
// | for details.
var fdRead = /* #__PURE__ */ toAff5(Node_FS_Async.fdRead);

// | Open a file asynchronously. See the [Node Documentation](https://nodejs.org/api/fs.html#fs_fs_open_path_flags_mode_callback)
// | for details.
var fdOpen = /* #__PURE__ */ toAff3(Node_FS_Async.fdOpen);

// | Convenience function to fill the whole buffer from the current
// | file position.
var fdNext = /* #__PURE__ */ toAff2(Node_FS_Async.fdNext);

// | Close a file asynchronously. See the [Node Documentation](https://nodejs.org/api/fs.html#fs_fs_close_fd_callback)
// | for details.
var fdClose = /* #__PURE__ */ toAff1(Node_FS_Async.fdClose);

// | Convenience function to append the whole buffer to the current
// | file position.
var fdAppend = /* #__PURE__ */ toAff2(Node_FS_Async.fdAppend);
var copyFile$prime = /* #__PURE__ */ toAff3(Node_FS_Async["copyFile$prime"]);
var copyFile = /* #__PURE__ */ toAff2(Node_FS_Async.copyFile);

// |
// | Changes the ownership of a file.
// |
var chown = /* #__PURE__ */ toAff3(Node_FS_Async.chown);

// |
// | Changes the permissions of a file.
// |
var chmod = /* #__PURE__ */ toAff2(Node_FS_Async.chmod);

// |
// | Appends text to a file using the specified encoding.
// |
var appendTextFile = /* #__PURE__ */ toAff3(Node_FS_Async.appendTextFile);

// |
// | Appends the contents of a buffer to a file.
// |
var appendFile = /* #__PURE__ */ toAff2(Node_FS_Async.appendFile);
var access$prime = function (path) {
    return function (mode) {
        return Effect_Aff.makeAff(function (k) {
            return function __do() {
                Node_FS_Async["access$prime"](path)(mode)(function ($0) {
                    return k(Data_Either.Right.create($0));
                })();
                return Effect_Aff.nonCanceler;
            };
        });
    };
};
var access = function (path) {
    return Effect_Aff.makeAff(function (k) {
        return function __do() {
            Node_FS_Async.access(path)(function ($1) {
                return k(Data_Either.Right.create($1));
            })();
            return Effect_Aff.nonCanceler;
        };
    });
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
    fdOpen,
    fdRead,
    fdNext,
    fdWrite,
    fdAppend,
    fdClose
};
