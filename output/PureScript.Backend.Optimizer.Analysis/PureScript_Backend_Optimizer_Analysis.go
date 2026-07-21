package PureScript_Backend_Optimizer_Analysis

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var union = gopurs_runtime.Apply(union, ordInt)
var ordQualified = gopurs_runtime.Apply(ordQualified, ordString)
var union1 = gopurs_runtime.Apply(union, ordQualified)
var pop = gopurs_runtime.Apply(pop, ordInt)
var KnownNeutral = gopurs_runtime.Value{}
var Unknown = gopurs_runtime.Value{}
var Trivial = gopurs_runtime.Value{}
var Deref = gopurs_runtime.Value{}
var KnownSize = gopurs_runtime.Value{}
var NonTrivial = gopurs_runtime.Value{}
var CaptureNone = gopurs_runtime.Value{}
var CaptureBranch = gopurs_runtime.Value{}
var CaptureClosure = gopurs_runtime.Value{}
var Usage = gopurs_runtime.Value{}
var BackendAnalysis = gopurs_runtime.Value{}
var semigroupResultTerm = gopurs_runtime.Value{}
var newtypeUsage_ = gopurs_runtime.Value{}
var newtypeBackendAnalysis_ = gopurs_runtime.Value{}
var monoidResultTerm = gopurs_runtime.Value{}
var foldMap1 = gopurs_runtime.Apply(gopurs_runtime.Value{}, monoidResultTerm)
var eqResultTerm = gopurs_runtime.Value{}
var eqComplexity = gopurs_runtime.Value{}
var ordComplexity = gopurs_runtime.Value{}
var semigroupComplexity = gopurs_runtime.Value{}
var monoidComplexity = gopurs_runtime.Value{}
var eqCapture = gopurs_runtime.Value{}
var ordCapture = gopurs_runtime.Value{}
var semigroupCapture = gopurs_runtime.Value{}
var monoidCapture = gopurs_runtime.Value{}
var semigroupUsage = gopurs_runtime.Value{}
var monoidUsage = gopurs_runtime.Value{}
var semigroupBackendAnalysis = gopurs_runtime.Value{}
var monoidBackendAnalysis = gopurs_runtime.Value{}
var foldMap2 = gopurs_runtime.Apply(gopurs_runtime.Value{}, monoidBackendAnalysis)
var foldMap3 = gopurs_runtime.Apply(gopurs_runtime.Value{}, monoidBackendAnalysis)
var foldMap4 = gopurs_runtime.Apply(gopurs_runtime.Value{}, monoidBackendAnalysis)
var foldMap6 = gopurs_runtime.Value{}
var withRewrite = gopurs_runtime.Value{}
var withResult = gopurs_runtime.Value{}
var withArgs = gopurs_runtime.Value{}
var usedDep = gopurs_runtime.Value{}
var used = gopurs_runtime.Value{}
var updated = gopurs_runtime.Value{}
var externs = gopurs_runtime.Value{}
var complex = gopurs_runtime.Value{}
var cased = gopurs_runtime.Value{}
var capture = gopurs_runtime.Value{}
var callArity = gopurs_runtime.Value{}
var bump = gopurs_runtime.Value{}
var boundArg = gopurs_runtime.Value{}
var bound = gopurs_runtime.Value{}
var analysisOf = gopurs_runtime.Value{}
var analyzeDefault = gopurs_runtime.Value{}
var resultOf = gopurs_runtime.Value{}
var accessed = gopurs_runtime.Value{}
var analyze = gopurs_runtime.Value{}
var analyzeEffectBlock = gopurs_runtime.Value{}
