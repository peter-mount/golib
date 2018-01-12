# codec
--
    import "github.com/peter-mount/golib/codec"

A simple codec used to read or write binary data in a compact format.

For an example of this in use see the nrod-cif repository

## Usage

#### type BinaryCodec

```go
type BinaryCodec struct {
}
```

An instance of a BinaryCodec

#### func  NewBinaryCodec

```go
func NewBinaryCodec() *BinaryCodec
```
NewBinaryCodec creates a BinaryCodec instance that can be used to marshal data
to

#### func  NewBinaryCodecFrom

```go
func NewBinaryCodecFrom(b []byte) *BinaryCodec
```
NewBinaryCodecFrom creates a BinaryCodec instance based on a []byte used to
unmarshal from

#### func (*BinaryCodec) Bytes

```go
func (c *BinaryCodec) Bytes() []byte
```
Bytes returns the byte array of the underlying buffer.

#### func (*BinaryCodec) Error

```go
func (c *BinaryCodec) Error() error
```
Error returns any error incurred during marshalling or nil if everything was ok.

#### func (*BinaryCodec) Read

```go
func (c *BinaryCodec) Read(i interface {
	Read(*BinaryCodec)
}) *BinaryCodec
```
Read allows a struct to read from a BinraryCodec as long as it implements the
Read(*BinaryCodec) function.

#### func (*BinaryCodec) ReadBool

```go
func (c *BinaryCodec) ReadBool(i *bool) *BinaryCodec
```
ReadBool reads a bool. The underlying storage is a single byte

#### func (*BinaryCodec) ReadByte

```go
func (c *BinaryCodec) ReadByte(b *byte) *BinaryCodec
```
ReadByte reads a single byte

#### func (*BinaryCodec) ReadBytes

```go
func (c *BinaryCodec) ReadBytes(b *[]byte) *BinaryCodec
```
ReadBytes reads a byte slice

#### func (*BinaryCodec) ReadInt

```go
func (c *BinaryCodec) ReadInt(i *int) *BinaryCodec
```
ReadInt reads an integer. The underlying storage is 64bits in length.

#### func (*BinaryCodec) ReadInt16

```go
func (c *BinaryCodec) ReadInt16(i *int16) *BinaryCodec
```
ReadInt16 reads an int16

#### func (*BinaryCodec) ReadInt32

```go
func (c *BinaryCodec) ReadInt32(i *int32) *BinaryCodec
```
ReadInt32 reads an int32

#### func (*BinaryCodec) ReadInt64

```go
func (c *BinaryCodec) ReadInt64(i *int64) *BinaryCodec
```
ReadInt64 reads an int64

#### func (*BinaryCodec) ReadString

```go
func (c *BinaryCodec) ReadString(s *string) *BinaryCodec
```
ReadString reads a string

#### func (*BinaryCodec) ReadStringArray

```go
func (c *BinaryCodec) ReadStringArray(s *[]string) *BinaryCodec
```
ReadStringArray reads an array of strings

#### func (*BinaryCodec) ReadTime

```go
func (c *BinaryCodec) ReadTime(i *time.Time) *BinaryCodec
```
ReadTime reads a time.Time value

#### func (*BinaryCodec) Write

```go
func (c *BinaryCodec) Write(i interface {
	Write(*BinaryCodec)
}) *BinaryCodec
```
Write allows a struct to write to the BinaryCodec as long as it implements
Write(*BinaryCodec)

#### func (*BinaryCodec) WriteBool

```go
func (c *BinaryCodec) WriteBool(b bool) *BinaryCodec
```
WriteBool writes a bool. The underlying storage is a single byte.

#### func (*BinaryCodec) WriteByte

```go
func (c *BinaryCodec) WriteByte(b byte) *BinaryCodec
```
WriteByte writes a single byte

#### func (*BinaryCodec) WriteBytes

```go
func (c *BinaryCodec) WriteBytes(b []byte) *BinaryCodec
```
WriteBytes writes a byte array. The underlying storage is an int16 containing
the length followed by the array.

#### func (*BinaryCodec) WriteInt

```go
func (c *BinaryCodec) WriteInt(i int) *BinaryCodec
```
WriteInt writes an int. This is the same as WriteInt64( int64( int ) )

#### func (*BinaryCodec) WriteInt16

```go
func (c *BinaryCodec) WriteInt16(v int16) *BinaryCodec
```
WriteInt16 writes an int16

#### func (*BinaryCodec) WriteInt32

```go
func (c *BinaryCodec) WriteInt32(v int32) *BinaryCodec
```
WriteInt32 writes an int32

#### func (*BinaryCodec) WriteInt64

```go
func (c *BinaryCodec) WriteInt64(v int64) *BinaryCodec
```
WriteInt64 writes an int64.

#### func (*BinaryCodec) WriteString

```go
func (c *BinaryCodec) WriteString(s string) *BinaryCodec
```
WriteString writes a string This is the same as calling WriteBytes( []byte(
string ) ) so it's an int16 containing the length of the strings byte
representation followed by the bytes.

#### func (*BinaryCodec) WriteStringArray

```go
func (c *BinaryCodec) WriteStringArray(s []string) *BinaryCodec
```
WriteStringArray writes an array of strings. The underlying storage is an int16
containing the number of entries in the array and a string written by
WriteString(string) for each entry.

#### func (*BinaryCodec) WriteTime

```go
func (c *BinaryCodec) WriteTime(t time.Time) *BinaryCodec
```
WriteTime writes a time.Time
