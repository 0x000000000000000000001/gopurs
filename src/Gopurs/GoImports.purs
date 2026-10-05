module Gopurs.GoImports (collectImports) where

import Prelude
import Data.Array as Array
import Data.Set (Set)
import Data.Set as Set
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..))
import Gopurs.GoCode (referencedImports)

type Imports = Set String

add :: String -> Imports -> Imports
add name imports = if Set.member name imports then imports else Set.insert name imports

runtime :: Imports -> Imports
runtime = add "gopurs/output/gopurs_runtime"

-- Both FFI implementations borrow the inputs and return a fresh array,
-- preserving order and duplicates. This collector owns normalization.
foreign import concatStringArrays :: Array (Array String) -> Array String

-- One file-level accumulator avoids copying a growing list of duplicate import
-- paths at every ancestor of a large expression. Ordering is normalized once.
collectImports :: Array (Array GoDecl) -> Array String
collectImports groups = Set.toUnfoldable (many (many declImports) groups Set.empty)

many :: forall a. (a -> Imports -> Imports) -> Array a -> Imports -> Imports
many visit items imports = Array.foldl (flip visit) imports items

fieldTypes :: Array (Tuple String GoType) -> Imports -> Imports
fieldTypes = many (\(Tuple _ ty) -> typeImports ty)

fieldExprs :: forall key. Array (Tuple key GoExpr) -> Imports -> Imports
fieldExprs = many (\(Tuple _ expr) -> exprImports expr)

referenced :: String -> Imports -> Imports
referenced = many add <<< referencedImports

declImports :: GoDecl -> Imports -> Imports
declImports decl imports = case decl of
  GoCachedValue binding -> imports # runtime # add "sync" # typeImports binding.goType # exprImports binding.expression
  GoStructDecl fields -> imports # fieldTypes fields.typeParams # fieldTypes fields.fields
  GoFunctionDecl fn -> imports # fieldTypes fn.params # typeImports fn.result # exprImports fn.body
  GoInitDecl body -> exprImports body imports
  GoForeignGetter _ -> runtime imports

typeImports :: GoType -> Imports -> Imports
typeImports ty imports = case ty of
  TypeValue -> runtime imports
  TypeStructPointer pointer -> many typeImports pointer.typeArgs imports
  TypeRecord fields -> fieldTypes fields imports
  TypeInterface name -> referenced name imports
  TypeNativeArray element -> typeImports element imports
  TypeFunc args ret -> imports # many typeImports args # typeImports ret
  TypeStructValue _ fields -> many typeImports fields imports
  _ -> imports

exprImports :: GoExpr -> Imports -> Imports
exprImports expression imports = case expression of
  GoVar name -> referenced name imports
  GoString _ -> imports
  GoInt _ -> imports
  GoCall fn args -> imports # exprImports fn # many exprImports args
  GoSelector obj field ->
    let base = exprImports obj imports
    in case obj of
      GoVar name -> referenced (name <> "." <> field) base
      _ -> base
  GoBlock stmts -> many exprImports stmts imports
  GoReturn expr -> exprImports expr imports
  GoAssign _ expr -> bindingImports expr imports
  GoRecordDict ty props ->
    let base = imports # typeImports ty # fieldExprs props
    in case ty of
      TypeRecord _ -> base
      TypeStructPointer _ -> base
      _ -> runtime base
  GoRecordUpdateDict obj props -> imports # runtime # exprImports obj # fieldExprs props
  GoRecordUpdateStatic obj _ updates fallback -> imports # runtime # add "unsafe" # exprImports obj
    # fieldExprs updates # fieldExprs fallback
  GoRecordUpdateNative ty obj props -> imports # typeImports ty # exprImports obj # fieldExprs props
  GoIIFE _ value body -> imports # runtime # bindingImports value # exprImports body
  GoLetRec bindings body -> imports # runtime # fieldExprs bindings # exprImports body
  GoRecordAccess obj _ -> imports # runtime # exprImports obj
  GoStructAccess obj _ -> exprImports obj imports
  GoRecordAccessStatic obj _ _ -> imports # runtime # exprImports obj
  GoConstructor _ _ types args ->
    -- Only a binding without a type context prints a typed nil.
    if Array.null args then imports else imports # many typeImports types # many exprImports args
  GoConstructorDict _ args -> imports # runtime # many exprImports args
  GoConstructorAccess obj _ types _ native ->
    let base = exprImports obj imports
    in if native then base else many typeImports types base
  GoBranch branches def -> imports # runtime
    # many (\(Tuple cond value) -> exprImports value <<< exprImports cond) branches # exprImports def
  GoBinOp _ left right -> imports # exprImports left # exprImports right
  GoPrefixOp _ expr -> exprImports expr imports
  GoTypeAssertion expr ty -> imports # exprImports expr # referenced ty
  GoIndex expr index -> imports # exprImports expr # exprImports index
  GoBoxStructPointer _ expr -> imports # runtime # add "unsafe" # exprImports expr
  GoBoxIntArray expr -> imports # runtime # exprImports expr
  GoUnboxIntArray expr -> imports # runtime # exprImports expr
  GoFreshFilterArray expr -> exprImports expr imports
  GoRaw code -> many add code.imports imports
  GoFor _ stmts -> many exprImports stmts imports
  GoForRange range stmts -> imports # referenced range # many exprImports stmts
  GoContinue _ -> imports
  GoMutate _ expr -> exprImports expr imports
  GoIfElse cond yes no -> imports # exprImports cond # many exprImports yes # many exprImports no
  GoFuncBlock params stmts ret -> imports # fieldTypes params # many exprImports stmts # typeImports ret
  GoFuncLit params stmts expr ret -> imports # fieldTypes params # many exprImports stmts # exprImports expr # typeImports ret
  GoStructValue _ types exprs -> imports # many typeImports types # many exprImports exprs

bindingImports :: GoExpr -> Imports -> Imports
bindingImports (GoConstructor _ _ types []) imports = many typeImports types imports
bindingImports expr imports = exprImports expr imports
