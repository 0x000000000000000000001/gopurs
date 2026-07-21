package PureScript_Backend_Optimizer_Syntax

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var eq = gopurs_runtime.Value{}
var compare = gopurs_runtime.Value{}
var eqMaybe = gopurs_runtime.Value{}
var eq7 = gopurs_runtime.Apply(eqArrayImpl, eqStringImpl)
var Pair = gopurs_runtime.Value{}
var Level = gopurs_runtime.Value{}
var OpEq = gopurs_runtime.Value{}
var OpNotEq = gopurs_runtime.Value{}
var OpGt = gopurs_runtime.Value{}
var OpGte = gopurs_runtime.Value{}
var OpLt = gopurs_runtime.Value{}
var OpLte = gopurs_runtime.Value{}
var OpAdd = gopurs_runtime.Value{}
var OpDivide = gopurs_runtime.Value{}
var OpMultiply = gopurs_runtime.Value{}
var OpSubtract = gopurs_runtime.Value{}
var OpArrayIndex = gopurs_runtime.Value{}
var OpBooleanAnd = gopurs_runtime.Value{}
var OpBooleanOr = gopurs_runtime.Value{}
var OpBooleanOrd = gopurs_runtime.Value{}
var OpCharOrd = gopurs_runtime.Value{}
var OpIntBitAnd = gopurs_runtime.Value{}
var OpIntBitOr = gopurs_runtime.Value{}
var OpIntBitShiftLeft = gopurs_runtime.Value{}
var OpIntBitShiftRight = gopurs_runtime.Value{}
var OpIntBitXor = gopurs_runtime.Value{}
var OpIntBitZeroFillShiftRight = gopurs_runtime.Value{}
var OpIntNum = gopurs_runtime.Value{}
var OpIntOrd = gopurs_runtime.Value{}
var OpNumberNum = gopurs_runtime.Value{}
var OpNumberOrd = gopurs_runtime.Value{}
var OpStringAppend = gopurs_runtime.Value{}
var OpStringOrd = gopurs_runtime.Value{}
var OpBooleanNot = gopurs_runtime.Value{}
var OpIntBitNot = gopurs_runtime.Value{}
var OpIntNegate = gopurs_runtime.Value{}
var OpNumberNegate = gopurs_runtime.Value{}
var OpArrayLength = gopurs_runtime.Value{}
var OpIsTag = gopurs_runtime.Value{}
var Op1 = gopurs_runtime.Value{}
var Op2 = gopurs_runtime.Value{}
var EffectRefNew = gopurs_runtime.Value{}
var EffectRefRead = gopurs_runtime.Value{}
var EffectRefWrite = gopurs_runtime.Value{}
var GetProp = gopurs_runtime.Value{}
var GetIndex = gopurs_runtime.Value{}
var GetCtorField = gopurs_runtime.Value{}
var Var = gopurs_runtime.Value{}
var Local = gopurs_runtime.Value{}
var Lit = gopurs_runtime.Value{}
var App = gopurs_runtime.Value{}
var Abs = gopurs_runtime.Value{}
var UncurriedApp = gopurs_runtime.Value{}
var UncurriedAbs = gopurs_runtime.Value{}
var UncurriedEffectApp = gopurs_runtime.Value{}
var UncurriedEffectAbs = gopurs_runtime.Value{}
var Accessor = gopurs_runtime.Value{}
var Update = gopurs_runtime.Value{}
var CtorSaturated = gopurs_runtime.Value{}
var CtorDef = gopurs_runtime.Value{}
var LetRec = gopurs_runtime.Value{}
var Let = gopurs_runtime.Value{}
var EffectBind = gopurs_runtime.Value{}
var EffectPure = gopurs_runtime.Value{}
var EffectDefer = gopurs_runtime.Value{}
var Branch = gopurs_runtime.Value{}
var PrimOp = gopurs_runtime.Value{}
var PrimEffect = gopurs_runtime.Value{}
var PrimUndefined = gopurs_runtime.Value{}
var Fail = gopurs_runtime.Value{}
var ordLevel = ordInt
var newtypeLevel_ = gopurs_runtime.Value{}
var functorPair = gopurs_runtime.Value{}
var functorBackendOperator = gopurs_runtime.Value{}
var functorBackendEffect = gopurs_runtime.Value{}
var functorBackendSyntax = gopurs_runtime.Value{}
var foldablePair = gopurs_runtime.Value{}

var traversablePair = gopurs_runtime.Value{}



var foldableBackendOperator = gopurs_runtime.Value{}



var traversableBackendOperato = gopurs_runtime.Value{}



var foldableBackendEffect = gopurs_runtime.Value{}



var foldableBackendSyntax = gopurs_runtime.Value{}

var traversableBackendEffect = gopurs_runtime.Value{}



var traversableBackendSyntax = gopurs_runtime.Value{}

var eqPair = gopurs_runtime.Value{}
var eqLevel = eqInt
var eqTuple2 = gopurs_runtime.Apply(eqTuple, eqMaybe)
var eq9 = gopurs_runtime.Apply(eqArrayImpl, gopurs_runtime.Value{})
var eq10 = gopurs_runtime.Apply(eqArrayImpl, gopurs_runtime.Value{})
var eqBackendOperatorOrd = gopurs_runtime.Value{}
var ordBackendOperatorOrd = gopurs_runtime.Value{}
var eqBackendOperatorNum = gopurs_runtime.Value{}
var ordBackendOperatorNum = gopurs_runtime.Value{}
var eqBackendOperator2 = gopurs_runtime.Value{}
var ordBackendOperator2 = gopurs_runtime.Value{}
var eqBackendOperator1 = gopurs_runtime.Value{}
var ordBackendOperator1 = gopurs_runtime.Value{}
var eqBackendOperator = gopurs_runtime.Value{}
var eqBackendEffect = gopurs_runtime.Value{}
var eqBackendAccessor = gopurs_runtime.Value{}
var eqBackendSyntax = gopurs_runtime.Value{}
var ordBackendAccessor = gopurs_runtime.Value{}
var syntaxOf = gopurs_runtime.Value{}
var sndPair = gopurs_runtime.Value{}
var fstPair = gopurs_runtime.Value{}