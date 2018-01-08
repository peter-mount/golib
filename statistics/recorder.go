// Thread safe statistics

package statistics

import (
  "log"
)

const (
  // 72 entries * STATS_HISTORY_PERIOD m = 6 hours when default schedule of 1 minute
  STATS_MAX_HISTORY = 72
  // Period of history in schedule units
  STATS_HISTORY_PERIOD = 5
)

func (s *Statistic) recordHistory() {
  // Add to last 5 entries
  s.latest = s.clone()
  s.lastFive = append( s.lastFive, s.latest )
  // If full then collate and push to history
  if len( s.lastFive ) >= STATS_HISTORY_PERIOD {
    // Form new statistoc of sum of all entries within it
    var hist = new(Statistic)
    for _, val := range s.lastFive {
      hist.Value += val.Value
      hist.Count += val.Count
      hist.Sum += val.Sum
    }
    hist.update()
    s.History = append( s.History, hist )
    s.lastFive = nil

    // Keep history down to size
    if len( s.History ) > STATS_MAX_HISTORY {
      s.History = s.History[1:]
    }
  }
}

func (value *Statistic) logState() {
  // Don't report stats with no submitted values, i.e. Min > Max
  if logStats &&  value.Min <= value.Max {
    log.Printf(
      "%s Val %d Count %d Min %d Max %d Sum %d Ave %d\n",
      value.name,
      value.Value,
      value.Count,
      value.Min,
      value.Max,
      value.Sum,
      value.Ave )
  }
}

// Record then reset all Statistics
func statsRecord() {
  mutex.Lock()

  for _, value := range stats {
    value.logState()
    value.recordHistory()
    value.reset()
  }

  mutex.Unlock()
}
