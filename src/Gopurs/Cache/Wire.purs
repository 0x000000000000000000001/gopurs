module Gopurs.Cache.Wire
  ( class Codec, encode, decode, write, read, hash, genericEncode, genericDecode
  , class Constructors, class Arguments, class Fields
  , encodeConstructor, decodeConstructor, argumentCount, encodeArguments, decodeArguments
  , encodeFields, decodeFields
  , class KeyOrder, compareKey
  , TypeTable, EncodeContext, DecodeContext
  ) where

import Prelude

import Control.Alt ((<|>))
import Control.Alternative (guard)
import Data.Argonaut.Core (Json, fromString, toArray, toBoolean, toString)
import Data.Argonaut.Parser (jsonParser)
import Data.Array as Array
import Data.Array.NonEmpty (NonEmptyArray)
import Data.Array.NonEmpty as NEA
import Data.Either (hush)
import Data.Foldable (foldM)
import Data.FoldableWithIndex (foldlWithIndex)
import Data.Generic.Rep (class Generic, Argument(..), Constructor(..), NoArguments(..), Product(..), Sum(..), from, to)
import Data.Int as Int
import Data.List (List)
import Data.List as List
import Data.Map (Map)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.Set (Set)
import Data.Set as Set
import Data.String as String
import Data.String.CodeUnits as CU
import Data.Symbol (class IsSymbol, reflectSymbol)
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..), fst)
import Prim.Row as Row
import Prim.RowList (class RowToList, Cons, Nil, RowList)
import Record as Record
import Effect (Effect)
import Gopurs.Cache.Pool as Pool
import Gopurs.Cache.TypeMemo as TypeMemo
import Gopurs.GoAst (GoType)
import Gopurs.GoSymbol (GoReference)
import PureScript.Backend.Optimizer.Analysis as A
import PureScript.Backend.Optimizer.CoreFn as C
import PureScript.Backend.Optimizer.Semantics as S
import PureScript.Backend.Optimizer.Syntax as B
import Type.Proxy (Proxy(..))

-- Host helpers handle scalars, string interning and a private identity shortcut.
-- All business constructors, field names, arities and ordered-container invariants
-- are defined here; host identities never enter the portable representation.
foreign import stringUnits :: String -> String
foreign import unitsString :: String -> String
foreign import numberBits :: Number -> String
foreign import bitsNumber :: String -> Number
foreign import hash :: String -> String
foreign import compareStrings :: String -> String -> Int
foreign import hex :: String -> Boolean

-- A type's canonical graph reference is a structural key. The separate pool
-- assigns type indices in first-encounter postorder, independent of identity.
data TypeTable = TypeTable Pool.Pool
type EncodeContext = { types :: TypeTable, pool :: Pool.Pool, typeMemo :: TypeMemo.TypeMemo }
type DecodeContext =
  { types :: Map Int C.ExprType
  , nodes :: Array Json
  , expressions :: Map Int S.NeutralExpr
  }

class Codec a where
  encode :: EncodeContext -> a -> Effect String
  decode :: DecodeContext -> Json -> Maybe a

write :: forall a. Codec a => a -> Effect String
write value = do
  table <- Pool.create
  pool <- Pool.create
  typeMemo <- TypeMemo.create
  root <- encode { types: TypeTable table, pool, typeMemo } value
  entries <- Pool.finish table
  nodes <- Pool.finish pool
  let body = ctor "WithTypes" [ node "array" [ "[" <> entries <> "]" ], root ]
  pure ("[\"ctor\",\"WithGraph\",[[\"array\",[" <> nodes <> "]]," <> body <> "]]")

read :: forall a. Codec a => String -> Maybe a
read text = do
  json <- hush (jsonParser text)
  graph <- rawCtorArgs "WithGraph" json
  case graph of
    [ entries, body ] -> do
      nodes <- rawArray entries
      guard (Array.all identity (Array.mapWithIndex validNode nodes))
      args <- rawCtorArgs "WithTypes" body
      case args of
        [ table, root ] -> do
          entries' <- rawArray table
          types <- foldM (\acc entry -> do
            value <- decodeType { types: acc, nodes, expressions: Map.empty } entry
            pure (Map.insert (Map.size acc) value acc)) Map.empty entries'
          -- Graph edges point backwards, so decode each expression once using
          -- only previously decoded expressions. Preserve sharing instead of
          -- expanding repeated graph references into an exponentially large tree.
          -- Failed/ill-typed nodes remain absent and fail any business read of them.
          let expressions = foldlWithIndex (decodeExpression types nodes) Map.empty nodes
          decode { types, nodes, expressions } root
        _ -> Nothing
    _ -> Nothing

decodeExpression :: Map Int C.ExprType -> Array Json -> Int -> Map Int S.NeutralExpr -> Json -> Map Int S.NeutralExpr
decodeExpression types nodes index expressions json =
  case rawCtorArgs "NeutralExpr" json of
    Just [ body ] -> case S.NeutralExpr <$> decodeBackendSyntax { types, nodes, expressions } body of
      Just value -> Map.insert index value expressions
      Nothing -> expressions
    _ -> expressions

