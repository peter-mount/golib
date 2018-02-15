# statistics
--
    import "github.com/peter-mount/golib/statistics"

Thread safe statistics

Statistics are a collection of named values which are sampled over a
predetermined sample period (default 1 minute).

An application can then update the statistic whenever it needs to and doesn't
need to worry about either exposing or logging those values as the library
handles that transparently.

## Setting up

At the start of an application, it needs to initialise the library:

    statConfig := statistics.Statistics{ Log: true }
    statConfig.Configure()

This will configure the library to sample statistics once a minute logging the
sampled values (if not 0) to the console.

Once that has been done, then the application can simply call one of the
functions that update a named statistic, usually Incr():

    statistics.Incr( "some.stat.name" )

Anther example is sampling some state. In this case you can use Set() to set the
value of the statistic:

    statistics.Set( "td.all", 10 )

An examle of this one can be seen in the nrod-td repository where we record the
latency (i.e. how late a message is received) of messages from the Network Rail
Train Describer feed. In that instance the statistic's value is the most recent
latency but min/max are the min/max value during the sample period, ave the
average and count the number of messages received.

Exposing the statistics over a Rest interface is also supported via the
StatsHandler method.

## Usage

```go
const (
	// 72 entries * STATS_HISTORY_PERIOD m = 6 hours when default schedule of 1 minute
	STATS_MAX_HISTORY = 72
	// Period of history in schedule units
	STATS_HISTORY_PERIOD = 5
)
```

#### func  Decr

```go
func Decr(n string)
```
Decr decrements the named statistic value by 1 The count field is also
incremented by 1.

#### func  Incr

```go
func Incr(n string)
```
Incr increments the named statistic value by 1. The count field is also
incremented by 1.

#### func  IncrVal

```go
func IncrVal(n string, v int64)
```
IncrVal increments the named statistic value by an arbitary amount. The count
field is also incremented by 1.

#### func  Set

```go
func Set(n string, v int64)
```
Set sets the value of a named statistic. The count field is also incremented by
1.

#### func  StatsHandler

```go
func StatsHandler(router *mux.Router)
```
StatsHandler installs a github.com/gorilla/mux handler under the path /stats
which exposes the current statistics & any history via a simple HTTP GET
request.

#### func  StatsRestHandler

```go
func StatsRestHandler(w http.ResponseWriter, r *http.Request)
```
Handler for /stats

#### type Statistic

```go
type Statistic struct {

	// The timestamp of the last operation
	Timestamp int64 `json:"timestamp"`
	// the current value
	Value int64 `json:"value"`
	// the number of updates to this statistic during the sample period
	Count int64 `json:"count"`
	// The minimum value during the sample period
	Min int64 `json:"min"`
	// The maximum value during the sample period
	Max int64 `json:"max"`
	// The average value during the sample period
	Ave int64 `json:"average"`
	// The sum of all values during the sample period
	Sum int64 `json:"sum"`
	// Historic data, max 72 entries at 5 minute intervals
	History []*Statistic `json:"history,omitempty"`
}
```

A basic statistic

#### func  Get

```go
func Get(n string) *Statistic
```
Get returns a Statistic instance based on the current value. Note this will be a
snapshot only of the current value. Changing anything in the returned Statistic
will not affect what the library is managing.

#### type Statistics

```go
type Statistics struct {
	// If set then log stats to the log every duration
	Log bool
	// The schedule to use to collect statistics, defaults to every minute
	Schedule string
	// The Cron instance in use. Either provide this before calling Configure()
	// or one will be created for you.
	Cron *cron.Cron
}
```

Global statistics configuration

#### func (*Statistics) Configure

```go
func (s *Statistics) Configure()
```
Configure takes a Statistics instance and configures the system to use that
configuration. Note: You should only call this once during an application's
lifetime.
