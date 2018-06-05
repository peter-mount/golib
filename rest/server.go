// A basic REST server supporting HTTP.
//
// This package implements a HTTP server using net/http and github.com/gorilla/mux
// taking away most of the required boiler plate code usually needed when implementing
// basic REST services. It also provides many utility methods for handling both JSON and XML responses.
package rest

import (
  "fmt"
  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
  "log"
  "net/http"
)

// The internal config of a Server
type Server struct {
  // The permitted headers
  Headers     []string
  // The permitted Origins
  Origins     []string
  // The permitted methods
  Methods     []string
  // Port to listen to
  Port          int
  // The mux Router
  router       *mux.Router
  // Base Context
  ctx          *ServerContext
}

func (a *Server) Name() string {
  return "Rest Server"
}

func (s *Server) PostInit() error {
  s.router = mux.NewRouter()
  s.ctx = &ServerContext{ context: "", server: s }
  return nil
}

func (s *Server) Run() error {
  // If not defined then use port 80
  port := s.Port
  if port < 1 || port > 65534 {
    port = 8080
  }

  // The permitted headers
  if len( s.Headers ) == 0 {
    s.Headers = []string{"X-Requested-With", "Content-Type"}
  }
  if len( s.Origins ) == 0 {
    s.Origins = []string{"*"}
  }
  if len( s.Methods ) == 0 {
    s.Origins = []string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}
  }
  headersOk := handlers.AllowedHeaders( s.Headers )
  originsOk := handlers.AllowedOrigins( s.Origins )
  methodsOk := handlers.AllowedMethods( s.Methods )
  handler := handlers.CORS( originsOk, headersOk, methodsOk )( s.router )

  log.Printf( "Listening on port %d\n", port )
  return http.ListenAndServe( fmt.Sprintf( ":%d", port ), handler )
}
