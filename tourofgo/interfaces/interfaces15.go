package interfaces

import (
		"golang.org/x/tour/reader"
)

type MyReader struct {

}

func (r MyReader) Read(buffer []byte) (int, error) {
	for i := 0; i < len(buffer); i++ {
		buffer[i] = 'A'
	}

	return len(buffer), nil
}


func interfaces15() {
	reader.Validate(MyReader{})
}