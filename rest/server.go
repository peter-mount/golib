// A basic REST server supporting HTTP.
//
// This package implements a HTTP server using net/http and github.com/gorilla/mux
// taking away most of the required boiler plate code usually needed when implementing
// basic REST services. It also provides many utility methods for handling both JSON and XML responses.
package rest

import (
  "flag"
  "fmt"
  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
  "github.com/peter-mount/golib/kernel"
  "log"
  "net/http"
  "os"
  "strconv"
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
  port         *int
  // The mux Router
  router       *mux.Router
  // Base Context
  ctx          *ServerContext
}

func (a *Server) Name() string {
  return "Rest Server"
}

func (a *Server) Init( k *kernel.Kernel ) error {
  a.port = flag.Int( "rest-port", 0, "Port to use for http" )
  return nil
}

func (s *Server) PostInit() error {
  if *s.port < 1 || *s.port > 65534 {
    p, err := strconv.Atoi( os.Getenv( "RESTPORT" ) )
    if err == nil {
      *s.port = p
    }
  }
  if *s.port >0 && *s.port < 65535 {
    s.Port = *s.port
  }

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
