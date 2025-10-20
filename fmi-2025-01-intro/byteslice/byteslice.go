package byteslice

type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
	return append([]byte(slice), data...)
}
func (slice *ByteSlice) AppendPointer(data []byte) {
	*slice = append([]byte(*slice), data...)
}
func (slice *ByteSlice) Write(data []byte) (n int, err error) {
	*slice = append([]byte(*slice), data...)
	return len(data), nil
}

func New(size int) *ByteSlice {
	byteslice := ByteSlice(make([]byte, size))
	return &byteslice
}
