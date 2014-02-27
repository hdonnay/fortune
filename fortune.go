package fortune

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"io"
	"math/big"
)

var Delim = []byte("%\n")

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

func NewScanner(r io.Reader) *bufio.Scanner {
	ret := bufio.NewScanner(r)
	ret.Split(SplitFortune)
	return ret
}

func Count(r io.Reader) int64 {
	var b bytes.Buffer
	io.Copy(&b, r)
	max := bytes.Count(b.Bytes(), Delim)
	ret, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return ret.Int64()
}
