package store

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Model[T Schemable] struct {
	ID int64
	T  T
}

// NewStore ...
func NewStore[T Schemable](tableName string, ctor func() T) (*Store[T], error) {
	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		return nil, err
	}
	t := ctor()
	schema := t.Schema()
	stmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT ", tableName)
	for i := range schema {
		stmt += fmt.Sprintf(",`%s` %s", schema[i].ColumnName, schema[i].ColumnType)
	}
	stmt += ")"
	if _, err = db.Exec(stmt); err != nil {
		return nil, err
	}
	return &Store[T]{TableName: tableName, Ctor: ctor, DB: db}, nil
}

// Schemable ...
type Schemable interface {
	Schema() []SchemaEntry
}

// Store ...
type Store[T Schemable] struct {
	tx        *sql.Tx
	DB        *sql.DB
	Ctor      func() T
	TableName string
}

// SchemaEntry ...
type SchemaEntry struct {
	ColumnName  string
	ColumnType  string
	ColumnValue any
}

//
//func (s Store[T]) NewTX(ctx context.Context) Store[T] {
//
//	tx, err:=s.DB.BeginTx(ctx, nil)
//	tx.
//	return
//}

// Insert ...
func (s Store[T]) Insert(ctx context.Context, t T) (*Model[T], error) {
	m := Model[T]{T: t}
	str := fmt.Sprintf("INSERT INTO %s (%s) values(%s)", s.TableName, strings.Join(m.ColumnNames(), ","), strings.Repeat("?,", len(m.ColumnNames())-1)+"?")
	stmt, err := s.DB.Prepare(str)
	if err != nil {
		return nil, err
	}
	res, err := stmt.ExecContext(ctx, m.ColumnValues()...)
	if err != nil {
		return nil, err
	}
	m.ID, err = res.LastInsertId()
	return &m, err
}

// Get ...
func (s Store[T]) Get(ctx context.Context, id int64) (*Model[T], error) {
	m := Model[T]{T: s.Ctor()}
	query := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", strings.Join(m.ColumnNames(), ","), s.TableName)
	row := s.DB.QueryRowContext(ctx, query, id)
	if err := row.Scan(m.ColumnValues()...); err != nil {
		return nil, err
	}
	m.ID = id
	return &m, nil
}

// Update ...
func (s Store[T]) Update(ctx context.Context, m Model[T]) error {
	fields := ""
	for i, col := range m.ColumnNames() {
		if i != 0 {
			fields += ","
		}
		fields += fmt.Sprintf(" %s = ? ", col)
	}
	stmt, err := s.DB.PrepareContext(ctx, fmt.Sprintf("UPDATE %s SET %s WHERE ID = ?", s.TableName, fields))
	if err != nil {
		return err
	}

	vals := append(m.ColumnValues(), m.ID)
	_, err = stmt.ExecContext(ctx, vals...)
	return err
}

// All ...
func (s Store[T]) All(ctx context.Context) ([]Model[T], error) {
	var entries []Model[T]
	m := Model[T]{T: s.Ctor()}
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(m.ColumnNames(), ","), s.TableName)
	rows, err := s.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		t := s.Ctor()
		m := Model[T]{T: t}
		if err := rows.Scan(m.ColumnValues()...); err != nil {
			return nil, err
		}
		entries = append(entries, m)
	}
	return entries, nil
}

// ColumnNames ...
func (m *Model[T]) ColumnNames() []string {
	schema := m.T.Schema()
	f := make([]string, 0, len(schema))
	for _, schema := range schema {
		f = append(f, schema.ColumnName)
	}
	return f
}

// ColumnValues ...
func (m *Model[T]) ColumnValues() []any {
	schema := m.T.Schema()
	v := make([]any, 0, len(schema))
	for _, schema := range schema {
		v = append(v, schema.ColumnValue)
	}
	return v
}
