package codec

import (
  "time"
)

// Read allows a struct to read from a BinraryCodec as long as it implements
// the Read(*BinaryCodec) function.
func (c *BinaryCodec) Read( i interface{ Read(*BinaryCodec) } ) *BinaryCodec {
  if c.err == nil {
    i.Read( c )
  }
  return c
}

// ReadByte reads a single byte
func (c *BinaryCodec) ReadByte( b *byte ) *BinaryCodec {
  if c.err == nil {
    if t, err := c.buf.ReadByte(); err != nil {
      c.err = err
    } else {
      *b = t
    }
  }
  return c
}

// ReadBytes reads a byte slice
func (c *BinaryCodec) ReadBytes( b *[]byte ) *BinaryCodec {
  var l int16
  if c.err == nil {
    c.ReadInt16( &l )
  }
  if c.err == nil {
    ary := make( []byte, l )
    if _, err := c.buf.Read( ary ); err != nil {
      c.err = err
    } else {
      *b = ary
    }
  }
  return c
}

// ReadString reads a string
func (c *BinaryCodec) ReadString( s *string) *BinaryCodec {
  var b []byte
  if c.err == nil {
    c.ReadBytes( &b )
  }
  if c.err == nil {
    *s = string(b[:])
  }
  return c
}

// ReadStringArray reads an array of strings
func (c *BinaryCodec) ReadStringArray( s *[]string) *BinaryCodec {
  var l int16
  if c.err == nil {
    c.ReadInt16( &l )
  }
  if c.err == nil {
    var a []string = make( []string, l )
    for i := 0; i < int(l); i++ {
      c.ReadString( &(a[i]) )
    }
    if c.err == nil {
      *s = a
    }
  }
  return c
}

// ReadInt reads an integer. The underlying storage is 64bits in length.
func (c *BinaryCodec) ReadInt( i *int ) *BinaryCodec {
  var t int64
  c.ReadInt64( &t )
  if( c.err == nil ) {
    *i = int(t)
  }
  return c
}

// ReadInt64 reads an int64
func (c *BinaryCodec) ReadInt64( i *int64) *BinaryCodec {
  if c.err == nil {
    var b []byte = make( []byte, 8 )
    if _, err := c.buf.Read( b ); err != nil {
      c.err = err
    } else {
      *i = int64( uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
  		  uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56 )
    }
  }
  return c
}

// ReadInt32 reads an int32
func (c *BinaryCodec) ReadInt32( i *int32 ) *BinaryCodec {
  if c.err == nil {
    var b []byte = make( []byte, 4 )
    if _, err := c.buf.Read( b ); err != nil {
      c.err = err
    } else {
      *i = int32( uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24 )
    }
  }
  return c
}

// ReadInt16 reads an int16
func (c *BinaryCodec) ReadInt16( i *int16 ) *BinaryCodec {
  if c.err == nil {
    var b []byte = make( []byte, 2 )
    if _, err := c.buf.Read( b ); err != nil {
      c.err = err
    } else {
      *i = int16( uint16(b[0]) | uint16(b[1])<<8 )
    }
  }
  return c
}

// ReadBool reads a bool. The underlying storage is a single byte
func (c *BinaryCodec) ReadBool( i *bool ) *BinaryCodec {
  if c.err == nil {
    if b, err := c.buf.ReadByte(); err != nil {
      c.err = err
    } else {
      if b == 0 {
        *i = false
      } else {
        *i = true
      }
    }
  }
  return c
}

// ReadTime reads a time.Time value
func (c *BinaryCodec) ReadTime( i *time.Time ) *BinaryCodec {
  var b []byte

  if c.err == nil {
    c.ReadBytes( &b )
  }

  if c.err == nil {
    c.err = i.UnmarshalBinary( b )
  }

  return c
}
