package main

// import "golang.org/x/tour/reader"

type MyReader struct{}
type MyReaderError bool

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (myReaderError MyReaderError) Error() string {
	return "error in byte[]"
}

func (my MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, MyReaderError(true)
	}
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
