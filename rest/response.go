package rest

import (
  "encoding/json"
  "encoding/xml"
  "errors"
  "log"
)

const (
  APPLICATION_JSON string = "application/json"
  APPLICATION_XML string = "application/xml"
  TEXT_JSON string = "text/json"
  TEXT_HTML string = "text/html"
  TEXT_XML string = "text/xml"
)

var (
  resposeUsed = errors.New( "Response already written to" )
)

// Send returns data to the client.
// If the request has the Accept header of "text/xml" or "application/xml" then
// the response will be in XML, otherwise in JSON.
func (r *Rest) Send() error {
  if( r.sent ) {
    return resposeUsed
  }

  r.sent = true

  if r.status <= 0 {
    r.status = 200
  }

  accept := r.GetHeader( "Accept" )
  isXml := accept == TEXT_XML || accept == APPLICATION_XML
  isJson := accept == TEXT_JSON || accept == APPLICATION_JSON

  // Ensure we have a valid contentType default to APPLICATION_JSON if not
  if !isXml && !isJson {
    accept = APPLICATION_JSON
    isJson = true
  }

  // Force the Content-Type if the response contentType is not set
  if r.contentType == "" {
    r.AddHeader( "Content-Type", accept )
  }else {
    r.AddHeader( "Content-Type", r.contentType )
  }

  // Until we get CORS handling correctly
  r.AddHeader( "Access-Control-Allow-Origin", "*" )

  // Write the headers
  h := r.writer.Header()
  for k, v := range r.headers {
    h.Add( k, v)
  }

  // Write the status
  r.writer.WriteHeader( r.status )

  // Write from a reader
  if r.source != nil {
    err := r.source( r.writer )
    log.Printf( "copy: %v %v", err )
    return err
  } else if r.value != nil {
    // Finally the content, encode if an object
    if isXml {
      return xml.NewEncoder( r.writer ).Encode( r.value )
    } else {
      return json.NewEncoder( r.writer ).Encode( r.value )
    }
  }

  return nil
}
