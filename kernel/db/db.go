package db

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
  "flag"
  "github.com/peter-mount/golib/kernel"
  "os"
)

// database/sql bound with github.com/lib/pq as a Kernel Service
type DBService struct {
  postgresURI      *string
  db               *sql.DB
}

func (s *DBService) Name() string {
  return "kernel.DBService"
}

func (s *DBService) Init( k *kernel.Kernel ) error {
  s.postgresURI = flag.String( "d", "", "The database to connect to" )
  return nil
}

func (s *DBService) PostInit() error {
  if *s.postgresURI == "" {
    *s.postgresURI = os.Getenv( "POSTGRESDB" )
  }
  if *s.postgresURI == "" {
    return fmt.Errorf( "No database uri provided" )
  }
  return nil
}

func (s *DBService) Start() error {
  db, err := sql.Open( "postgres", *s.postgresURI )
  if err != nil {
    return err
  }
  s.db = db
  return nil
}

func (s *DBService) Stop() {
  if s.db != nil {

    s.db.Close()
    s.db = nil
  }
}

func (s *DBService) GetDB() *sql.DB {
  return s.db
}
