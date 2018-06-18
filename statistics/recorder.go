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

// Recorder allows an external service to be plugged in to receive statistics
// once they have been collected
type Recorder interface {
  // PublishStatistic accepts a statistic for external processing, usually
  // publishing to rabbitmq or statsd etc.
  // The passed Statistic is a copy so modifying it will not affect the stats.
  PublishStatistic( string, *Statistic )
}

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
func (s *Statistics) statsRecord() {
  mutex.Lock()
  defer mutex.Unlock()

  var publish []*Statistic

  for _, value := range stats {
    value.logState()
    value.recordHistory()
    value.reset()

    // Append to publish list if we have a latest value
    if s.Recorder != nil && value.latest != nil {
      publish = append( publish, value.latest.clone() )
    }
  }

  // Publish any to the optional external recorder
  if s.Recorder != nil && len( publish ) > 0 {
    go func() {
      for _, value := range publish {
        s.Recorder.PublishStatistic( value.name, value )
      }
    }()
  }
}
