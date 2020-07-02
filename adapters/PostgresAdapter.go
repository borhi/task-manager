package adapters

import (
	"database/sql"
)

type PostgresAdapter struct {
	Connection *sql.DB
}

type IDbAdapter interface {
	Execute(statement string) (sql.Result, error)
	Query(statement string) (IRow, error)
}

type IRow interface {
	Scan(dest ...interface{}) error
	Next() bool
}

func (adapter PostgresAdapter) Execute(statement string) (sql.Result, error) {
	return adapter.Connection.Exec(statement)
}

func (adapter PostgresAdapter) Query(statement string) (IRow, error) {
	rows, err := adapter.Connection.Query(statement)
	if err != nil {
		return new(PostgresRow), err
	}

	row := new(PostgresRow)
	row.Rows = rows

	return row, nil
}

type PostgresRow struct {
	Rows *sql.Rows
}

func (r PostgresRow) Scan(dest ...interface{}) error {
	err := r.Rows.Scan(dest...)
	if err != nil {
		return err
	}

	return nil
}

func (r PostgresRow) Next() bool {
	return r.Rows.Next()
}
