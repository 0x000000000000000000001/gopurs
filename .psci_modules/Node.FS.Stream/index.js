import * as $foreign from "./foreign.js";

// | Create a Writable stream which writes data to the specified file descriptor.
// | Unused options should not be specified. Some options
// | (e.g. `flags`, `encoding`, and `mode`) should convert their 
// | PureScript values to the corresponding JavaScript ones:
// | ```
// | filePath # fdCreateWriteStream'
// |   { flags: fileFlagsToNode R
// |   , encoding: encodingToNode UTF8
// |   , mode: permsToInt Perms.all
// |   }
// | ```
var fdCreateWriteStream$prime = function () {
    return function (f) {
        return function (opts) {
            return function () {
                return $foreign.fdCreateWriteStreamOptsImpl(f, opts);
            };
        };
    };
};

// | Create a Writable stream which writes data to the specified file
// | descriptor, using the default options.
var fdCreateWriteStream = function (f) {
    return function () {
        return $foreign.fdCreateWriteStreamImpl(f);
    };
};

// | Create a Readable stream which reads data to the specified file descriptor.
// | Unused options should not be specified. Some options
// | (e.g. `flags`, `encoding`, and `mode`) should convert their 
// | PureScript values to the corresponding JavaScript ones:
// | ```
// | filePath # fdCreateReadStream'
// |   { flags: fileFlagsToNode R
// |   , encoding: encodingToNode UTF8
// |   , mode: permsToInt Perms.all
// |   }
// | ```
var fdCreateReadStream$prime = function () {
    return function (f) {
        return function (opts) {
            return function () {
                return $foreign.fdCreateReadStreamOptsImpl(f, opts);
            };
        };
    };
};

// | Create a Readable stream which reads data to the specified file
// | descriptor, using the default options.
var fdCreateReadStream = function (f) {
    return function () {
        return $foreign.fdCreateReadStreamImpl(f);
    };
};

// | Create a Writable stream which writes data to the specified file.
// | Unused options should not be specified. Some options
// | (e.g. `flags`, `encoding`, and `mode`) should convert their 
// | PureScript values to the corresponding JavaScript ones:
// | ```
// | filePath # createWriteStream'
// |   { flags: fileFlagsToNode R
// |   , encoding: encodingToNode UTF8
// |   , mode: permsToInt Perms.all
// |   }
// | ```
var createWriteStream$prime = function () {
    return function (f) {
        return function (opts) {
            return function () {
                return $foreign.createWriteStreamOptsImpl(f, opts);
            };
        };
    };
};

// | Create a Writable stream which writes data to the specified file, using
// | the default options.
var createWriteStream = function (f) {
    return function () {
        return $foreign.createWriteStreamImpl(f);
    };
};

// | Create a Readable stream which reads data from the specified file.
// | Unused options should not be specified. Some options
// | (e.g. `flags`, `encoding`, and `mode`) should convert their 
// | PureScript values to the corresponding JavaScript ones:
// | ```
// | filePath # createReadStream'
// |   { flags: fileFlagsToNode R
// |   , encoding: encodingToNode UTF8
// |   , mode: permsToInt Perms.all
// |   }
// | ```
var createReadStream$prime = function () {
    return function (path) {
        return function (opts) {
            return function () {
                return $foreign.createReadStreamOptsImpl(path, opts);
            };
        };
    };
};

// | Create a Readable stream which reads data to the specified file, using
// | the default options.
var createReadStream = function (p) {
    return function () {
        return $foreign.createReadStreamImpl(p);
    };
};
export {
    createWriteStream,
    createWriteStream$prime,
    fdCreateWriteStream,
    fdCreateWriteStream$prime,
    createReadStream,
    createReadStream$prime,
    fdCreateReadStream,
    fdCreateReadStream$prime
};
