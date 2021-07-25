package interfaces

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader 
}

func (rot13 *rot13Reader) Read(buffer []byte) (int, error) {
	ibuff := make([]byte, len(buffer))
	n, err := rot13.r.Read(ibuff)
	if err == io.EOF {
		return 0, io.EOF
	}
	for i := 0; i < n; i++ {
		buffer[i] = RotateCharacter(ibuff[i])
		fmt.Printf("%v, %v\n", ibuff[i], buffer[i])
	}
	return n, nil
}
func RotateCharacter(r byte) byte {
	if r >= 'a' + 13 && r <= 'z' + 13 {
		// Rotate lowercase letters 13 places.
		if r >= 'm' + 13 {
				return r + 13
		} else {
				return r - 13
		}
	} else if r >= 'A' + 13 && r <= 'Z' + 13 {
			// Rotate uppercase letters 13 places.
			if r >= 'M' + 13 {
					return r + 13
			} else {
					return r - 13
			}
	}
	// Do nothing.
	return r	
}

func interfaces16() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println("\ntry again")
	io.Copy(os.Stdout, &r)
	fmt.Println()
}