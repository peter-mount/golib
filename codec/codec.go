// A simple codec used to read or write binary data in a compact format.
//
// For an example of this in use see the nrod-cif repository
//
package codec

import (
  "bytes"
)

// An instance of a BinaryCodec
type BinaryCodec struct {
  buf  *bytes.Buffer
  err   error
}

// NewBinaryCodec creates a BinaryCodec instance that can be used to
// marshal data to
func NewBinaryCodec() *BinaryCodec {
  var c *BinaryCodec = &BinaryCodec{}
  c.buf = &bytes.Buffer{}
  return c
}

// NewBinaryCodecFrom creates a BinaryCodec instance based on a []byte used to
// unmarshal from
func NewBinaryCodecFrom( b []byte ) *BinaryCodec {
  var c *BinaryCodec = &BinaryCodec{}
  buf := bytes.NewBuffer( b )
  c.buf = buf
  return c
}

// Bytes returns the byte array of the underlying buffer.
func (c *BinaryCodec) Bytes() []byte {
  return c.buf.Bytes()
}

// Error returns any error incurred during marshalling or nil if everything was ok.
func (c *BinaryCodec) Error() error {
  return c.err
}
