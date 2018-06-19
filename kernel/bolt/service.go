// A simple kernel service wich provides access to a single github.com/coreos/bbolt
// object store
package bolt

import (
  bbolt "github.com/coreos/bbolt"
  "fmt"
  "time"
)

type BoltService struct {
  FileName  string
  db       *bbolt.DB
}

func (s *BoltService) Name() string {
  return "bolt:" + s.FileName
}

func (s *BoltService) PostInit() error {
  if s.FileName == "" {
    return fmt.Errorf( "No filename supplied for bbolt" )
  }
  return nil
}

func (s *BoltService) Start() error {
  db, err := bbolt.Open( s.FileName, 0666, &bbolt.Options{
    Timeout: 5 * time.Second,
    } )
    if err != nil {
      return err
    }
    s.db = db
  return nil
}

func (s *BoltService) Stop() {
  s.db.Close()
}
