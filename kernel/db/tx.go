package db

import (
  "database/sql"
  "log"
)

// Tx wrapper around sql.Tx with additional functionality
type Tx struct {
  // wrapped tx
  tx    *sql.Tx
  // parent DBservice
  db    *DBService
  // ionCommit code
  onCommit []onCommit
}

type onCommit func() error

type TxHandler func( *Tx) error

// Update calls a TxHandler within a transaction.
// If the handler returns an error then the transaction is rolled back.
// If the handler returns nil then the transaction is committed
func (s *DBService) Update( f TxHandler ) error {
  tx := &Tx{ db: s }

  ptx, err := s.db.Begin()
  if err != nil {
    return err
  }
  tx.tx = ptx
  defer tx.commit()

  err = f( tx )

  if err != nil {
    tx.rollback()
    return err
  }

  return nil
}

func (tx *Tx) commit() error {
  err := tx.tx.Commit()
  if err != nil {
    return err
  }

  for _, f := range tx.onCommit {
    err := f()
    if err != nil {
      return err
    }
  }

  return nil
}

func (tx *Tx) rollback() error {
  return tx.tx.Rollback()
}

// Execute a query within this transaction
func (tx *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
  r, e := tx.tx.Exec( query, args... )
  if tx.db.Debug && e != nil {
    log.Printf( "Error: %s\n%v", query, e )
  }
  return r, e
}

// Prepare a statement that will run within this transaction
func (tx *Tx) Prepare(query string) (*sql.Stmt, error) {
  s, e := tx.tx.Prepare( query )
  return s, e
}

// Perform a Query within this transaction
func (tx *Tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
  r, e := tx.tx.Query( query, args... )
  if tx.db.Debug && e != nil {
    log.Printf( "Error: %s\n%v", query, e )
  }
  return r, e
}

func (tx *Tx) QueryRow(query string, args ...interface{}) *sql.Row {
  return tx.tx.QueryRow(query, args...)
}

// OnCommit will call the supplied function once this transaction commits
func (tx *Tx) OnCommit( f func() error ) {
  tx.onCommit = append( tx.onCommit, f )
}

// ExecOnCommit will execute a query once this transaction commits.
// Note this will not be within a transaction, see UpdateOnCommit for that
func (tx *Tx) ExecOnCommit( query string, args ...interface{} ) {
  tx.OnCommit( func() error {
    _, err := tx.db.db.Exec( query, args... )
    if tx.db.Debug && err != nil {
      log.Printf( "Error: %s\n%v", query, err )
    }
    return err
  })
}

// UpdateOnCommit will start a new update/transaction once this transaction commits
func (tx *Tx) UpdateOnCommit( f TxHandler ) {
  tx.OnCommit( func() error {
    return tx.db.Update( f )
  })
}
