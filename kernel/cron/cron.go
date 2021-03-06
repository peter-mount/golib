// This provides the Kernel a managed Cron service.
//
// This is simply a wrapper around gopkg.in/robfig/cron.v2
//
package cron

import (
  crn "gopkg.in/robfig/cron.v2"
  "github.com/peter-mount/golib/kernel"
)

// gopkg.in/robfig/crn.v2 as a Kernel Service
type CronService struct {
  cron   *crn.Cron
}

func (s *CronService) Name() string {
  return "kernel.CronService"
}

func (s *CronService) Init( k *kernel.Kernel ) error {
  s.cron = crn.New()
  return nil
}

func (s *CronService) Start() error {
  s.cron.Start()
  return nil
}

func (s *CronService) Stop() {
  s.cron.Stop()
}

func (s *CronService) AddFunc(spec string, cmd func()) (crn.EntryID, error) {
  id, err := s.cron.AddFunc( spec, cmd )
  return id, err
}

func (s *CronService) AddJob(spec string, cmd crn.Job) (crn.EntryID, error) {
  id, err := s.cron.AddJob( spec, cmd )
  return id, err
}

func (s *CronService) Entries() []crn.Entry {
  return s.cron.Entries()
}

func (s *CronService) Entry(id crn.EntryID) crn.Entry {
  return s.cron.Entry( id )
}

func (s *CronService) Remove(id crn.EntryID) {
  s.cron.Remove( id )
}

func (s *CronService) Schedule(schedule crn.Schedule, cmd crn.Job) crn.EntryID {
  return s.cron.Schedule( schedule, cmd )
}
