package util

import (
  "bytes"
  "strconv"
)

type BufferWriter struct {
  b             bytes.Buffer
  FieldWidth    int
}

func (b *BufferWriter) String() string {
  return b.b.String()
}

func (b *BufferWriter) Field( n string, s string ) *BufferWriter {
  if b.FieldWidth <=0 {
    b.FieldWidth = 15
  }

  if len( n ) <= b.FieldWidth {
    for i := len( n ); i < b.FieldWidth; i++ {
      b.b.WriteString( " " )
    }

    b.b.WriteString( n )
  } else {
    b.b.WriteString( n[0:b.FieldWidth] )
  }

  b.b.WriteString( " " )
  b.b.WriteString( s )
  b.b.WriteString( "\n" )
  return b
}

func (b *BufferWriter) FieldInt( n string, s int) *BufferWriter {
  b.Field( n, strconv.FormatInt( int64( s ), 10 ) )
  return b
}

func (b *BufferWriter) FieldBool( n string, s bool) *BufferWriter {
  if s {
    b.Field( n, "true" )
  } else {
    b.Field( n, "false" )
  }
  return b
}

func (b *BufferWriter) Row( ) *BufferWriter {
  b.b.WriteString( "\n| " )
  return b
}

func (b *BufferWriter) Pad( s string, l int ) *BufferWriter {
  if len( s ) <= l {
    b.b.WriteString( s )
  } else {
    b.b.WriteString( s[0:l] )
  }
  for i := len( s ); i < l; i++ {
    b.b.WriteString( " " )
  }
  b.b.WriteString( " | " )
  return b
}

func (b *BufferWriter)PadX( l int ) *BufferWriter {
  b.Pad( "", l )
  return b
}

func  (b *BufferWriter)PadBool( s bool, l int ) *BufferWriter {
  if s {
    b.Pad( "t", l )
  } else {
    b.Pad( "f", l )
  }
  return b
}

func (b *BufferWriter) PadInt( s int, l int ) *BufferWriter {
  b.Pad( strconv.FormatInt( int64( s ), 10 ), l )
  return b
}
