package byteslice

// ByteSlice type
type ByteSlice []byte

// Append appends data and returns new bvyte slice
func (slice ByteSlice) Append(data []byte) []byte {
	return append([]byte(slice), data...)
}

// AppendPointer appends data to existing ByteSlice
func (slice *ByteSlice) AppendPointer(data []byte) {
	*slice = append([]byte(*slice), data...)
}

func (slice *ByteSlice) Write(data []byte) (n int, err error) {
	*slice = append([]byte(*slice), data...)
	return len(data), nil
}

func (slice ByteSlice) String() string { // implements fmt.Stringer
	return string(slice)
}
