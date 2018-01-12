package codec

import (
  "time"
)

// Write allows a struct to write to the BinaryCodec as long as it implements
// Write(*BinaryCodec)
func (c *BinaryCodec) Write( i interface{ Write(*BinaryCodec) } ) *BinaryCodec {
  if c.err == nil {
    i.Write( c )
  }
  return c
}

// WriteByte writes a single byte
func (c *BinaryCodec) WriteByte( b byte ) *BinaryCodec {
  if c.err == nil {
    c.err = c.buf.WriteByte( b )
  }
  return c
}

// WriteBytes writes a byte array.
// The underlying storage is an int16 containing the length followed by the array.
func (c *BinaryCodec) WriteBytes( b []byte ) *BinaryCodec {
  if c.err == nil {
    c.WriteInt16( int16( len( b ) ) )
  }
  if c.err == nil {
    _, c.err = c.buf.Write( b )
  }
  return c
}

// WriteString writes a string
// This is the same as calling WriteBytes( []byte( string ) ) so it's an int16
// containing the length of the strings byte representation followed by the bytes.
func (c *BinaryCodec) WriteString( s string ) *BinaryCodec {
  if c.err == nil {
    c.WriteBytes( []byte( s ) )
  }
  return c
}

// WriteStringArray writes an array of strings.
// The underlying storage is an int16 containing the number of entries in the
// array and a string written by WriteString(string) for each entry.
func (c *BinaryCodec) WriteStringArray( s []string ) *BinaryCodec {
  if c.err == nil {
    c.WriteInt16( int16( len( s ) ) )
    for _, v := range s {
      c.WriteString( v )
    }
  }
  return c
}

// WriteInt writes an int.
// This is the same as WriteInt64( int64( int ) )
func (c *BinaryCodec) WriteInt( i int ) *BinaryCodec {
  return c.WriteInt64( int64(i) )
}

// WriteInt64 writes an int64.
func (c *BinaryCodec) WriteInt64( v int64 ) *BinaryCodec {
  if c.err == nil {
    var b []byte = make( []byte, 8 )
    b[0] = byte(v)
  	b[1] = byte(v >> 8)
  	b[2] = byte(v >> 16)
  	b[3] = byte(v >> 24)
    b[4] = byte(v >> 32)
  	b[5] = byte(v >> 40)
  	b[6] = byte(v >> 48)
  	b[7] = byte(v >> 56)
    _, c.err = c.buf.Write( b )
  }
  return c
}

// WriteInt32 writes an int32
func (c *BinaryCodec) WriteInt32( v int32 ) *BinaryCodec {
  if c.err == nil {
    var b []byte = make( []byte, 4 )
    b[0] = byte(v)
  	b[1] = byte(v >> 8)
  	b[2] = byte(v >> 16)
  	b[3] = byte(v >> 24)
    _, c.err = c.buf.Write( b )
  }
  return c
}

// WriteInt16 writes an int16
func (c *BinaryCodec) WriteInt16( v int16 ) *BinaryCodec {
  if c.err == nil {
    var b []byte = make( []byte, 2 )
    b[0] = byte(v)
  	b[1] = byte(v >> 8)
    _, c.err = c.buf.Write( b )
  }
  return c
}

// WriteBool writes a bool.
// The underlying storage is a single byte.
func (c *BinaryCodec) WriteBool( b bool ) *BinaryCodec {
  if c.err == nil {
    if b {
      c.err = c.buf.WriteByte( 0 )
    } else {
      c.err = c.buf.WriteByte( 1 )
    }
  }
  return c
}

// WriteTime writes a time.Time
func (c *BinaryCodec) WriteTime( t time.Time ) *BinaryCodec {
  if c.err == nil {
    if b, err := t.MarshalBinary(); err != nil {
      c.err = err
    } else {
      c.WriteBytes( b )
    }
  }
  return c
}
