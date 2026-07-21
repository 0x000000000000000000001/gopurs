module Node.Buffer.Constants where

import Effect (Effect)

foreign import inspectMaxBytes :: Effect Int

foreign import maxLength :: Int

foreign import maxStringLength :: Int
