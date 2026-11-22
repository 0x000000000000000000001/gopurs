package purescript

import "gopurs/output/gopurs_runtime"

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"

	_ "modernc.org/sqlite"
)

type Database struct {
	db           *sql.DB
	anchor       *sql.Conn
	mu           sync.Mutex
	closed       bool
	transactions map[*Transaction]struct{}
	closeOnce    sync.Once
	closeErr     error
}

type Transaction struct {
	owner  *Database
	conn   *sql.Conn
	tx     *sql.Tx
	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.Mutex
	done   bool
}

type ExecResult struct {
	Changes         int64
	LastInsertRowId string
}
type QueryResult struct {
	Columns []string
	Rows    [][]any
}

var memoryID atomic.Uint64

func Database_SQLite3_OpenImpl(ctx context.Context, path string) (*Database, func(), error) {
	memory := path == ":memory:"
	if memory {
		path = "file:gopurs-memory-" + strconv.FormatUint(memoryID.Add(1), 10) + "?mode=memory&cache=shared"
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, nil, err
	}
	// One operational connection; transactions reserve it until completion.
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	result := &Database{db: db, transactions: make(map[*Transaction]struct{})}
	cleanup := func() { _ = Database_SQLite3_CloseImpl(result) }
	if memory {
		// Interrupted SQLite operations can invalidate their physical connection.
		// An idle anchor keeps a private shared-memory database alive when the
		// operational connection is replaced after cancellation.
		db.SetMaxOpenConns(2)
		result.anchor, err = db.Conn(ctx)
		if err != nil {
			return nil, cleanup, err
		}
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, cleanup, err
	}
	return result, cleanup, nil
}

func Database_SQLite3_CloseImpl(db *Database) error {
	if db == nil {
		return errors.New("sqlite: nil database")
	}
	db.closeOnce.Do(func() {
		db.mu.Lock()
		db.closed = true
		transactions := make([]*Transaction, 0, len(db.transactions))
		for tx := range db.transactions {
			transactions = append(transactions, tx)
		}
		db.mu.Unlock()
		for _, tx := range transactions {
			db.closeErr = errors.Join(db.closeErr, Database_SQLite3_RollbackImpl(tx))
		}
		if db.anchor != nil {
			db.closeErr = errors.Join(db.closeErr, db.anchor.Close())
		}
		db.closeErr = errors.Join(db.closeErr, db.db.Close())
	})
	return db.closeErr
}

func sqlParameters(params []any) ([]any, error) {
	out := make([]any, len(params))
	for i, value := range params {
		switch v := value.(type) {
		case nil, bool, string, int64:
			out[i] = v
		case int:
			out[i] = int64(v)
		case float64:
			if math.IsNaN(v) || math.IsInf(v, 0) {
				return nil, fmt.Errorf("sqlite: parameter %d is not finite", i)
			}
			out[i] = v
		case []byte:
			out[i] = append([]byte{}, v...)
		default:
			return nil, fmt.Errorf("sqlite: unsupported parameter %d (%T)", i, value)
		}
	}
	return out, nil
}

type executor interface {
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
}

func executeSQL(ctx context.Context, target executor, statement string, params []any) (ExecResult, error) {
	args, err := sqlParameters(params)
	if err != nil {
		return ExecResult{}, err
	}
	result, err := target.ExecContext(ctx, statement, args...)
	if err != nil {
		return ExecResult{}, err
	}
	changes, err := result.RowsAffected()
	if err != nil {
		return ExecResult{}, err
	}
	if changes < 0 || changes > math.MaxInt32 {
		return ExecResult{}, errors.New("sqlite: affected row count exceeds PureScript Int")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return ExecResult{}, err
	}
	return ExecResult{Changes: changes, LastInsertRowId: strconv.FormatInt(id, 10)}, nil
}

func querySQL(ctx context.Context, target executor, statement string, params []any) (QueryResult, error) {
	args, err := sqlParameters(params)
	if err != nil {
		return QueryResult{}, err
	}
	rows, err := target.QueryContext(ctx, statement, args...)
	if err != nil {
		return QueryResult{}, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return QueryResult{}, err
	}
	result := QueryResult{Columns: columns, Rows: make([][]any, 0)}
	for rows.Next() {
		values := make([]any, len(columns))
		pointers := make([]any, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}
		if err := rows.Scan(pointers...); err != nil {
			return QueryResult{}, err
		}
		for i, value := range values {
			if bytes, ok := value.([]byte); ok {
				values[i] = append([]byte{}, bytes...)
			}
		}
		result.Rows = append(result.Rows, values)
	}
	if err := rows.Err(); err != nil {
		return QueryResult{}, err
	}
	if err := rows.Close(); err != nil {
		return QueryResult{}, err
	}
	return result, nil
}

