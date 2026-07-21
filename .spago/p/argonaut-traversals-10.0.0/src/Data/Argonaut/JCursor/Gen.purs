module Data.Argonaut.JCursor.Gen where

import Prelude

import Control.Lazy (class Lazy, defer)
import Control.Monad.Gen (class MonadGen)
import Control.Monad.Gen as Gen
import Control.Monad.Rec.Class (class MonadRec)

import Data.Argonaut.JCursor (JCursor(..))
import Data.String.Gen (genUnicodeString)

genJCursor :: forall m. MonadGen m => MonadRec m => Lazy (m JCursor) => m JCursor
genJCursor = Gen.resize (min 10) $ Gen.sized genJCursor'
  where
  genJCursor' size
    | size > 0 = Gen.resize (_ - 1) (Gen.choose genField genIndex)
    | otherwise = pure JCursorTop
  genField = JField <$> genUnicodeString <*> defer \_ -> genJCursor
  genIndex = JIndex <$> Gen.chooseInt 0 1000 <*> defer \_ -> genJCursor
