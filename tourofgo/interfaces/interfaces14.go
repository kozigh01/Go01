package interfaces

import (
	"fmt"
	"io"
	"strings"
)

func interfaces14() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err = %v, b = %v, b[:n] = %q\n", n, err, b, b[:n])
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}