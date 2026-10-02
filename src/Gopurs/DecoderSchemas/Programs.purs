module Gopurs.DecoderSchemas.Programs (DictionaryProofs, deriveProgram) where

import Prelude hiding (one)

import Control.Alternative (guard)
import Data.Array as Array
import Data.Array.NonEmpty as NEA
import Data.Foldable (all)
import Data.Map as Map
import Data.Maybe (Maybe(..))
import Data.String as String
import Data.String.CodeUnits as CU
import Data.String.Pattern (Pattern(..))
import Data.Traversable (traverse)
import Data.Tuple (Tuple(..))
import Gopurs.DecoderSchemas.Source (standard, strip)
import Gopurs.DecoderSchemas.Types (Program(..), Schema, complete)
import PureScript.Backend.Optimizer.CoreFn (Ident(..), Literal(..), ModuleName(..), Prop(..), Qualified(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr(..))
import PureScript.Backend.Optimizer.Syntax (BackendAccessor(..), BackendOperator(..), BackendOperator1(..), BackendOperator2(..), BackendOperatorOrd(..), BackendSyntax(..), Level, Pair(..))

-- Track provenance separately from the result envelope. A successful object
-- borrow, a decoded field and an error payload have distinct permitted uses.
data Atom = Input | ObjectInput | Decoded Level | Failure Level | Result Boolean Atom

type DictionaryProofs =
  { resolveDictionary :: NeutralExpr -> Maybe (Tuple NeutralExpr (Array NeutralExpr))
  , readSchema :: NeutralExpr -> Maybe Schema
  }

deriveProgram :: DictionaryProofs -> NeutralExpr -> Maybe (Program Schema)
deriveProgram proofs expr = do
  method <- case strip expr of
    Lit (LitRecord [ Prop "decodeJson" value ]) -> Just value
    _ -> Nothing
  { refs, body } <- case strip method of
    Abs refs body -> Just { refs, body }
    _ -> Nothing
  Tuple _ input <- one (NEA.toArray refs)
  { result, producer, continuation } <- case strip body of
    Let _ result producer continuation -> Just { result, producer, continuation }
    _ -> Nothing
  { fn: helper, args } <- callExpr producer
  guard (isGlobal "Data.Argonaut.Decode.Internal.Record" "borrowObject" helper)
  { first: objectMethod, second: json } <- two args
  input' <- case strip json of
    Local _ level -> Just level
    _ -> Nothing
  guard (input == input')
  dictionary <- case strip objectMethod of
    Accessor dictionary (GetProp "decodeJson") -> Just dictionary
    _ -> Nothing
  Tuple head arguments <- proofs.resolveDictionary dictionary
  guard (standard "decodeForeignObject" head)
  inner <- one arguments
  guard (standard "decodeJsonJson" inner)
  let env = Map.singleton input Input
  guard (forwards 128 result (Map.insert result (Result false (Failure result)) env) continuation)
  lower 256 (Map.insert result (Result true ObjectInput) env) continuation
  where
  lower remaining env original
    | remaining <= 0 = Nothing
    | otherwise = case strip original of
        Let _ level value body -> case atom env value of
          Just bound -> lower (remaining - 1) (Map.insert level bound env) body
          Nothing -> do
            -- A shadowing producer can alias an earlier result. Decline this
            -- shape instead of conflating its generated local with that value.
            guard (not (Map.member level env))
            { fn, args } <- callExpr value
            Tuple optional nullable <- reader fn
            { first: method, second: object, third: label } <- three args
            guard (case atom env object of
              Just ObjectInput -> true
              _ -> false)
            key <- stringExpr label
            { fn: decode, args: dictionaries } <- callExpr method
            guard (standard "decodeJson" decode)
            dictionary <- one dictionaries
            schema <- proofs.readSchema dictionary
            guard (complete schema)
            guard (forwards 128 level (Map.insert level (Result false (Failure level)) env) body)
            next <- lower (remaining - 1) (Map.insert level (Result true (Decoded level)) env) body
            pure (ReadField level key optional nullable schema next)
        Branch cases other -> case knownBranch env (NEA.toArray cases) other of
          Just body -> lower (remaining - 1) env body
          Nothing -> do
            branches <- traverse (\(Pair condition body) -> do
              { left, right } <- case strip condition of
                PrimOp (Op2 (OpStringOrd OpEq) left right) -> Just { left, right }
                _ -> Nothing
              level <- atom env left >>= decodedAtom
              key <- stringExpr right
              Tuple (Tuple level key) <$> lower (remaining - 1) env body) (NEA.toArray cases)
            Choices branches <$> lower (remaining - 1) env other
        CtorSaturated ctor _ _ _ [ Tuple "value0" value ] | isEither "Right" ctor -> ReturnValue <$> valueExpr env value
        CtorSaturated ctor _ _ _ [ Tuple "value0" value ] | isEither "Left" ctor -> do
          { errorCtor, label } <- case strip value of
            CtorSaturated errorCtor _ _ _ [ Tuple "value0" label ] -> Just { errorCtor, label }
            _ -> Nothing
          guard (errorCtor == Qualified (Just (ModuleName "Data.Argonaut.Decode.Error")) (Ident "TypeMismatch"))
          message <- stringExpr label
          pure (TypeMismatch message)
        _ -> Nothing

atom :: Map.Map Level Atom -> NeutralExpr -> Maybe Atom
atom env expr = case strip expr of
  Local _ level -> Map.lookup level env
  Accessor value (GetCtorField ctor _ _ _ "value0" 0) -> do
    { success, payload } <- atom env value >>= resultAtom
    guard (isEither (if success then "Right" else "Left") ctor)
    pure payload
  _ -> Nothing

knownBranch :: Map.Map Level Atom -> Array (Pair NeutralExpr) -> NeutralExpr -> Maybe NeutralExpr
knownBranch env cases other = case Array.uncons cases of
  Nothing -> Just other
  Just { head: Pair condition body, tail } -> case strip condition of
    PrimOp (Op1 (OpIsTag ctor) value) -> do
      { success } <- atom env value >>= resultAtom
      guard (isEither "Right" ctor || isEither "Left" ctor)
      if success == isEither "Right" ctor then Just body else knownBranch env tail other
    _ -> Nothing

forwards :: Int -> Level -> Map.Map Level Atom -> NeutralExpr -> Boolean
forwards fuel expected env expr
  | fuel <= 0 = false
  | otherwise = case strip expr of
      Let _ level value body -> case atom env value of
        Just bound -> forwards (fuel - 1) expected (Map.insert level bound env) body
        _ -> false
      Branch cases other -> case knownBranch env (NEA.toArray cases) other of
        Just body -> forwards (fuel - 1) expected env body
        _ -> false
      CtorSaturated ctor _ _ _ [ Tuple "value0" value ] | isEither "Left" ctor -> case atom env value of
        Just (Failure level) -> level == expected
        _ -> false
      _ -> case atom env expr of
        Just (Result false (Failure level)) -> level == expected
        _ -> false

-- Successful construction is limited to literals, decoded values and real ADT
-- constructors. The ordinary code generator retains their annotations, native
-- field layouts, reboxing and constructor identities.
valueExpr :: Map.Map Level Atom -> NeutralExpr -> Maybe NeutralExpr
valueExpr env original@(NeutralExpr syn) = case syn of
  Typed ty inner -> NeutralExpr <<< Typed ty <$> valueExpr env inner
  TypeApp inner ty -> NeutralExpr <<< flip TypeApp ty <$> valueExpr env inner
  Lit (LitString _) -> Just original
  Lit (LitInt _) -> Just original
  Lit (LitNumber _) -> Just original
  Lit (LitBoolean _) -> Just original
  CtorSaturated ctor ct tn cn fields -> NeutralExpr <<< CtorSaturated ctor ct tn cn <$> traverse (\(Tuple key value) -> Tuple key <$> valueExpr env value) fields
  _ -> do
    level <- atom env original >>= decodedAtom
    pure (NeutralExpr (Local Nothing level))

reader :: NeutralExpr -> Maybe (Tuple Boolean Boolean)
reader expr = case strip expr of
  Var (Qualified (Just (ModuleName "Data.Argonaut.Decode.Decoders")) (Ident name)) -> do
    base <- Array.find (\base -> name == base || case String.stripPrefix (Pattern (base <> "__")) name of
      Just suffix -> suffix /= "" && all (\c -> c >= '0' && c <= '9') (CU.toCharArray suffix)
      _ -> false) [ "getField", "getFieldOptional", "getFieldOptional'" ]
    pure (Tuple (base /= "getField") (base == "getFieldOptional'"))
  _ -> Nothing

isGlobal :: String -> String -> NeutralExpr -> Boolean
isGlobal owner name expr = case strip expr of
  Var (Qualified (Just (ModuleName m)) (Ident n)) -> m == owner && n == name
  _ -> false

isEither :: String -> Qualified Ident -> Boolean
isEither name = (_ == Qualified (Just (ModuleName "Data.Either")) (Ident name))

resultAtom :: Atom -> Maybe { success :: Boolean, payload :: Atom }
resultAtom = case _ of
  Result success payload -> Just { success, payload }
  _ -> Nothing

decodedAtom :: Atom -> Maybe Level
decodedAtom = case _ of
  Decoded level -> Just level
  _ -> Nothing

callExpr :: NeutralExpr -> Maybe { fn :: NeutralExpr, args :: Array NeutralExpr }
callExpr expr = case strip expr of
  App fn args -> Just { fn, args: NEA.toArray args }
  _ -> Nothing

stringExpr :: NeutralExpr -> Maybe String
stringExpr expr = case strip expr of
  Lit (LitString value) -> Just value
  _ -> Nothing

one :: forall a. Array a -> Maybe a
one = case _ of
  [ value ] -> Just value
  _ -> Nothing

two :: forall a. Array a -> Maybe { first :: a, second :: a }
two = case _ of
  [ first, second ] -> Just { first, second }
  _ -> Nothing

three :: forall a. Array a -> Maybe { first :: a, second :: a, third :: a }
three = case _ of
  [ first, second, third ] -> Just { first, second, third }
  _ -> Nothing
