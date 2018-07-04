package rest

import (
  "github.com/gorilla/mux"
  "io"
  "net/http"
)

// A Rest query. This struct handles everything from the inbound request and
// sending the response
type Rest struct {
  writer        http.ResponseWriter
  request      *http.Request
  // Response contentType
  contentType   string
  // Response HTTP Status code, defaults to 200
  status        int
  // The value to send
  value         interface{}
  reader        io.Reader
  // Response headers
  headers       map[string]string
  // true if Send() has been called
  sent          bool
  // Request route variables
  vars          map[string]string
  // The context
  context       string
}

// NewRest creates a new Rest query
func NewRest( writer http.ResponseWriter, request *http.Request) *Rest {
  r := &Rest{}
  r.writer = writer
  r.request = request
  r.headers = make( map[string]string )
  return r
}

// Request return the underlying http.Request so that
func (r *Rest) Request() *http.Request {
  return r.request
}

// Var returns the named route variable or "" if none
func (r *Rest) Var(name string) string {
  if r.vars == nil {
    r.vars = mux.Vars( r.request )
  }
  if r.vars == nil {
    return ""
  }
  return r.vars[ name ]
}

// Status sets the HTTP status of the response.
func (r *Rest) Status( status int ) *Rest {
  r.status = status
  return r
}

// Value sets the response value
func (r *Rest) Value( value interface{} ) *Rest {
  r.value = value
  return r
}

// Value sets the response value
func (r *Rest) Reader( rdr io.Reader ) *Rest {
  r.reader = rdr
  return r
}

// Value sets the response value
func (r *Rest) ContentType( c string ) *Rest {
  r.contentType = c
  return r
}

func (r *Rest) HTML() *Rest { return r.ContentType( TEXT_HTML ) }
func (r *Rest) JSON() *Rest { return r.ContentType( APPLICATION_JSON ) }
func (r *Rest) XML() *Rest { return r.ContentType( APPLICATION_XML ) }

// Context returns the base context for this request
func (r *Rest) Context() string {
  return r.context
}
