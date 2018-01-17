// A basic REST server supporting HTTP.
//
// This package implements a HTTP server using net/http and github.com/gorilla/mux
// taking away most of the required boiler plate code usually needed when implementing
// basic REST services. It also provides many utility methods for handling both JSON and XML responses.
package rest

import (
  "fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

// The internal config of a Server
type Server struct {
  // Port to listen to
  port    int
  // The mux Router
  router  *mux.Router
}

// Create a new Server with the specified port
func NewServer( port int ) *Server {
  s := &Server{}

  // If not defined then use port 80
  if port < 1 || port > 65534 {
    s.port = 8080
  } else {
    s.port = port
  }

  s.router = mux.NewRouter()

  return s
}

// Start starts the server
func (s *Server) Start() error {
  log.Printf( "Listening on port %d\n", s.port )
  return http.ListenAndServe( fmt.Sprintf( ":%d", s.port ), s.router )
}