func Database_SQLite3_ExecuteImpl(ctx context.Context, db *Database, statement string, params []any) (ExecResult, error) {
	if db == nil {
		return ExecResult{}, errors.New("sqlite: nil database")
	}
	return executeSQL(ctx, db.db, statement, params)
}

func Database_SQLite3_QueryImpl(ctx context.Context, db *Database, statement string, params []any) (QueryResult, error) {
	if db == nil {
		return QueryResult{}, errors.New("sqlite: nil database")
	}
	return querySQL(ctx, db.db, statement, params)
}

func Database_SQLite3_BeginImpl(ctx context.Context, db *Database) (*Transaction, func(), error) {
	if db == nil {
		return nil, nil, errors.New("sqlite: nil database")
	}
	conn, err := db.db.Conn(ctx)
	if err != nil {
		return nil, nil, err
	}
	// The bridge releases ctx after acquisition. The transaction owns a separate
	// lifetime so returning successfully does not immediately roll it back.
	lifetime, cancel := context.WithCancel(context.Background())
	tx, err := conn.BeginTx(lifetime, nil)
	if err != nil {
		cancel()
		_ = conn.Close()
		return nil, nil, err
	}
	result := &Transaction{owner: db, conn: conn, tx: tx, ctx: lifetime, cancel: cancel}
	cleanup := func() { _ = Database_SQLite3_RollbackImpl(result) }
	db.mu.Lock()
	if db.closed {
		db.mu.Unlock()
		cleanup()
		return nil, nil, errors.New("sqlite: database is closed")
	}
	db.transactions[result] = struct{}{}
	db.mu.Unlock()
	if err := ctx.Err(); err != nil {
		return nil, cleanup, err
	}
	return result, cleanup, nil
}

func finishTransaction(tx *Transaction) {
	tx.done = true
	tx.cancel()
	_ = tx.conn.Close()
	tx.owner.mu.Lock()
	delete(tx.owner.transactions, tx)
	tx.owner.mu.Unlock()
}

func Database_SQLite3_CommitImpl(ctx context.Context, tx *Transaction) error {
	if tx == nil {
		return errors.New("sqlite: nil transaction")
	}
	tx.mu.Lock()
	defer tx.mu.Unlock()
	if tx.done {
		return sql.ErrTxDone
	}
	if err := ctx.Err(); err != nil {
		_ = tx.tx.Rollback()
		finishTransaction(tx)
		return err
	}
	err := tx.tx.Commit()
	if err != nil {
		// SQLite keeps a transaction open when a deferred constraint rejects
		// COMMIT, although database/sql has already marked its Tx done. Roll back
		// on the reserved connection before putting it back in the pool.
		if _, rollbackErr := tx.conn.ExecContext(context.Background(), "ROLLBACK"); rollbackErr != nil {
			// Never reuse a connection with uncertain transaction state.
			_ = tx.conn.Raw(func(any) error { return driver.ErrBadConn })
		}
	}
	finishTransaction(tx)
	return err
}

func Database_SQLite3_RollbackImpl(tx *Transaction) error {
	if tx == nil {
		return errors.New("sqlite: nil transaction")
	}
	tx.cancel()
	tx.mu.Lock()
	defer tx.mu.Unlock()
	if tx.done {
		return nil
	}
	err := tx.tx.Rollback()
	finishTransaction(tx)
	if errors.Is(err, sql.ErrTxDone) {
		return nil
	}
	return err
}

func Database_SQLite3_ExecuteTransactionImpl(ctx context.Context, tx *Transaction, statement string, params []any) (ExecResult, error) {
	if tx == nil {
		return ExecResult{}, errors.New("sqlite: nil transaction")
	}
	tx.mu.Lock()
	defer tx.mu.Unlock()
	if tx.done {
		return ExecResult{}, sql.ErrTxDone
	}
	operation, cancel := context.WithCancel(ctx)
	stop := context.AfterFunc(tx.ctx, cancel)
	defer func() { stop(); cancel() }()
	return executeSQL(operation, tx.tx, statement, params)
}

func Database_SQLite3_QueryTransactionImpl(ctx context.Context, tx *Transaction, statement string, params []any) (QueryResult, error) {
	if tx == nil {
		return QueryResult{}, errors.New("sqlite: nil transaction")
	}
	tx.mu.Lock()
	defer tx.mu.Unlock()
	if tx.done {
		return QueryResult{}, sql.ErrTxDone
	}
	operation, cancel := context.WithCancel(ctx)
	stop := context.AfterFunc(tx.ctx, cancel)
	defer func() { stop(); cancel() }()
	return querySQL(operation, tx.tx, statement, params)
}

var Database_SQLite3_Null any = nil

