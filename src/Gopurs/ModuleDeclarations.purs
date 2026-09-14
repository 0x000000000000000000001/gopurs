module Gopurs.ModuleDeclarations (constructors) where

import Prelude
import Data.Array as Array
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Newtype (unwrap)
import Data.String as String
import Data.String.Pattern (Pattern(..), Replacement(..))
import Data.Tuple (Tuple(..))
import Gopurs.CodegenState (CodegenMetadata)
import Gopurs.GoAst (GoDecl(..), GoType(..), rawGo, sanitizeName)
import Gopurs.GoTypes (structFieldGoType)
import PureScript.Backend.Optimizer.Convert (BackendModule)
import PureScript.Backend.Optimizer.FfiSupport (hashString)

-- Constructor structs and class getter registration share their source order.
constructors :: CodegenMetadata -> String -> BackendModule -> Array GoDecl
constructors { pointerAdtPaths, enumAdts, elidedCtors, classDeclsFields } modNameStr mod =
  Array.concatMap
    (\decl -> Array.concatMap
      (\ctor ->
        let
          goFieldTypes = map (structFieldGoType pointerAdtPaths enumAdts elidedCtors decl.vars modNameStr) ctor.fields
          structName = "Constructor_" <> modNameStr <> "_" <> sanitizeName ctor.name
          structDecl = GoStructDecl
            { name: structName
            , typeParams: map (\v -> Tuple ("T_" <> sanitizeName v) (TypeInterface "any")) decl.vars
            , fields: Array.cons (Tuple "Rc" TypeUint32) (Array.mapWithIndex (\i ty -> Tuple ("V" <> show i) ty) goFieldTypes)
            }
          fullName = unwrap mod.name <> "." <> ctor.name
        in
          case Map.lookup fullName classDeclsFields of
            Nothing -> [ structDecl ]
            Just info ->
              let
                typeParamsGetter = if Array.null decl.vars then "" else "[" <> String.joinWith ", " (map (const "gopurs_runtime.Value") decl.vars) <> "]"
                cases = Array.mapWithIndex
                  (\i f ->
                    let
                      field = "c.V" <> show i
                      -- Enum fields store constructor tags, not boxed foreign integers.
                      boxed = case Array.index goFieldTypes i of
                        Just TypeUint32 -> "gopurs_runtime.Value{Type: 9, IntVal: int64(" <> field <> ")}"
                        _ -> "gopurs_runtime.Box(" <> field <> ")"
                    in
                      "\t\tcase \"" <> f.name <> "\": return " <> boxed)
                  info.fields
                pkgNameStr = String.replaceAll (Pattern ".") (Replacement "_") (unwrap mod.name)
                baseStructName = "Data_" <> pkgNameStr <> "_" <> sanitizeName ctor.name
                hashStr = hashString baseStructName
                -- The registration body remains an opaque fragment with its
                -- runtime/unsafe dependencies; its declaration is structured.
                body = rawGo ("\tgopurs_runtime.StructGetters[" <> hashStr <> "] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {\n\t\tc := (*" <> structName <> typeParamsGetter <> ")(ptr)\n\t\t_ = c\n\t\tswitch key {\n" <> String.joinWith "\n" cases <> "\n\t\tdefault: panic(\"Key not found in dictionary " <> structName <> ": \" + key)\n\t\t}\n\t}")
              in
                [ structDecl, GoInitDecl body ])
      decl.constructors)
    mod.dataDecls