rawArray :: Json -> Maybe (Array Json)
rawArray json = do
  fields <- toArray json
  case fields of
    [ tag, values ] | toString tag == Just "array" -> toArray values
    _ -> Nothing

rawCtorArgs :: String -> Json -> Maybe (Array Json)
rawCtorArgs name json = do
  fields <- toArray json
  case fields of
    [ tag, name', values ] | toString tag == Just "ctor" && toString name' == Just name -> toArray values
    _ -> Nothing

reference :: Json -> Maybe Int
reference json = do
  fields <- toArray json
  case fields of
    [ tag, value ] | toString tag == Just "int" -> do
      text <- toString value
      value' <- Int.fromString text
      guard (value' >= 0 && show value' == text)
      pure value'
    _ -> Nothing

-- Every graph edge points backwards. Reject cycles/forward references before
-- business decoding follows the graph's edges.
validNode :: Int -> Json -> Boolean
validNode limit json = case toArray json of
  Just [ tag, values ] | toString tag == Just "array" -> edges values
  Just [ tag, _, values ] | toString tag == Just "ctor" -> edges values
  Just [ tag, values ] | toString tag == Just "record" -> case toArray values of
    Just fields -> Array.all (\field -> case toArray field of
      Just [ _, value ] -> edge value
      _ -> false) fields
    _ -> false
  Just [ tag, _ ] -> Array.elem (toString tag) (map Just [ "boolean", "int", "number", "string" ])
  Just [ tag ] -> toString tag == Just "unit"
  _ -> false
  where
  edge value = case reference value of
    Just index -> index < limit
    Nothing -> false
  edges value = case toArray value of
    Just values -> Array.all edge values
    Nothing -> false

resolve :: DecodeContext -> Json -> Maybe Json
resolve context json = reference json >>= Array.index context.nodes

intern :: EncodeContext -> String -> Effect String
intern context value = do
  index <- Pool.intern context.pool value
  pure ("[\"int\",\"" <> show index <> "\"]")

internCtor :: EncodeContext -> String -> Array String -> Effect String
internCtor context name args = intern context (ctor name args)

-- Direct constructor codecs avoid rebuilding recursive Generic dictionaries and
-- materializing Sum/Product trees at every AST node. The wire layout is shared
-- with genericEncode, including left-to-right child interning.
one :: forall a. Codec a => EncodeContext -> String -> a -> Effect String
one context tag a = do
  a' <- encode context a
  internCtor context tag [ a' ]

two :: forall a b. Codec a => Codec b => EncodeContext -> String -> a -> b -> Effect String
two context tag a b = do
  a' <- encode context a
  b' <- encode context b
  internCtor context tag [ a', b' ]

three :: forall a b c. Codec a => Codec b => Codec c => EncodeContext -> String -> a -> b -> c -> Effect String
three context tag a b c = do
  a' <- encode context a
  b' <- encode context b
  c' <- encode context c
  internCtor context tag [ a', b', c' ]

four :: forall a b c d. Codec a => Codec b => Codec c => Codec d => EncodeContext -> String -> a -> b -> c -> d -> Effect String
four context tag a b c d = do
  a' <- encode context a
  b' <- encode context b
  c' <- encode context c
  d' <- encode context d
  internCtor context tag [ a', b', c', d' ]

five :: forall a b c d e. Codec a => Codec b => Codec c => Codec d => Codec e => EncodeContext -> String -> a -> b -> c -> d -> e -> Effect String
five context tag a b c d e = do
  a' <- encode context a
  b' <- encode context b
  c' <- encode context c
  d' <- encode context d
  e' <- encode context e
  internCtor context tag [ a', b', c', d', e' ]

-- All text passed to quoted is a fixed schema name, canonical numeric text,
-- or UTF-16 hex. No source string enters this renderer without stringUnits.
quoted :: String -> String
quoted text = "\"" <> text <> "\""

jsonArray :: Array String -> String
jsonArray values = "[" <> String.joinWith "," values <> "]"

node :: String -> Array String -> String
node tag args = "[\"" <> tag <> "\"" <> (if Array.null args then "" else "," <> String.joinWith "," args) <> "]"

payload :: DecodeContext -> String -> Json -> Maybe Json
payload context tag json = do
  fields <- resolve context json >>= toArray
  case fields of
    [ name, value ] | toString name == Just tag -> pure value
    _ -> Nothing

instance Codec Unit where
  encode context _ = intern context "[\"unit\"]"
  decode context json = do
    fields <- resolve context json >>= toArray
    guard (fields == [ fromString "unit" ])
    pure unit

instance Codec Boolean where
  encode context value = intern context ("[\"boolean\"," <> show value <> "]")
  decode context json = payload context "boolean" json >>= toBoolean

instance Codec Int where
  encode context value = intern context ("[\"int\",\"" <> show value <> "\"]")
  decode context json = do
    text <- payload context "int" json >>= toString
    value <- Int.fromString text
    guard (show value == text)
    pure value

instance Codec Number where
  encode context value = intern context ("[\"number\",\"" <> numberBits value <> "\"]")
  decode context json = do
    text <- payload context "number" json >>= toString
    guard (CU.length text == 16 && hex text)
    let value = bitsNumber text
    guard (numberBits value == text)
    pure value

instance Codec String where
  encode context value = intern context ("[\"string\",\"" <> stringUnits value <> "\"]")
  decode context json = do
    text <- payload context "string" json >>= toString
    guard (CU.length text `mod` 4 == 0 && hex text)
    let value = unitsString text
    guard (stringUnits value == text)
    pure value

instance Codec Char where
  encode context value = encode context (CU.singleton value) >>= internCtor context "Char" <<< Array.singleton
  decode context json = do
    args <- ctorArgs context "Char" json
    case args of
      [ value ] -> do
        text <- decode context value
        guard (CU.length text == 1)
        CU.charAt 0 text
      _ -> Nothing

instance Codec a => Codec (Array a) where
  encode context values = do
    encoded <- traverse (encode context) values
    intern context ("[\"array\",[" <> String.joinWith "," encoded <> "]]")
  decode context json = payload context "array" json >>= toArray >>= traverse (decode context)

instance Codec a => Codec (NonEmptyArray a) where
  encode context values = do
    encoded <- encode context (NEA.toArray values)
    internCtor context "NonEmptyArray" [ encoded ]
  decode context json = do
    args <- ctorArgs context "NonEmptyArray" json
    case args of
      [ values ] -> decode context values >>= NEA.fromArray
      _ -> Nothing

instance Codec a => Codec (List a) where
  encode context values = do
    encoded <- encode context (Array.fromFoldable values)
    internCtor context "List" [ encoded ]
  decode context json = do
    args <- ctorArgs context "List" json
    case args of
      [ values ] -> List.fromFoldable <$> (decode context values :: Maybe (Array a))
      _ -> Nothing

instance Codec a => Codec (Maybe a) where
  encode context = case _ of
    Nothing -> internCtor context "Nothing" []
    Just value -> one context "Just" value
  decode context json = do
    value <- constructor context json
    case value.tag, value.args of
      "Nothing", [] -> Just Nothing
      "Just", [ a ] -> Just <$> decode context a
      _, _ -> Nothing

instance (Codec a, Codec b) => Codec (Tuple a b) where
  encode context (Tuple a b) = two context "Tuple" a b
  decode context json = do
    args <- ctorArgs context "Tuple" json
    case args of
      [ a, b ] -> Tuple <$> decode context a <*> decode context b
      _ -> Nothing

-- Native maps retain their host's ordering internally. The persistent format
-- uses PureScript/JS ordering, including UTF-16 (not UTF-8) string comparison.
-- Decoding rebuilds the map with its ordinary host comparator after checking
-- the canonical portable list; it never serializes the balancing/tree shape.
class KeyOrder a where
  compareKey :: a -> a -> Ordering

instance KeyOrder String where
  compareKey a b = compare (compareStrings a b) 0
instance KeyOrder Int where
  compareKey = compare
instance KeyOrder C.Ident where
  compareKey (C.Ident a) (C.Ident b) = compareKey a b
instance KeyOrder C.ModuleName where
  compareKey (C.ModuleName a) (C.ModuleName b) = compareKey a b
instance KeyOrder C.ProperName where
  compareKey (C.ProperName a) (C.ProperName b) = compareKey a b
instance KeyOrder B.Level where
  compareKey (B.Level a) (B.Level b) = compareKey a b
instance KeyOrder a => KeyOrder (Maybe a) where
  compareKey Nothing Nothing = EQ
  compareKey Nothing (Just _) = LT
  compareKey (Just _) Nothing = GT
  compareKey (Just a) (Just b) = compareKey a b
instance KeyOrder a => KeyOrder (C.Qualified a) where
  compareKey (C.Qualified am a) (C.Qualified bm b) = case compareKey am bm of
    EQ -> compareKey a b
    order -> order
instance KeyOrder C.ReExport where
  compareKey (C.ReExport am a) (C.ReExport bm b) = case compareKey am bm of
    EQ -> compareKey a b
    order -> order
instance KeyOrder S.EvalRef where
  compareKey (S.EvalExtern a) (S.EvalExtern b) = compareKey a b
  compareKey (S.EvalExtern _) (S.EvalLocal _ _) = LT
  compareKey (S.EvalLocal _ _) (S.EvalExtern _) = GT
  compareKey (S.EvalLocal am a) (S.EvalLocal bm b) = case compareKey am bm of
    EQ -> compareKey a b
    order -> order
instance KeyOrder S.InlineAccessor where
  compareKey (S.InlineProp a) (S.InlineProp b) = compareKey a b
  compareKey (S.InlineProp _) _ = LT
  compareKey _ (S.InlineProp _) = GT
  compareKey (S.InlineSpineProp a) (S.InlineSpineProp b) = compareKey a b
  compareKey (S.InlineSpineProp _) S.InlineRef = LT
  compareKey S.InlineRef (S.InlineSpineProp _) = GT
  compareKey S.InlineRef S.InlineRef = EQ

ascending :: forall a. KeyOrder a => Array a -> Boolean
ascending values = Array.all identity (Array.zipWith (\a b -> compareKey a b == LT) values (Array.drop 1 values))

ordered :: forall a. KeyOrder a => Array a -> Array a
ordered values = if ascending values then values else Array.sortBy compareKey values

instance (Ord k, KeyOrder k, Codec k, Codec v) => Codec (Map k v) where
  encode context value = do
    let entries = Map.toUnfoldable value :: Array (Tuple k v)
        sorted = if ascending (map fst entries) then entries else Array.sortBy (\a b -> compareKey (fst a) (fst b)) entries
    encoded <- encode context sorted
    internCtor context "Map" [ encoded ]
  decode context json = do
    args <- ctorArgs context "Map" json
    case args of
      [ values ] -> do
        entries <- decode context values :: Maybe (Array (Tuple k v))
        guard (ascending (map fst entries))
        pure (Map.fromFoldable entries)
      _ -> Nothing

instance (Ord a, KeyOrder a, Codec a) => Codec (Set a) where
  encode context value = do
    encoded <- encode context (ordered (Set.toUnfoldable value :: Array a))
    internCtor context "Set" [ encoded ]
  decode context json = do
    args <- ctorArgs context "Set" json
    case args of
      [ values ] -> do
        entries <- decode context values :: Maybe (Array a)
        guard (ascending entries)
        pure (Set.fromFoldable entries)
      _ -> Nothing

instance (RowToList row fields, Fields fields row) => Codec (Record row) where
  encode context value = do
    fields <- encodeFields (Proxy :: Proxy fields) context value
    intern context ("[\"record\",[" <> String.joinWith "," fields <> "]]")
  decode context json = payload context "record" json >>= toArray >>= decodeFields (Proxy :: Proxy fields) context

class Fields (fields :: RowList Type) (row :: Row Type) | fields -> row where
  encodeFields :: Proxy fields -> EncodeContext -> Record row -> Effect (Array String)
  decodeFields :: Proxy fields -> DecodeContext -> Array Json -> Maybe (Record row)

instance Fields Nil () where
  encodeFields _ _ _ = pure []
  decodeFields _ _ values = if Array.null values then Just {} else Nothing

instance (IsSymbol label, Codec a, Fields tail rest, Row.Cons label a rest row, Row.Lacks label rest) => Fields (Cons label a tail) row where
  encodeFields _ context record = do
    value <- encode context (Record.get label record)
    rest <- encodeFields (Proxy :: Proxy tail) context (Record.delete label record)
    pure (Array.cons (jsonArray [ quoted (stringUnits (reflectSymbol label)), value ]) rest)
    where label = Proxy :: Proxy label
  decodeFields _ context values = do
    { head, tail } <- Array.uncons values
    fields <- toArray head
    case fields of
      [ name, value ] | toString name == Just (stringUnits (reflectSymbol label)) -> do
        decoded <- decode context value
        rest <- decodeFields (Proxy :: Proxy tail) context tail
        pure (Record.insert label decoded rest)
      _ -> Nothing
    where label = Proxy :: Proxy label

ctor :: String -> Array String -> String
ctor tag args = "[\"ctor\",\"" <> tag <> "\",[" <> String.joinWith "," args <> "]]"

constructor :: DecodeContext -> Json -> Maybe { tag :: String, args :: Array Json }
constructor context json = do
  fields <- resolve context json >>= toArray
  case fields of
    [ kind, name, arguments ] | toString kind == Just "ctor" -> do
      tag <- toString name
      args <- toArray arguments
      pure { tag, args }
    _ -> Nothing

ctorArgs :: DecodeContext -> String -> Json -> Maybe (Array Json)
ctorArgs context tag json = do
  value <- constructor context json
  guard (tag == value.tag)
  pure value.args

genericEncode :: forall a rep. Generic a rep => Constructors rep => EncodeContext -> a -> Effect String
genericEncode context value = do
  result <- encodeConstructor context (from value)
  internCtor context result.tag result.args

genericDecode :: forall a rep. Generic a rep => Constructors rep => DecodeContext -> Json -> Maybe a
genericDecode context json = do
  value <- constructor context json
  to <$> decodeConstructor context value.tag value.args

class Constructors rep where
  encodeConstructor :: EncodeContext -> rep -> Effect { tag :: String, args :: Array String }
  decodeConstructor :: DecodeContext -> String -> Array Json -> Maybe rep

instance (Constructors a, Constructors b) => Constructors (Sum a b) where
  encodeConstructor context (Inl a) = encodeConstructor context a
  encodeConstructor context (Inr b) = encodeConstructor context b
  decodeConstructor context tag args = (Inl <$> decodeConstructor context tag args) <|> (Inr <$> decodeConstructor context tag args)

instance (IsSymbol name, Arguments args) => Constructors (Constructor name args) where
  encodeConstructor context (Constructor args) = do
    values <- encodeArguments context args
    pure { tag: reflectSymbol (Proxy :: Proxy name), args: values }
  decodeConstructor context tag args = do
    guard (tag == reflectSymbol (Proxy :: Proxy name))
    Constructor <$> decodeArguments context args

class Arguments rep where
  argumentCount :: Proxy rep -> Int
  encodeArguments :: EncodeContext -> rep -> Effect (Array String)
  decodeArguments :: DecodeContext -> Array Json -> Maybe rep

instance Arguments NoArguments where
  argumentCount _ = 0
  encodeArguments _ _ = pure []
  decodeArguments _ values = if Array.null values then Just NoArguments else Nothing

instance Codec a => Arguments (Argument a) where
  argumentCount _ = 1
  encodeArguments context (Argument a) = Array.singleton <$> encode context a
  decodeArguments context values = case values of
    [ value ] -> Argument <$> decode context value
    _ -> Nothing

instance (Arguments a, Arguments b) => Arguments (Product a b) where
  argumentCount _ = argumentCount (Proxy :: Proxy a) + argumentCount (Proxy :: Proxy b)
  encodeArguments context (Product a b) = append <$> encodeArguments context a <*> encodeArguments context b
  decodeArguments context values = Product
    <$> decodeArguments context (Array.take count values)
    <*> decodeArguments context (Array.drop count values)
    where count = argumentCount (Proxy :: Proxy a)

-- This explicit list is the versioned business schema: adding a constructor or
-- changing a record changes its portable representation and compiler identity.
instance Codec C.Ident where
  encode context (C.Ident value) = one context "Ident" value
  decode context json = do
    args <- ctorArgs context "Ident" json
    case args of
      [ value ] -> C.Ident <$> decode context value
      _ -> Nothing
instance Codec C.ModuleName where
  encode context (C.ModuleName value) = one context "ModuleName" value
  decode context json = do
    args <- ctorArgs context "ModuleName" json
    case args of
      [ value ] -> C.ModuleName <$> decode context value
      _ -> Nothing
instance Codec C.ProperName where
  encode context (C.ProperName value) = one context "ProperName" value
  decode context json = do
    args <- ctorArgs context "ProperName" json
    case args of
      [ value ] -> C.ProperName <$> decode context value
      _ -> Nothing
instance Codec a => Codec (C.Qualified a) where
  encode context (C.Qualified name value) = two context "Qualified" name value
  decode context json = do
    args <- ctorArgs context "Qualified" json
    case args of
      [ name, value ] -> C.Qualified <$> decode context name <*> decode context value
      _ -> Nothing
instance Codec C.Ann where
  encode = genericEncode
  decode = genericDecode
instance Codec C.SourceBindingId where
  encode = genericEncode
  decode = genericDecode
instance Codec C.ExprType where
  encode context value = do
    known <- TypeMemo.lookup context.typeMemo value
    index <- if known >= 0 then pure known else do
      -- Memoize physical subtrees before revisiting shared children. Equal but
      -- distinct types reproduce existing graph nodes and the same type index;
      -- they do not add entries or change either table's encounter order.
      encoded <- encodeType context value
      let TypeTable keys = context.types
      index <- Pool.intern keys encoded
      TypeMemo.insert context.typeMemo value index
      pure index
    encoded <- encode context index
    internCtor context "TypeRef" [ encoded ]
  decode context value = do
    args <- ctorArgs context "TypeRef" value
    case args of
      [ raw ] -> do
        index <- decode context raw
        Map.lookup index context.types
      _ -> Nothing

encodeType :: EncodeContext -> C.ExprType -> Effect String
encodeType context = case _ of
  C.Int -> internCtor context "Int" []
  C.Number -> internCtor context "Number" []
  C.String -> internCtor context "String" []
  C.Char -> internCtor context "Char" []
  C.Boolean -> internCtor context "Boolean" []
  C.Unit -> internCtor context "Unit" []
  C.Any -> internCtor context "Any" []
  C.TypeLevelString a -> one context "TypeLevelString" a
  C.Array a -> one context "Array" a
  C.TypeVar a -> one context "TypeVar" a
  C.ADT a b c -> three context "ADT" a b c
  C.TypeApp a b -> two context "TypeApp" a b
  C.Func a b -> two context "Func" a b
  C.Row a b -> two context "Row" a b
  C.Record a -> one context "Record" a
  C.ForAll a b -> two context "ForAll" a b
  C.ConstrainedType a b -> two context "ConstrainedType" a b

decodeType :: DecodeContext -> Json -> Maybe C.ExprType
decodeType context json = do
  value <- constructor context json
  case value.tag, value.args of
    "Int", [] -> Just C.Int
    "Number", [] -> Just C.Number
    "String", [] -> Just C.String
    "Char", [] -> Just C.Char
    "Boolean", [] -> Just C.Boolean
    "Unit", [] -> Just C.Unit
    "Any", [] -> Just C.Any
    "TypeLevelString", [ a ] -> C.TypeLevelString <$> decode context a
    "Array", [ a ] -> C.Array <$> decode context a
    "TypeVar", [ a ] -> C.TypeVar <$> decode context a
    "ADT", [ a, b, c ] -> C.ADT <$> decode context a <*> decode context b <*> decode context c
    "TypeApp", [ a, b ] -> C.TypeApp <$> decode context a <*> decode context b
    "Func", [ a, b ] -> C.Func <$> decode context a <*> decode context b
    "Row", [ a, b ] -> C.Row <$> decode context a <*> decode context b
    "Record", [ a ] -> C.Record <$> decode context a
    "ForAll", [ a, b ] -> C.ForAll <$> decode context a <*> decode context b
    "ConstrainedType", [ a, b ] -> C.ConstrainedType <$> decode context a <*> decode context b
    _, _ -> Nothing

instance Codec C.Meta where
  encode = genericEncode
  decode = genericDecode
instance Codec C.ConstructorType where
  encode = genericEncode
  decode = genericDecode
instance Codec C.Comment where
  encode = genericEncode
  decode = genericDecode
instance Codec a => Codec (C.Module a) where
  encode = genericEncode
  decode = genericDecode
instance Codec a => Codec (C.Import a) where
  encode = genericEncode
  decode = genericDecode
instance Codec C.ReExport where
  encode = genericEncode
  decode = genericDecode
instance Codec a => Codec (C.Bind a) where
  encode context = case _ of
    C.NonRec a -> one context "NonRec" a
    C.Rec a -> one context "Rec" a
  decode context json = do
    value <- constructor context json
    case value.tag, value.args of
      "NonRec", [ a ] -> C.NonRec <$> decode context a
      "Rec", [ a ] -> C.Rec <$> decode context a
      _, _ -> Nothing
instance Codec a => Codec (C.Binding a) where
  encode context (C.Binding a b c) = three context "Binding" a b c
  decode context json = do
    args <- ctorArgs context "Binding" json
    case args of
      [ a, b, c ] -> C.Binding <$> decode context a <*> decode context b <*> decode context c
      _ -> Nothing
instance Codec a => Codec (C.Expr a) where
  encode context = case _ of
    C.ExprVar a b -> two context "ExprVar" a b
    C.ExprLit a b -> two context "ExprLit" a b
    C.ExprConstructor a b c d -> four context "ExprConstructor" a b c d
    C.ExprAccessor a b c -> three context "ExprAccessor" a b c
    C.ExprUpdate a b c -> three context "ExprUpdate" a b c
    C.ExprAbs a b c -> three context "ExprAbs" a b c
    C.ExprApp a b c -> three context "ExprApp" a b c
    C.ExprCase a b c -> three context "ExprCase" a b c
    C.ExprLet a b c -> three context "ExprLet" a b c
    C.ExprTypeApp a b c -> three context "ExprTypeApp" a b c
  decode context json = do
    value <- constructor context json
    case value.tag, value.args of
      "ExprVar", [ a, b ] -> C.ExprVar <$> decode context a <*> decode context b
      "ExprLit", [ a, b ] -> C.ExprLit <$> decode context a <*> decode context b
      "ExprConstructor", [ a, b, c, d ] -> C.ExprConstructor <$> decode context a <*> decode context b <*> decode context c <*> decode context d
      "ExprAccessor", [ a, b, c ] -> C.ExprAccessor <$> decode context a <*> decode context b <*> decode context c
      "ExprUpdate", [ a, b, c ] -> C.ExprUpdate <$> decode context a <*> decode context b <*> decode context c
      "ExprAbs", [ a, b, c ] -> C.ExprAbs <$> decode context a <*> decode context b <*> decode context c
      "ExprApp", [ a, b, c ] -> C.ExprApp <$> decode context a <*> decode context b <*> decode context c
      "ExprCase", [ a, b, c ] -> C.ExprCase <$> decode context a <*> decode context b <*> decode context c
      "ExprLet", [ a, b, c ] -> C.ExprLet <$> decode context a <*> decode context b <*> decode context c
      "ExprTypeApp", [ a, b, c ] -> C.ExprTypeApp <$> decode context a <*> decode context b <*> decode context c
      _, _ -> Nothing
instance Codec a => Codec (C.CaseAlternative a) where
  encode context (C.CaseAlternative a b) = two context "CaseAlternative" a b
  decode context json = do
    args <- ctorArgs context "CaseAlternative" json
    case args of
      [ a, b ] -> C.CaseAlternative <$> decode context a <*> decode context b
      _ -> Nothing
instance Codec a => Codec (C.CaseGuard a) where
  encode context = case _ of
    C.Unconditional a -> one context "Unconditional" a
    C.Guarded a -> one context "Guarded" a
  decode context json = do
    value <- constructor context json
    case value.tag, value.args of
      "Unconditional", [ a ] -> C.Unconditional <$> decode context a
      "Guarded", [ a ] -> C.Guarded <$> decode context a
      _, _ -> Nothing
instance Codec a => Codec (C.Guard a) where
  encode context (C.Guard a b) = two context "Guard" a b
  decode context json = do
    args <- ctorArgs context "Guard" json
    case args of
      [ a, b ] -> C.Guard <$> decode context a <*> decode context b
      _ -> Nothing
instance Codec a => Codec (C.Prop a) where
  encode context (C.Prop name value) = two context "Prop" name value
  decode context json = do
    args <- ctorArgs context "Prop" json
    case args of
      [ name, value ] -> C.Prop <$> decode context name <*> decode context value
      _ -> Nothing
instance Codec a => Codec (C.Literal a) where
  encode context = case _ of
    C.LitInt value -> one context "LitInt" value
    C.LitNumber value -> one context "LitNumber" value
    C.LitString value -> one context "LitString" value
    C.LitChar value -> one context "LitChar" value
    C.LitBoolean value -> one context "LitBoolean" value
    C.LitArray value -> one context "LitArray" value
    C.LitRecord value -> one context "LitRecord" value
  decode context json = do
    value <- constructor context json
    case value.tag, value.args of
      "LitInt", [ a ] -> C.LitInt <$> decode context a
      "LitNumber", [ a ] -> C.LitNumber <$> decode context a
      "LitString", [ a ] -> C.LitString <$> decode context a
      "LitChar", [ a ] -> C.LitChar <$> decode context a
      "LitBoolean", [ a ] -> C.LitBoolean <$> decode context a
      "LitArray", [ a ] -> C.LitArray <$> decode context a
      "LitRecord", [ a ] -> C.LitRecord <$> decode context a
      _, _ -> Nothing
instance Codec a => Codec (C.Binder a) where
  encode context = case _ of
    C.BinderNull a -> one context "BinderNull" a
    C.BinderVar a b -> two context "BinderVar" a b
    C.BinderNamed a b c -> three context "BinderNamed" a b c
    C.BinderLit a b -> two context "BinderLit" a b
    C.BinderConstructor a b c d -> four context "BinderConstructor" a b c d
  decode context json = do
    value <- constructor context json
    case value.tag, value.args of
      "BinderNull", [ a ] -> C.BinderNull <$> decode context a
      "BinderVar", [ a, b ] -> C.BinderVar <$> decode context a <*> decode context b
      "BinderNamed", [ a, b, c ] -> C.BinderNamed <$> decode context a <*> decode context b <*> decode context c
      "BinderLit", [ a, b ] -> C.BinderLit <$> decode context a <*> decode context b
      "BinderConstructor", [ a, b, c, d ] -> C.BinderConstructor <$> decode context a <*> decode context b <*> decode context c <*> decode context d
      _, _ -> Nothing
instance Codec a => Codec (B.BackendSyntax a) where
  encode context value = encodeBackendSyntax context value
  decode context json = decodeBackendSyntax context json

-- The recursive NeutralExpr paths call these functions directly. Building the
-- complete BackendSyntax dictionary at each node eagerly constructs codecs for
-- every branch (including Generic operators/effects), even for a plain Var.
encodeBackendSyntax :: forall a. Codec a => EncodeContext -> B.BackendSyntax a -> Effect String
encodeBackendSyntax context = case _ of
  B.Var a -> one context "Var" a
  B.Local a b -> two context "Local" a b
  B.Lit a -> one context "Lit" a
  B.App a b -> two context "App" a b
  B.TypeApp a b -> two context "TypeApp" a b
  B.Abs a b -> two context "Abs" a b
  B.UncurriedApp a b -> two context "UncurriedApp" a b
  B.UncurriedAbs a b -> two context "UncurriedAbs" a b
  B.UncurriedEffectApp a b -> two context "UncurriedEffectApp" a b
  B.UncurriedEffectAbs a b -> two context "UncurriedEffectAbs" a b
  B.Accessor a b -> two context "Accessor" a b
  B.Update a b -> two context "Update" a b
  B.CtorSaturated a b c d e -> five context "CtorSaturated" a b c d e
  B.CtorDef a b c d -> four context "CtorDef" a b c d
  B.LetRec a b c -> three context "LetRec" a b c
  B.Let a b c d -> four context "Let" a b c d
  B.EffectBind a b c d -> four context "EffectBind" a b c d
  B.EffectPure a -> one context "EffectPure" a
  B.EffectDefer a -> one context "EffectDefer" a
  B.Branch a b -> two context "Branch" a b
  B.PrimOp a -> one context "PrimOp" a
  B.PrimEffect a -> one context "PrimEffect" a
  B.PrimUndefined -> internCtor context "PrimUndefined" []
  B.Fail a -> one context "Fail" a
  B.Typed a b -> two context "Typed" a b

decodeBackendSyntax :: forall a. Codec a => DecodeContext -> Json -> Maybe (B.BackendSyntax a)
decodeBackendSyntax context json = do
  value <- constructor context json
  case value.tag, value.args of
    "Var", [ a ] -> B.Var <$> decode context a
    "Local", [ a, b ] -> B.Local <$> decode context a <*> decode context b
    "Lit", [ a ] -> B.Lit <$> decode context a
    "App", [ a, b ] -> B.App <$> decode context a <*> decode context b
    "TypeApp", [ a, b ] -> B.TypeApp <$> decode context a <*> decode context b
    "Abs", [ a, b ] -> B.Abs <$> decode context a <*> decode context b
    "UncurriedApp", [ a, b ] -> B.UncurriedApp <$> decode context a <*> decode context b
    "UncurriedAbs", [ a, b ] -> B.UncurriedAbs <$> decode context a <*> decode context b
    "UncurriedEffectApp", [ a, b ] -> B.UncurriedEffectApp <$> decode context a <*> decode context b
    "UncurriedEffectAbs", [ a, b ] -> B.UncurriedEffectAbs <$> decode context a <*> decode context b
    "Accessor", [ a, b ] -> B.Accessor <$> decode context a <*> decode context b
    "Update", [ a, b ] -> B.Update <$> decode context a <*> decode context b
    "CtorSaturated", [ a, b, c, d, e ] -> B.CtorSaturated <$> decode context a <*> decode context b <*> decode context c <*> decode context d <*> decode context e
    "CtorDef", [ a, b, c, d ] -> B.CtorDef <$> decode context a <*> decode context b <*> decode context c <*> decode context d
    "LetRec", [ a, b, c ] -> B.LetRec <$> decode context a <*> decode context b <*> decode context c
    "Let", [ a, b, c, d ] -> B.Let <$> decode context a <*> decode context b <*> decode context c <*> decode context d
    "EffectBind", [ a, b, c, d ] -> B.EffectBind <$> decode context a <*> decode context b <*> decode context c <*> decode context d
    "EffectPure", [ a ] -> B.EffectPure <$> decode context a
    "EffectDefer", [ a ] -> B.EffectDefer <$> decode context a
    "Branch", [ a, b ] -> B.Branch <$> decode context a <*> decode context b
    "PrimOp", [ a ] -> B.PrimOp <$> decode context a
    "PrimEffect", [ a ] -> B.PrimEffect <$> decode context a
    "PrimUndefined", [] -> Just B.PrimUndefined
    "Fail", [ a ] -> B.Fail <$> decode context a
    "Typed", [ a, b ] -> B.Typed <$> decode context a <*> decode context b
    _, _ -> Nothing

instance Codec B.Level where
  encode = genericEncode
  decode = genericDecode
instance Codec a => Codec (B.Pair a) where
  encode = genericEncode
  decode = genericDecode
instance Codec B.BackendAccessor where
  encode = genericEncode
  decode = genericDecode
instance Codec a => Codec (B.BackendOperator a) where
  encode = genericEncode
  decode = genericDecode
instance Codec B.BackendOperator1 where
  encode = genericEncode
  decode = genericDecode
instance Codec B.BackendOperator2 where
  encode = genericEncode
  decode = genericDecode
instance Codec B.BackendOperatorNum where
  encode = genericEncode
  decode = genericDecode
instance Codec B.BackendOperatorOrd where
  encode = genericEncode
  decode = genericDecode
instance Codec a => Codec (B.BackendEffect a) where
  encode = genericEncode
  decode = genericDecode
instance Codec A.Capture where
  encode = genericEncode
  decode = genericDecode
instance Codec A.Usage where
  encode = genericEncode
  decode = genericDecode
instance Codec A.Complexity where
  encode = genericEncode
  decode = genericDecode
instance Codec A.ResultTerm where
  encode = genericEncode
  decode = genericDecode
instance Codec A.BackendAnalysis where
  encode = genericEncode
  decode = genericDecode
instance Codec S.ExternImpl where
  encode = genericEncode
  decode = genericDecode
instance Codec S.EvalRef where
  encode = genericEncode
  decode = genericDecode
instance Codec S.InlineAccessor where
  encode = genericEncode
  decode = genericDecode
instance Codec S.InlineDirective where
  encode = genericEncode
  decode = genericDecode
instance Codec S.NeutralExpr where
  encode context (S.NeutralExpr value) = do
    encoded <- encodeBackendSyntax context value
    internCtor context "NeutralExpr" [ encoded ]
  decode context json = do
    index <- reference json
    Map.lookup index context.expressions
instance Codec GoType where
  encode value = genericEncode value
  decode value = genericDecode value
instance Codec GoReference where
  encode = genericEncode
  decode = genericDecode
