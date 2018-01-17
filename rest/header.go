package rest

// AddHeader adds a header to the response, replacing any existing entry
func (r *Rest) AddHeader( name string, value string ) *Rest {
  r.headers[ name ] = value
  return r
}

// GetHeader returns a header from the request or "" if not present
func (r *Rest) GetHeader( name string ) string {
  return r.request.Header.Get( name )
}
