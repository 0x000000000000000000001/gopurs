module Gopurs.DecoderSchemas.Types
  ( Schema(..)
  , Program(..)
  , weight
  , tagged
  , complete
  , textComplete
  ) where

import Prelude

import Data.Foldable (all, foldl)
import Data.Tuple (Tuple(..))
import PureScript.Backend.Optimizer.Semantics (NeutralExpr)
import PureScript.Backend.Optimizer.Syntax (Level)

-- Custom leaves retain ordinary dispatch. Derived leaves carry a proof of a
-- real decoder body; neither one supplies the runtime tag required at a root.
data Schema
  = Scalar String
  | Sequence Schema
  | Optional Schema
  | ObjectSchema (Array (Tuple String Schema))
  | Custom
  | Derived (Program Schema)

-- Every read, branch and constructor comes from the source body. Error
-- continuations have already been proved to forward the exact Left payload.
data Program a
  = ReadField Level String Boolean Boolean a (Program a)
  | Choices (Array (Tuple (Tuple Level String) (Program a))) (Program a)
  | ReturnValue NeutralExpr
  | TypeMismatch String

weight :: Schema -> Int
weight = case _ of
  Scalar _ -> 1
  Custom -> 0
  Derived program -> 1 + programWeight program
  Sequence inner -> 1 + weight inner
  Optional inner -> 1 + weight inner
  ObjectSchema fields -> 1 + foldl (\n (Tuple _ child) -> n + weight child) 0 fields

programWeight :: Program Schema -> Int
programWeight = case _ of
  ReadField _ _ _ _ schema next -> 1 + weight schema + programWeight next
  Choices cases other -> 1 + foldl (\n (Tuple _ branch) -> n + programWeight branch) (programWeight other) cases
  _ -> 1

tagged :: Schema -> Boolean
tagged = case _ of
  Custom -> false
  Derived _ -> false
  _ -> true

-- Reads inside a derived program need fully standard schemas. In particular,
-- a derived decoder cannot recursively justify another derived field read.
complete :: Schema -> Boolean
complete = case _ of
  Custom -> false
  Derived _ -> false
  Sequence child -> complete child
  Optional child -> complete child
  ObjectSchema fields -> all (\(Tuple _ child) -> complete child) fields
  _ -> true

-- An opaque callback may mutate or retain its Json argument. Materializing only
-- its subtree would lose aliases across repeated reads, so text specialization
-- requires a whole-schema proof. Derived programs already prove all their reads.
textComplete :: Schema -> Boolean
textComplete = case _ of
  Custom -> false
  Sequence child -> textComplete child
  Optional child -> textComplete child
  ObjectSchema fields -> all (\(Tuple _ child) -> textComplete child) fields
  _ -> true
