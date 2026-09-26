module Gopurs.GoImports (collectImports) where

import Prelude
import Data.Array as Array
import Data.Tuple (Tuple(..))
import Gopurs.GoAst (GoDecl(..), GoExpr(..), GoType(..))
import Gopurs.GoCode (referencedImports)

runtime :: Array String
runtime = [ "gopurs/output/gopurs_runtime" ]

-- Concaténation native de tableaux de chaînes. Les folds PureScript
-- recopient l'accumulateur à chaque étape (`foldMap` par défaut des tableaux
-- est quadratique), et le regroupement des imports d'un gros module domine
-- alors le codegen.
foreign import concatStringArrays :: Array (Array String) -> Array String

-- One file-level owner. Opaque fragments arrive with their dependencies;
-- structured nodes supply theirs without rendering declarations a second time.
collectImports :: Array (Array GoDecl) -> Array String
collectImports groups = Array.sort (Array.nub (concatStringArrays (map declGroupImports groups)))
  where
  declGroupImports decls = concatStringArrays (map declImports decls)

declImports :: GoDecl -> Array String
declImports = case _ of
  GoCachedValue binding -> runtime <> [ "sync" ] <> typeImports binding.goType <> exprImports binding.expression
  GoStructDecl decl -> concatStringArrays (map (\(Tuple _ ty) -> typeImports ty) (decl.typeParams <> decl.fields))
  GoFunctionDecl decl -> concatStringArrays (map (\(Tuple _ ty) -> typeImports ty) decl.params) <> typeImports decl.result <> exprImports decl.body
  GoInitDecl body -> exprImports body
  GoForeignGetter _ -> runtime

typeImports :: GoType -> Array String
typeImports = case _ of
  TypeValue -> runtime
  TypeStructPointer pointer -> concatStringArrays (map typeImports pointer.typeArgs)
  TypeRecord fields -> concatStringArrays (map (\(Tuple _ ty) -> typeImports ty) fields)
  TypeInterface name -> referencedImports name
  TypeNativeArray element -> typeImports element
  TypeFunc args ret -> concatStringArrays (map typeImports args) <> typeImports ret
  TypeStructValue _ fields -> concatStringArrays (map typeImports fields)
  _ -> []

exprImports :: GoExpr -> Array String
exprImports = case _ of
  GoVar name -> referencedImports name
  GoString _ -> []
  GoInt _ -> []
  GoCall fn args -> exprImports fn <> concatStringArrays (map exprImports args)
  GoSelector obj field -> exprImports obj <> case obj of
    GoVar name -> referencedImports (name <> "." <> field)
    _ -> []
  GoBlock stmts -> concatStringArrays (map exprImports stmts)
  GoReturn expr -> exprImports expr
  GoAssign _ expr -> bindingImports expr
  GoRecordDict ty props ->
    typeImports ty <> concatStringArrays (map (\(Tuple _ expr) -> exprImports expr) props) <> case ty of
      TypeRecord _ -> []
      TypeStructPointer _ -> []
      _ -> runtime
  GoRecordUpdateDict obj props -> runtime <> exprImports obj <> concatStringArrays (map (\(Tuple _ expr) -> exprImports expr) props)
  GoRecordUpdateStatic obj _ updates fallback -> runtime <> [ "unsafe" ] <> exprImports obj
    <> concatStringArrays (map (\(Tuple _ expr) -> exprImports expr) updates) <> concatStringArrays (map (\(Tuple _ expr) -> exprImports expr) fallback)
  GoRecordUpdateNative ty obj props -> typeImports ty <> exprImports obj <> concatStringArrays (map (\(Tuple _ expr) -> exprImports expr) props)
  GoIIFE _ value body -> runtime <> bindingImports value <> exprImports body
  GoLetRec bindings body -> runtime <> concatStringArrays (map (\(Tuple _ expr) -> exprImports expr) bindings) <> exprImports body
  GoRecordAccess obj _ -> runtime <> exprImports obj
  GoStructAccess obj _ -> exprImports obj
  GoRecordAccessStatic obj _ _ -> runtime <> exprImports obj
  GoConstructor _ _ types args ->
    -- Only a binding without a type context prints a typed nil.
    if Array.null args then [] else concatStringArrays (map typeImports types) <> concatStringArrays (map exprImports args)
  GoConstructorDict _ args -> runtime <> concatStringArrays (map exprImports args)
  GoConstructorAccess obj _ types _ native -> exprImports obj <> if native then [] else concatStringArrays (map typeImports types)
  GoBranch branches def -> runtime <> concatStringArrays (map (\(Tuple cond value) -> exprImports cond <> exprImports value) branches) <> exprImports def
  GoBinOp _ left right -> exprImports left <> exprImports right
  GoPrefixOp _ expr -> exprImports expr
  GoTypeAssertion expr ty -> exprImports expr <> referencedImports ty
  GoIndex expr index -> exprImports expr <> exprImports index
  GoBoxStructPointer _ expr -> runtime <> [ "unsafe" ] <> exprImports expr
  GoBoxIntArray expr -> runtime <> exprImports expr
  GoUnboxIntArray expr -> runtime <> exprImports expr
  GoFreshFilterArray expr -> exprImports expr
  GoRaw code -> code.imports
  GoFor _ stmts -> concatStringArrays (map exprImports stmts)
  GoForRange range stmts -> referencedImports range <> concatStringArrays (map exprImports stmts)
  GoContinue _ -> []
  GoMutate _ expr -> exprImports expr
  GoIfElse cond yes no -> exprImports cond <> concatStringArrays (map exprImports yes) <> concatStringArrays (map exprImports no)
  GoFuncBlock params stmts ret -> concatStringArrays (map (\(Tuple _ ty) -> typeImports ty) params) <> concatStringArrays (map exprImports stmts) <> typeImports ret
  GoFuncLit params stmts expr ret -> concatStringArrays (map (\(Tuple _ ty) -> typeImports ty) params) <> concatStringArrays (map exprImports stmts) <> exprImports expr <> typeImports ret
  GoStructValue _ types exprs -> concatStringArrays (map typeImports types) <> concatStringArrays (map exprImports exprs)

bindingImports :: GoExpr -> Array String
bindingImports (GoConstructor _ _ types []) = concatStringArrays (map typeImports types)
bindingImports expr = exprImports expr
