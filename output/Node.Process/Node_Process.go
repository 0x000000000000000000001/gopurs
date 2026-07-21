package Node_Process

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var show = gopurs_runtime.Value{}
var showCpuUsage = gopurs_runtime.Value{}
var eqCpuUsage = gopurs_runtime.Value{}
var workerH = gopurs_runtime.Value{}
var warningH = gopurs_runtime.Value{}
var unsetEnv = gopurs_runtime.Value{}
var unsafeSendOptsCb = gopurs_runtime.Value{}
var unsafeSendOpts = gopurs_runtime.Value{}
var unsafeSendCb = gopurs_runtime.Value{}
var unsafeSend = gopurs_runtime.Value{}
var unhandledRejectionH = gopurs_runtime.Value{}
var uncaughtExceptionH = gopurs_runtime.Value{}
var setUncaughtExceptionCaptureCallback = gopurs_runtime.Value{}
var setTitle = gopurs_runtime.Value{}
var setExitCode = gopurs_runtime.Value{}
var setEnv = gopurs_runtime.Value{}
var rejectionHandledH = gopurs_runtime.Value{}
var platform = gopurs_runtime.Value{}
var nextTick' = gopurs_runtime.Value{}
var nextTick = gopurs_runtime.Value{}
var mkSignalH' = gopurs_runtime.Value{}
var mkSignalH = gopurs_runtime.Value{}
var messageH = gopurs_runtime.Value{}
var lookupEnv = gopurs_runtime.Value{}
var killStr = gopurs_runtime.Value{}
var killInt = gopurs_runtime.Value{}
var kill = gopurs_runtime.Value{}
var getUid = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Apply(pureE, toMaybe))
var getGid = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Apply(pureE, toMaybe))
var getExitCode = gopurs_runtime.Apply(gopurs_runtime.Value{}, gopurs_runtime.Apply(pureE, toMaybe))
var exitH = gopurs_runtime.Value{}
var exit' = gopurs_runtime.Value{}
var disconnectH = gopurs_runtime.Value{}
var disconnect = gopurs_runtime.Apply(runFn3, nullable)
var cpuUsageToRecord = gopurs_runtime.Value{}
var cpuUsageDiff = gopurs_runtime.Value{}
var chdir = gopurs_runtime.Value{}
var channelUnref = gopurs_runtime.Apply(runFn3, nullable)
var channelRef = gopurs_runtime.Apply(runFn3, nullable)
var beforeExitH = gopurs_runtime.Value{}
var abort = gopurs_runtime.Apply(runFn3, nullable)