func Database_SQLite3_IntegerTextImpl[R any](just func(string) R, nothing R, value any) R {
	if integer, ok := value.(int64); ok {
		return just(strconv.FormatInt(integer, 10))
	}
	return nothing
}

func Database_SQLite3_Integer64Impl[R any](just func(any) R, nothing R, text string) R {
	value, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return nothing
	}
	return just(value)
}

// --- Auto-generated FFI wrappers ---
var _Gopurs_Database_SQLite3_BeginImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAffAcquire(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, func(), error) {
		var go_arg0 *Database
		if arg0.Type == gopurs_runtime.TypeAny {
			go_arg0 = arg0.AnyVal().(*Database)
		} else {
			go_arg0 = *(**Database)(arg0.UnsafePtr)
		}
		go_res, go_cleanup, go_err := Database_SQLite3_BeginImpl(ctx, go_arg0)
		if go_err != nil {
			return gopurs_runtime.Value{}, go_cleanup, go_err
		}
		return gopurs_runtime.Box(go_res), go_cleanup, nil
	})
})
var _Gopurs_Database_SQLite3_CloseImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAff(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, error) {
		var go_arg0 *Database
		if arg0.Type == gopurs_runtime.TypeAny {
			go_arg0 = arg0.AnyVal().(*Database)
		} else {
			go_arg0 = *(**Database)(arg0.UnsafePtr)
		}
		go_err := Database_SQLite3_CloseImpl(go_arg0)
		return gopurs_runtime.Value{}, go_err
	})
})
var _Gopurs_Database_SQLite3_CommitImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAff(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, error) {
		var go_arg0 *Transaction
		if arg0.Type == gopurs_runtime.TypeAny {
			go_arg0 = arg0.AnyVal().(*Transaction)
		} else {
			go_arg0 = *(**Transaction)(arg0.UnsafePtr)
		}
		go_err := Database_SQLite3_CommitImpl(ctx, go_arg0)
		return gopurs_runtime.Value{}, go_err
	})
})
var _Gopurs_Database_SQLite3_ExecuteImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAff(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, error) {
		var go_arg0 *Database
		if arg0.Type == gopurs_runtime.TypeAny {
			go_arg0 = arg0.AnyVal().(*Database)
		} else {
			go_arg0 = *(**Database)(arg0.UnsafePtr)
		}
		go_arg1 := gopurs_runtime.Unbox[string](arg1)
		go_arg2 := gopurs_runtime.ForeignValue(arg2).([]any)
		go_res, go_err := Database_SQLite3_ExecuteImpl(ctx, go_arg0, go_arg1, go_arg2)
		if go_err != nil {
			return gopurs_runtime.Value{}, go_err
		}
		return func() gopurs_runtime.Value {
			_source := go_res
			_ = _source
			return gopurs_runtime.RecordDict2("changes", "lastInsertRowId", gopurs_runtime.Int(_source.Changes), gopurs_runtime.Box(_source.LastInsertRowId))
		}(), nil
	})
})
var _Gopurs_Database_SQLite3_ExecuteTransactionImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAff(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, error) {
		var go_arg0 *Transaction
		if arg0.Type == gopurs_runtime.TypeAny {
			go_arg0 = arg0.AnyVal().(*Transaction)
		} else {
			go_arg0 = *(**Transaction)(arg0.UnsafePtr)
		}
		go_arg1 := gopurs_runtime.Unbox[string](arg1)
		go_arg2 := gopurs_runtime.ForeignValue(arg2).([]any)
		go_res, go_err := Database_SQLite3_ExecuteTransactionImpl(ctx, go_arg0, go_arg1, go_arg2)
		if go_err != nil {
			return gopurs_runtime.Value{}, go_err
		}
		return func() gopurs_runtime.Value {
			_source := go_res
			_ = _source
			return gopurs_runtime.RecordDict2("changes", "lastInsertRowId", gopurs_runtime.Int(_source.Changes), gopurs_runtime.Box(_source.LastInsertRowId))
		}(), nil
	})
})
var _Gopurs_Database_SQLite3_Integer64Impl = // TAST: (ForAll [r] (Func [(Func [(ADT ["Foreign","Foreign"] [])] (TypeVar r)), (TypeVar r), String] (TypeVar r)))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) gopurs_runtime.Value {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Any(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Database_SQLite3_Integer64Impl[gopurs_runtime.Value](go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Database_SQLite3_IntegerTextImpl = // TAST: (ForAll [r] (Func [(Func [String] (TypeVar r)), (TypeVar r), (ADT ["Foreign","Foreign"] [])] (TypeVar r)))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) gopurs_runtime.Value {
		return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
	}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.ForeignValue(arg2)
	go_res := Database_SQLite3_IntegerTextImpl[gopurs_runtime.Value](go_arg0, go_arg1, go_arg2)
	return go_res
})
var _Gopurs_Database_SQLite3_Null = // TAST: (ADT ["Foreign","Foreign"] [])
gopurs_runtime.Box(Database_SQLite3_Null)
var _Gopurs_Database_SQLite3_OpenImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAffAcquire(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, func(), error) {
		go_arg0 := gopurs_runtime.Unbox[string](arg0)
		go_res, go_cleanup, go_err := Database_SQLite3_OpenImpl(ctx, go_arg0)
		if go_err != nil {
			return gopurs_runtime.Value{}, go_cleanup, go_err
		}
		return gopurs_runtime.Box(go_res), go_cleanup, nil
	})
})
var _Gopurs_Database_SQLite3_QueryImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAff(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, error) {
		var go_arg0 *Database
		if arg0.Type == gopurs_runtime.TypeAny {
			go_arg0 = arg0.AnyVal().(*Database)
		} else {
			go_arg0 = *(**Database)(arg0.UnsafePtr)
		}
		go_arg1 := gopurs_runtime.Unbox[string](arg1)
		go_arg2 := gopurs_runtime.ForeignValue(arg2).([]any)
		go_res, go_err := Database_SQLite3_QueryImpl(ctx, go_arg0, go_arg1, go_arg2)
		if go_err != nil {
			return gopurs_runtime.Value{}, go_err
		}
		return func() gopurs_runtime.Value {
			_source := go_res
			_ = _source
			return gopurs_runtime.RecordDict2("columns", "rows", func() gopurs_runtime.Value {
				_source := _source.Columns
				_out := make([]gopurs_runtime.Value, len(_source))
				for _i, _item := range _source {
					_out[_i] = gopurs_runtime.Box(_item)
				}
				return gopurs_runtime.Array(_out)
			}(), func() gopurs_runtime.Value {
				_source := _source.Rows
				_out := make([]gopurs_runtime.Value, len(_source))
				for _i, _item := range _source {
					_out[_i] = func() gopurs_runtime.Value {
						_source := _item
						_out := make([]gopurs_runtime.Value, len(_source))
						for _i, _item := range _source {
							_out[_i] = gopurs_runtime.Any(_item)
						}
						return gopurs_runtime.Array(_out)
					}()
				}
				return gopurs_runtime.Array(_out)
			}())
		}(), nil
	})
})
var _Gopurs_Database_SQLite3_QueryTransactionImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAff(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, error) {
		var go_arg0 *Transaction
		if arg0.Type == gopurs_runtime.TypeAny {
			go_arg0 = arg0.AnyVal().(*Transaction)
		} else {
			go_arg0 = *(**Transaction)(arg0.UnsafePtr)
		}
		go_arg1 := gopurs_runtime.Unbox[string](arg1)
		go_arg2 := gopurs_runtime.ForeignValue(arg2).([]any)
		go_res, go_err := Database_SQLite3_QueryTransactionImpl(ctx, go_arg0, go_arg1, go_arg2)
		if go_err != nil {
			return gopurs_runtime.Value{}, go_err
		}
		return func() gopurs_runtime.Value {
			_source := go_res
			_ = _source
			return gopurs_runtime.RecordDict2("columns", "rows", func() gopurs_runtime.Value {
				_source := _source.Columns
				_out := make([]gopurs_runtime.Value, len(_source))
				for _i, _item := range _source {
					_out[_i] = gopurs_runtime.Box(_item)
				}
				return gopurs_runtime.Array(_out)
			}(), func() gopurs_runtime.Value {
				_source := _source.Rows
				_out := make([]gopurs_runtime.Value, len(_source))
				for _i, _item := range _source {
					_out[_i] = func() gopurs_runtime.Value {
						_source := _item
						_out := make([]gopurs_runtime.Value, len(_source))
						for _i, _item := range _source {
							_out[_i] = gopurs_runtime.Any(_item)
						}
						return gopurs_runtime.Array(_out)
					}()
				}
				return gopurs_runtime.Array(_out)
			}())
		}(), nil
	})
})
var _Gopurs_Database_SQLite3_RollbackImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.EffectFnAff(func(ctx gopurs_runtime.FFIContext) (gopurs_runtime.Value, error) {
		var go_arg0 *Transaction
		if arg0.Type == gopurs_runtime.TypeAny {
			go_arg0 = arg0.AnyVal().(*Transaction)
		} else {
			go_arg0 = *(**Transaction)(arg0.UnsafePtr)
		}
		go_err := Database_SQLite3_RollbackImpl(go_arg0)
		return gopurs_runtime.Value{}, go_err
	})
})
