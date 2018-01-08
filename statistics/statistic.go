// Thread safe statistics

package statistics

import (
  "time"
)

type Statistic struct {
  // The statistic name
  name      string        `json:"-"`
  // The timestamp of the last operation
  Timestamp int64         `json:"timestamp"`
  // the current value
  Value     int64         `json:"value"`
  // the number of updates
  Count     int64         `json:"count"`
  // The minimum value
  Min       int64         `json:"min"`
  // The maximum value
  Max       int64         `json:"max"`
  // The average value
  Ave       int64         `json:"average"`
  // The sum of all values
  Sum       int64         `json:"sum"`
  // Historic data, max 72 entries at 5 minute intervals
  History   []*Statistic  `json:"history,omitempty"`
  // The last 5 minutes data, used to build the history
  lastFive  []*Statistic  `json:"-"`
  latest    *Statistic    `json:"-"`
}

func (s *Statistic) reset() {
  s.Timestamp = time.Now().Unix()
  s.Value = 0
  s.Count = 0
  s.Min = int64(^uint64(0) >> 1)
  s.Max = -s.Min - 1
  s.Ave = 0
  s.Sum = 0
}

func (s *Statistic) clone() *Statistic {
  var r *Statistic = new(Statistic)
  r.name = s.name
  r.Timestamp = s.Timestamp
  r.Value = s.Value
  r.Count = s.Count
  r.Min = s.Min
  r.Max = s.Max
  r.Ave = s.Ave
  r.Sum = s.Sum
  return r
}

func statsCopyArray( s []*Statistic ) []*Statistic {
  var a []*Statistic
  if len( s ) > 0 {
    for _, v := range s {
      a = append( a, v.clone() )
    }
  }
  return a
}

func (s *Statistic) update() {
  s.Timestamp = time.Now().Unix()
  if( s.Value < s.Min ) {
    s.Min = s.Value
  }
  if( s.Value > s.Max ) {
    s.Max = s.Value
  }

  // protect against /0 - incase count is reset for some reason
  if( s.Count != 0 && s.Sum != 0 ) {
    s.Ave = s.Sum / s.Count
    } else {
      s.Ave = 0
    }
}

func (s *Statistic) set( v int64 ) {
  s.Value = v
  s.Sum += v
  s.Count ++
  s.update()
}

func (s *Statistic) incr( v int64 ) {
  s.Value += v
  s.Sum += v
  s.Count ++
  s.update()
}
