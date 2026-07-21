package PureScript_CST_Traversal

import (
	"gopurs/output/gopurs_runtime"
	"fmt"
)

var _ = fmt.Println
var _ = gopurs_runtime.TypeInt

var ltraverse = gopurs_runtime.Value{}
var applicativeReaderT = gopurs_runtime.Value{}
var applyCompose = gopurs_runtime.Value{}
var applicativeCompose = gopurs_runtime.Apply(applicativeCompose, freeApplicative)
var traverseWrapped = gopurs_runtime.Value{}
var traverseSeparated = gopurs_runtime.Value{}

var traverseRecordUpdate = gopurs_runtime.Value{}

var traverseRecordLabeled = gopurs_runtime.Value{}
var traverseRecordAccessor = gopurs_runtime.Value{}
var traversePatternGuard = gopurs_runtime.Value{}
var traverseModuleBody = gopurs_runtime.Value{}
var traverseModule = gopurs_runtime.Value{}
var traverseModule1 = gopurs_runtime.Apply(traverseModule, freeApplicative)
var traverseLambda = gopurs_runtime.Value{}
var traverseLabeled = gopurs_runtime.Value{}
var traverseRow = gopurs_runtime.Value{}
var traverseTypeVarBinding = gopurs_runtime.Value{}
var traverseType = gopurs_runtime.Value{}
var traverseType1 = gopurs_runtime.Apply(traverseType, applicativeReaderT)
var traverseType2 = gopurs_runtime.Apply(traverseType, freeApplicative)
var traverseIfThenElse = gopurs_runtime.Value{}

var traverseWhere = gopurs_runtime.Value{}
var traverseValueBindingFields = gopurs_runtime.Value{}
var traverseLetBinding = gopurs_runtime.Value{}
var traverseGuardedExpr = gopurs_runtime.Value{}
var traverseGuarded = gopurs_runtime.Value{}

