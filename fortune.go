/*
Package fortune sets up a bufio.Scanner for reading a fortune file.

	r, err := os.Open("test/test.fortune")
	if err != nil {
		t.Fatal(err)
	}
	r2, _ := os.Open("test/test.fortune")
	rn := Count(r2)
	s := NewScanner(r)
	for i := 0; s.Scan(); i++ {
		if int64(i) == rn {
			t.Log(s.Text())
			break
		}
	}

*/
package fortune

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"io"
	"math/big"
)

// Delim is the byte sequent that delimits fortunes
var Delim = []byte("%\n")

// SplitFortune is a bufio.SplitFunc to use on an io.Reader from a fortune file
func SplitFortune(data []byte, atEOF bool) (advance int, token []byte, err error) {
	idx := bytes.Index(data, []byte("%\n"))
	switch {
	case idx == -1 && !atEOF:
		break
	case idx == -1 && atEOF:
		advance = len(data)
		token = data
	case idx != -1:
		token = data[:idx]
		advance = idx + 2
	}
	return
}

// NewScanner returns a Scanner with Split() set already
func NewScanner(r io.Reader) *bufio.Scanner {
	ret := bufio.NewScanner(r)
	ret.Split(SplitFortune)
	return ret
}

// Count returns number of fortunes - 1
func Count(r io.Reader) int64 {
	var b bytes.Buffer
	io.Copy(&b, r)
	max := bytes.Count(b.Bytes(), Delim)
	ret, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return ret.Int64()
}
