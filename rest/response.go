package rest

import (
  "encoding/json"
  "encoding/xml"
)

const (
  APPLICATION_JSON string = "application/json"
  APPLICATION_XML string = "application/xml"
  TEXT_JSON string = "text/json"
  TEXT_XML string = "text/xml"
)

// Send returns data to the client.
// If the request has the Accept header of "text/xml" or "application/xml" then
// the response will be in XML, otherwise in JSON.
func (r *Rest) Send() {
  if( !r.sent ) {

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

    // Force the Content-Type
    r.AddHeader( "Content-Type", accept )

    // Write the headers
    h := r.writer.Header()
    for k, v := range r.headers {
      h.Add( k, v)
    }

    // Write the status
    r.writer.WriteHeader( r.status )

    // Finally the content
    if r.value != nil {
      if isXml {
        xml.NewEncoder( r.writer ).Encode( r.value )
      } else {
        json.NewEncoder( r.writer ).Encode( r.value )
      }
    }

    r.sent = true
  }
}