var traverseInstanceBinding = gopurs_runtime.Value{}
var traverseLetIn = gopurs_runtime.Value{}
var traverseForeign = gopurs_runtime.Value{}
var traverseExprAppSpine = gopurs_runtime.Value{}
var traverseDoStatement = gopurs_runtime.Value{}
var traverseDoBlock = gopurs_runtime.Value{}
var traverseDelimitedNonEmpty = gopurs_runtime.Value{}
var traverseOneOrDelimited = gopurs_runtime.Value{}
var traverseInstanceHead = gopurs_runtime.Value{}
var traverseInstance = gopurs_runtime.Value{}
var traverseDelimited = gopurs_runtime.Value{}
var traverseDataHead = gopurs_runtime.Value{}
var traverseDataCtor = gopurs_runtime.Value{}
var traverseClassHead = gopurs_runtime.Value{}
var traverseDecl = gopurs_runtime.Value{}
var traverseDecl1 = gopurs_runtime.Apply(traverseDecl, applicativeReaderT)
var traverseDecl2 = gopurs_runtime.Apply(traverseDecl, freeApplicative)
var traverseCaseOf = gopurs_runtime.Value{}
var traverseBinder = gopurs_runtime.Value{}
var traverseBinder1 = gopurs_runtime.Apply(traverseBinder, applicativeReaderT)
var traverseBinder2 = gopurs_runtime.Apply(traverseBinder, freeApplicative)
var traverseAdoBlock = gopurs_runtime.Value{}
var traverseExpr = gopurs_runtime.Value{}
var traverseExpr1 = gopurs_runtime.Apply(traverseExpr, applicativeReaderT)
var traverseExpr2 = gopurs_runtime.Apply(traverseExpr, freeApplicative)
var topDownTraversalWithContextM = gopurs_runtime.Value{}
var topDownTraversalWithContext = gopurs_runtime.Value{}
var topDownTraversal = gopurs_runtime.Value{}
var topDownPureTraversal = gopurs_runtime.Value{}
var rewriteWithContextM = gopurs_runtime.Value{}
var rewriteWithContext = gopurs_runtime.Value{}
var rewriteTypeWithContextM = gopurs_runtime.Value{}
var rewriteTypeWithContext = gopurs_runtime.Value{}
var rewriteTopDownM = gopurs_runtime.Value{}
var rewriteTypeTopDownM = gopurs_runtime.Value{}
var rewriteTopDown = gopurs_runtime.Value{}
var rewriteTypeTopDown = gopurs_runtime.Apply(rewriteTopDown, gopurs_runtime.Value{})
var rewriteModuleWithContextM = gopurs_runtime.Value{}
var rewriteModuleWithContext = gopurs_runtime.Apply(rewriteWithContext, gopurs_runtime.Apply(traverseModule, applicativeReaderT))
var rewriteModuleTopDownM = gopurs_runtime.Value{}
var rewriteModuleTopDown = gopurs_runtime.Apply(rewriteTopDown, traverseModule1)
var rewriteExprWithContextM = gopurs_runtime.Value{}
var rewriteExprWithContext = gopurs_runtime.Value{}
var rewriteExprTopDownM = gopurs_runtime.Value{}
var rewriteExprTopDown = gopurs_runtime.Apply(rewriteTopDown, gopurs_runtime.Value{})
var rewriteDeclWithContextM = gopurs_runtime.Value{}
var rewriteDeclWithContext = gopurs_runtime.Value{}
var rewriteDeclTopDownM = gopurs_runtime.Value{}
var rewriteDeclTopDown = gopurs_runtime.Apply(rewriteTopDown, gopurs_runtime.Value{})
var rewriteBinderWithContextM = gopurs_runtime.Value{}
var rewriteBinderWithContext = gopurs_runtime.Value{}
var rewriteBinderTopDownM = gopurs_runtime.Value{}
var rewriteBinderTopDown = gopurs_runtime.Apply(rewriteTopDown, gopurs_runtime.Value{})
var defer = gopurs_runtime.Value{}
var topDownMonoidalTraversal = gopurs_runtime.Value{}
var monoidalRewrite = gopurs_runtime.Value{}
var foldMapBinder = gopurs_runtime.Value{}
var foldMapDecl = gopurs_runtime.Value{}
var foldMapExpr = gopurs_runtime.Value{}
var foldMapModule = gopurs_runtime.Value{}
var foldMapType = gopurs_runtime.Value{}
var defaultVisitorWithContextM = gopurs_runtime.Value{}
var defaultVisitorWithContext = gopurs_runtime.Value{}
var defaultVisitorM = gopurs_runtime.Value{}
var defaultVisitor = gopurs_runtime.Value{}
var defaultMonoidalVisitor = gopurs_runtime.Value{}
var bottomUpTraversal = gopurs_runtime.Value{}
var rewriteBottomUpM = gopurs_runtime.Value{}
var rewriteBinderBottomUpM = gopurs_runtime.Value{}
var rewriteDeclBottomUpM = gopurs_runtime.Value{}
var rewriteExprBottomUpM = gopurs_runtime.Value{}
var rewriteModuleBottomUpM = gopurs_runtime.Value{}
var rewriteTypeBottomUpM = gopurs_runtime.Value{}
var bottomUpPureTraversal = gopurs_runtime.Value{}
var rewriteBottomUp = gopurs_runtime.Value{}
var rewriteBinderBottomUp = gopurs_runtime.Apply(rewriteBottomUp, gopurs_runtime.Value{})
var rewriteDeclBottomUp = gopurs_runtime.Apply(rewriteBottomUp, gopurs_runtime.Value{})
var rewriteExprBottomUp = gopurs_runtime.Apply(rewriteBottomUp, gopurs_runtime.Value{})
var rewriteModuleBottomUp = gopurs_runtime.Apply(rewriteBottomUp, traverseModule1)
var rewriteTypeBottomUp = gopurs_runtime.Apply(rewriteBottomUp, gopurs_runtime.Value{})