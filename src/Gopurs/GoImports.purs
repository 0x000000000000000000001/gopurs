module Gopurs.GoImports (collectImports) where

import Prelude
import Data.Array as Array
import Data.Foldable (foldMap)
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..))
import Gopurs.GoCode (referencedImports)

runtime :: Array String
runtime = [ "gopurs/output/gopurs_runtime" ]

-- One file-level owner. Opaque fragments arrive with their dependencies;
-- structured nodes supply theirs without rendering declarations a second time.
collectImports :: Array (Array GoDecl) -> Array String
collectImports = Array.sort <<< Array.nub <<< foldMap (foldMap declImports)

declImports :: GoDecl -> Array String
declImports = case _ of
  GoCachedValue binding -> runtime <> [ "sync" ] <> typeImports binding.goType <> exprImports binding.expression
  GoStructDecl decl -> foldMap (\(Tuple _ ty) -> typeImports ty) (decl.typeParams <> decl.fields)
  GoFunctionDecl decl -> foldMap (\(Tuple _ ty) -> typeImports ty) decl.params <> typeImports decl.result <> exprImports decl.body
  GoInitDecl body -> exprImports body
  GoForeignGetter _ -> runtime

typeImports :: GoType -> Array String
typeImports = case _ of
  TypeValue -> runtime
  TypeStructPointer pointer -> foldMap typeImports pointer.typeArgs
  TypeRecord fields -> foldMap (\(Tuple _ ty) -> typeImports ty) fields
  TypeInterface name -> referencedImports name
  TypeNativeArray element -> typeImports element
  TypeFunc args ret -> foldMap typeImports args <> typeImports ret
  TypeStructValue _ fields -> foldMap typeImports fields
  _ -> []

exprImports :: GoExpr -> Array String
exprImports = case _ of
  GoVar name -> referencedImports name
  GoString _ -> []
  GoInt _ -> []
  GoCall fn args -> exprImports fn <> foldMap exprImports args
  GoSelector obj field -> exprImports obj <> case obj of
    GoVar name -> referencedImports (name <> "." <> field)
    _ -> []
  GoBlock stmts -> foldMap exprImports stmts
  GoReturn expr -> exprImports expr
  GoAssign _ expr -> bindingImports expr
  GoRecordDict ty props ->
    typeImports ty <> foldMap (\(Tuple _ expr) -> exprImports expr) props <> case ty of
      TypeRecord _ -> []
      TypeStructPointer _ -> []
      _ -> runtime
  GoRecordUpdateDict obj props -> runtime <> exprImports obj <> foldMap (\(Tuple _ expr) -> exprImports expr) props
  GoRecordUpdateStatic obj _ updates fallback -> runtime <> [ "unsafe" ] <> exprImports obj
    <> foldMap (\(Tuple _ expr) -> exprImports expr) updates <> foldMap (\(Tuple _ expr) -> exprImports expr) fallback
  GoRecordUpdateNative ty obj props -> typeImports ty <> exprImports obj <> foldMap (\(Tuple _ expr) -> exprImports expr) props
  GoIIFE _ value body -> runtime <> bindingImports value <> exprImports body
  GoLetRec bindings body -> runtime <> foldMap (\(Tuple _ expr) -> exprImports expr) bindings <> exprImports body
  GoRecordAccess obj _ -> runtime <> exprImports obj
  GoStructAccess obj _ -> exprImports obj
  GoRecordAccessStatic obj _ _ -> runtime <> exprImports obj
  GoConstructor _ _ types args ->
    -- Only a binding without a type context prints a typed nil.
    if Array.null args then [] else foldMap typeImports types <> foldMap exprImports args
  GoConstructorDict _ args -> runtime <> foldMap exprImports args
  GoConstructorAccess obj _ types _ native -> exprImports obj <> if native then [] else foldMap typeImports types
  GoBranch branches def -> runtime <> foldMap (\(Tuple cond value) -> exprImports cond <> exprImports value) branches <> exprImports def
  GoBinOp _ left right -> exprImports left <> exprImports right
  GoPrefixOp _ expr -> exprImports expr
  GoTypeAssertion expr ty -> exprImports expr <> referencedImports ty
  GoIndex expr index -> exprImports expr <> exprImports index
  GoBoxStructPointer _ expr -> runtime <> [ "unsafe" ] <> exprImports expr
  GoBoxIntArray expr -> runtime <> exprImports expr
  GoUnboxIntArray expr -> runtime <> exprImports expr
  GoFreshFilterArray expr -> exprImports expr
  GoRaw code -> code.imports
  GoFor _ stmts -> foldMap exprImports stmts
  GoForRange range stmts -> referencedImports range <> foldMap exprImports stmts
  GoContinue _ -> []
  GoMutate _ expr -> exprImports expr
  GoIfElse cond yes no -> exprImports cond <> foldMap exprImports yes <> foldMap exprImports no
  GoFuncBlock params stmts ret -> foldMap (\(Tuple _ ty) -> typeImports ty) params <> foldMap exprImports stmts <> typeImports ret
  GoFuncLit params stmts expr ret -> foldMap (\(Tuple _ ty) -> typeImports ty) params <> foldMap exprImports stmts <> exprImports expr <> typeImports ret
  GoStructValue _ types exprs -> foldMap typeImports types <> foldMap exprImports exprs

bindingImports :: GoExpr -> Array String
bindingImports (GoConstructor _ _ types []) = foldMap typeImports types
bindingImports expr = exprImports expr
