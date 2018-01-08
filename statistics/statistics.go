// Thread safe statistics

package statistics

import (
  "sync"
)

var (
  // Predefined Statistics
  stats   map[string]*Statistic = make( map[string]*Statistic )
  // Mutex
  mutex  *sync.Mutex = &sync.Mutex{}
)

type Statistics struct {
  // If set then log stats to the log every duration
  Log         bool
  // If set then present /stats endpoint with json output
  Statistics  bool
  // The schedule to use to collect statistics, defaults to every minute
  Schedule    string
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

func Incr( n string ) {
  IncrVal( n, 1 )
}

func Decr( n string ) {
  IncrVal( n, -1 )
}

func IncrVal( n string, v int64 ) {
  mutex.Lock()
  getOrCreate( n ).incr( v )
  mutex.Unlock()
}

func Set( n string, v int64 ) {
  mutex.Lock()
  getOrCreate( n ).set( v )
  mutex.Unlock()
}

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
