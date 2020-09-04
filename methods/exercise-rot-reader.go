package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, e := rot.r.Read(b)

	for i, c := range b {
		if 'A' <= c && c < ('Z' - 13) || 'a' <= c && c < ('z' - 13) {
			b[i] += 13
		} else if ('Z' - 13) <= c && c <= 'Z' || ('z' - 13) <= c && c <= 'z' {
			b[i] -= 13
		}
	}

	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

