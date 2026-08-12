import * as Data_Eq from "../Data.Eq/index.js";
import * as Data_Ord from "../Data.Ord/index.js";
import * as Data_Show from "../Data.Show/index.js";

// | A user ID (for a process or a file).
var Uid = function (x) {
    return x;
};

// | A process ID.
var Pid = function (x) {
    return x;
};

// | A group ID (for a process or a file).
var Gid = function (x) {
    return x;
};
var showUid = {
    show: function (v) {
        return "(Uid " + (Data_Show.show(Data_Show.showInt)(v) + ")");
    }
};
var showPid = {
    show: function (v) {
        return "(Pid " + (Data_Show.show(Data_Show.showInt)(v) + ")");
    }
};
var showGid = {
    show: function (v) {
        return "(Gid " + (Data_Show.show(Data_Show.showInt)(v) + ")");
    }
};
var ordUid = Data_Ord.ordInt;
var ordPid = Data_Ord.ordInt;
var ordGid = Data_Ord.ordInt;
var newtypeUid = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypePid = {
    Coercible0: function () {
        return undefined;
    }
};
var newtypeGid = {
    Coercible0: function () {
        return undefined;
    }
};
var eqUid = Data_Eq.eqInt;
var eqPid = Data_Eq.eqInt;
var eqGid = Data_Eq.eqInt;
export {
    Pid,
    Gid,
    Uid,
    newtypePid,
    eqPid,
    ordPid,
    showPid,
    newtypeGid,
    eqGid,
    ordGid,
    showGid,
    newtypeUid,
    eqUid,
    ordUid,
    showUid
};
