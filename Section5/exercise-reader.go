package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, err) method to MyReader.
func (reader MyReader) Read(array []byte) (int, error) {
    size := len(array)
    for i := 0; i<size; i++ {
        array[i] = 'A'
    }
    return size, nil
}

func main() {
    reader.Validate(MyReader{})
}
