package codec

import (
  "bytes"
)

type BinaryCodec struct {
  buf  *bytes.Buffer
  err   error
}

func NewBinaryCodec() *BinaryCodec {
  var c *BinaryCodec = &BinaryCodec{}
  c.buf = &bytes.Buffer{}
  return c
}

func NewBinaryCodecFrom( b []byte ) *BinaryCodec {
  var c *BinaryCodec = &BinaryCodec{}
  buf := bytes.NewBuffer( b )
  c.buf = buf
  return c
}
