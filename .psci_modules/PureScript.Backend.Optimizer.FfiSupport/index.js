import * as $foreign from "./foreign.js";
import * as Data_Nullable from "../Data.Nullable/index.js";

// | Finds the FFI file corresponding to a PureScript module
// | extension: e.g. ".go" or ".php"
// | extraSpagoDirs: optional extra directories to scan like "bak/spago.d/php/p"
// | mbFfiDir: optional directory to search in (if not provided, searches .spago and local dirs)
// | modName: the PureScript module name (e.g. "Data.Show")
// | mbModulePath: the path to the original .purs file if known
var findFfiFile = function (extension) {
    return function (extraSpagoDirs) {
        return function (mbFfiDir) {
            return function (modName) {
                return function (mbModulePath) {
                    return function __do() {
                        var path = $foreign.findFfiFileImpl(extension)(extraSpagoDirs)(Data_Nullable.toNullable(mbFfiDir))(modName)(Data_Nullable.toNullable(mbModulePath))();
                        return Data_Nullable.toMaybe(path);
                    };
                };
            };
        };
    };
};
export {
    hashString
} from "./foreign.js";
export {
    findFfiFile
};
