package statistics

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
)

// StatsHandler installs a github.com/gorilla/mux handler under the path /stats
// which exposes the current statistics & any history via a simple HTTP GET request.
func StatsHandler( router *mux.Router ) {
  router.HandleFunc( "/stats", getStats ).Methods( "GET" )
}

// Handler for /stats
func getStats(w http.ResponseWriter, r *http.Request) {
  var result = make( map[string]*Statistic )

  mutex.Lock()

  for key,value := range stats {
    if value.latest != nil {
      result[key] = value.latest.clone()
      result[key].History = statsCopyArray( value.History )
    }
  }

  mutex.Unlock()

  json.NewEncoder(w).Encode( result )
}
