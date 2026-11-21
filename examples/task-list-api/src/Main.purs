module Main where

import Prelude

import Control.Monad.Except (runExcept)
import Data.Argonaut.Decode (JsonDecodeError, fromJsonString)
import Data.Argonaut.Encode (class EncodeJson, toJsonString)
import Data.Array as Array
import Data.Either (Either(..))
import Data.Maybe (Maybe(..), fromMaybe)
import Data.String (Pattern(..), split, trim)
import Data.String.CodeUnits as String
import Data.Traversable (traverse)
import Database.SQLite3 as SQLite
import Effect (Effect)
import Effect.Aff (Aff, runAff_, throwError)
import Effect.Class (liftEffect)
import Effect.Console as Console
import Effect.Exception (error, message)
import Example.HTTP (Request, Response, serve)
import Foreign (Foreign, readInt, readString, unsafeToForeign)
import Node.Process as Process

type Task = { id :: String, title :: String, completed :: Boolean }
type NewTask = { title :: String }
type UpdateTask = { completed :: Boolean }

main :: Effect Unit
main = runAff_ finish do
  address <- liftEffect $ fromMaybe "127.0.0.1:8080" <$> Process.lookupEnv "TASK_LIST_ADDRESS"
  path <- liftEffect $ fromMaybe "tasks.db" <$> Process.lookupEnv "TASK_LIST_DATABASE"
  SQLite.withDatabase path \database -> do
    void $ SQLite.execute database
      "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT NOT NULL, completed INTEGER NOT NULL DEFAULT 0 CHECK (completed IN (0, 1)))" []
    serve address \request reply -> runAff_ (respond reply) (route database request)
  liftEffect $ Console.log "STOPPED"
  where
  finish (Left err) = Console.error (message err) *> Process.exit' 1
  finish (Right _) = pure unit
  respond reply (Right response) = reply response
  respond reply (Left err) = do
    Console.error (message err)
    reply $ failure 500 "Internal server error"

json :: forall a. EncodeJson a => Int -> a -> Response
json status value = { status, body: toJsonString value }

failure :: Int -> String -> Response
failure status text = json status { error: text }

route :: SQLite.Database -> Request -> Aff Response
route database request = case request.method, split (Pattern "/") request.path of
  "GET", [ "", "health" ] -> pure $ json 200 { status: "ok" }
  "GET", [ "", "tasks" ] -> json 200 <$> tasks database
  "POST", [ "", "tasks" ] -> case fromJsonString request.body :: Either JsonDecodeError NewTask of
    Left _ -> pure $ failure 400 "Expected a JSON object with a string title"
    Right input -> do
      let title = trim input.title
      if title == "" || String.length title > 200 then
        pure $ failure 400 "Title must contain 1 to 200 UTF-16 code units"
      else do
        result <- SQLite.execute database "INSERT INTO tasks (title) VALUES (?)" [ unsafeToForeign title ]
        pure $ json 201 { id: result.lastInsertRowId, title, completed: false }
  method, [ "", "tasks", id ] | validId id -> case method of
    "GET" -> getTask database id
    "PATCH" -> case fromJsonString request.body :: Either JsonDecodeError UpdateTask of
      Left _ -> pure $ failure 400 "Expected a JSON object with a boolean completed"
      Right input -> do
        result <- SQLite.execute database "UPDATE tasks SET completed = ? WHERE id = ?"
          [ unsafeToForeign (if input.completed then 1 else 0), unsafeToForeign id ]
        if result.changes == 0 then pure $ failure 404 "Task not found"
        else getTask database id
    "DELETE" -> do
      result <- SQLite.execute database "DELETE FROM tasks WHERE id = ?" [ unsafeToForeign id ]
      pure $ if result.changes == 0 then failure 404 "Task not found" else { status: 204, body: "" }
    _ -> pure $ failure 405 "Method not allowed"
  _, _ -> pure $ failure 404 "Route not found"

validId :: String -> Boolean
validId value = value /= "" && String.length value <= 19 &&
  Array.all (\char -> char >= '0' && char <= '9') (String.toCharArray value)

selectTasks :: String
selectTasks = "SELECT CAST(id AS TEXT), title, completed FROM tasks"

tasks :: SQLite.Database -> Aff (Array Task)
tasks database = do
  result <- SQLite.query database (selectTasks <> " ORDER BY id") []
  traverse decodeTask result.rows

getTask :: SQLite.Database -> String -> Aff Response
getTask database id = do
  result <- SQLite.query database (selectTasks <> " WHERE id = ?") [ unsafeToForeign id ]
  case Array.head result.rows of
    Nothing -> pure $ failure 404 "Task not found"
    Just row -> json 200 <$> decodeTask row

decodeTask :: Array Foreign -> Aff Task
decodeTask [ rawId, rawTitle, rawCompleted ] = case runExcept do
  id <- readString rawId
  title <- readString rawTitle
  completed <- (_ /= 0) <$> readInt rawCompleted
  pure { id, title, completed } of
    Left errors -> throwError $ error $ "Unexpected SQLite row: " <> show errors
    Right task -> pure task
decodeTask _ = throwError $ error "Unexpected SQLite column count"
