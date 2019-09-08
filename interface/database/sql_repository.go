package database

type ConnectedSql interface {
	Exec(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Rows, error)
	QueryRow(string, ...interface{}) Row
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Rows interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

type Row interface {
	Scan(...interface{}) error
}
