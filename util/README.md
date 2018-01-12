# util
--
    import "github.com/peter-mount/golib/util"


## Usage

#### type BufferWriter

```go
type BufferWriter struct {
	FieldWidth int
}
```


#### func (*BufferWriter) Field

```go
func (b *BufferWriter) Field(n string, s string) *BufferWriter
```

#### func (*BufferWriter) FieldBool

```go
func (b *BufferWriter) FieldBool(n string, s bool) *BufferWriter
```

#### func (*BufferWriter) FieldInt

```go
func (b *BufferWriter) FieldInt(n string, s int) *BufferWriter
```

#### func (*BufferWriter) Pad

```go
func (b *BufferWriter) Pad(s string, l int) *BufferWriter
```

#### func (*BufferWriter) PadBool

```go
func (b *BufferWriter) PadBool(s bool, l int) *BufferWriter
```

#### func (*BufferWriter) PadInt

```go
func (b *BufferWriter) PadInt(s int, l int) *BufferWriter
```

#### func (*BufferWriter) PadX

```go
func (b *BufferWriter) PadX(l int) *BufferWriter
```

#### func (*BufferWriter) Row

```go
func (b *BufferWriter) Row() *BufferWriter
```

#### func (*BufferWriter) String

```go
func (b *BufferWriter) String() string
```
