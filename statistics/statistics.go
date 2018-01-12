// Thread safe statistics
//
// Statistics are a collection of named values which are sampled over a
// predetermined sample period (default 1 minute).
//
// An application can then update the statistic whenever it needs to and doesn't
// need to worry about either exposing or logging those values as the library
// handles that transparently.
//
// ## Setting up
//
// At the start of an application, it needs to initialise the library:
//
//    statConfig := statistics.Statistics{ Log: true }
//    statConfig.Configure()
//
// This will configure the library to sample statistics once a minute logging
// the sampled values (if not 0) to the console.
//
// Once that has been done, then the application can simply call one of the
// functions that update a named statistic, usually Incr():
//
//    statistics.Incr( "some.stat.name" )
//
// Anther example is sampling some state. In this case you can use Set() to set
// the value of the statistic:
//
//    statistics.Set( "td.all", 10 )
//
// An examle of this one can be seen in the nrod-td repository where we record
// the latency (i.e. how late a message is received) of messages from the
// Network Rail Train Describer feed. In that instance the statistic's value is
// the most recent latency but min/max are the min/max value during the sample
// period, ave the average and count the number of messages received.
//
// Exposing the statistics over a Rest interface is also supported via the
// StatsHandler method.
package statistics

import (
  "gopkg.in/robfig/cron.v2"
  "sync"
)

var (
  // Predefined Statistics
  stats       map[string]*Statistic = make( map[string]*Statistic )
  // Mutex
  mutex      *sync.Mutex = &sync.Mutex{}
  logStats    bool
)

// Global statistics configuration
type Statistics struct {
  // If set then log stats to the log every duration
  Log         bool
  // The schedule to use to collect statistics, defaults to every minute
  Schedule    string
  // The Cron instance in use. Either provide this before calling Configure()
  // or one will be created for you.
  Cron        *cron.Cron
}

// Configure takes a Statistics instance and configures the system to use that
// configuration.
// Note: You should only call this once during an application's lifetime.
func (s *Statistics) Configure() {
  logStats = s.Log

  // If no Cron defined create one
  if s.Cron == nil {
    s.Cron = cron.New()
    s.Cron.Start()
  }

  // The schedule, default to every minute if not set
  var schedule = s.Schedule
  if schedule == "" {
    schedule = "0 * * * * *"
  }

  // Schedule the recorder service
  s.Cron.AddFunc( schedule, statsRecord )
}

// return a statistic creating it as needed
func getOrCreate( n string ) *Statistic {
  if val, ok := stats[n]; ok {
    return val
  }
  stats[n] = new(Statistic)
  stats[n].name = n
  stats[n].reset()
  return stats[n]
}

// Incr increments the named statistic value by 1.
// The count field is also incremented by 1.
func Incr( n string ) {
  IncrVal( n, 1 )
}

// Decr decrements the named statistic value by 1
// The count field is also incremented by 1.
func Decr( n string ) {
  IncrVal( n, -1 )
}

// IncrVal increments the named statistic value by an arbitary amount.
// The count field is also incremented by 1.
func IncrVal( n string, v int64 ) {
  mutex.Lock()
  getOrCreate( n ).incr( v )
  mutex.Unlock()
}

// Set sets the value of a named statistic.
// The count field is also incremented by 1.
func Set( n string, v int64 ) {
  mutex.Lock()
  getOrCreate( n ).set( v )
  mutex.Unlock()
}

// Get returns a Statistic instance based on the current value.
// Note this will be a snapshot only of the current value. Changing anything
// in the returned Statistic will not affect what the library is managing.
func Get( n string ) *Statistic {
  var stat *Statistic

  mutex.Lock()

  val, ok := stats[n]

  if( ok ) {
    if val.latest != nil {
      stat = val.latest.clone()
    } else {
      stat =  val.clone()
    }
  }

  mutex.Unlock()

  return stat
}
