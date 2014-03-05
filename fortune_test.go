package fortune

import (
	"crypto/rand"
	"math/big"
	"os"
	"strings"
	"testing"
)

var (
	test = `a
%
b
%
asdfasdf
`
	testCheck = []string{"a\n", "b\n", "asdfasdf\n"}
)

func TestFortune(t *testing.T) {
	r := strings.NewReader(test)
	s := NewScanner(r)
	for i := 0; s.Scan(); i++ {
		if o := s.Text(); o != testCheck[i] {
			t.Fatalf("%s != %s\n", o, testCheck[i])
		}
	}
}

func TestZeroLen(t *testing.T) {
	s := NewScanner(strings.NewReader(""))
	s.Scan()
	if o := s.Text(); o != "" {
		t.Fatalf("'' != '%s'\n", o)
	}
}

func TestNoSep(t *testing.T) {
	text := "blahblahblah"
	s := NewScanner(strings.NewReader(text))
	s.Scan()
	if o := s.Text(); o != text {
		t.Fatalf("%s != %s\n", o, text)
	}
}

func TestRandomUse(t *testing.T) {
	r, err := os.Open("test/test.fortune")
	if err != nil {
		t.Fatal(err)
	}
	x, _ := rand.Int(rand.Reader, big.NewInt(Count(r)))
	rn := x.Int64()
	r.Seek(0, 0)
	s := NewScanner(r)
	for i := 0; s.Scan(); i++ {
		if int64(i) == rn {
			t.Log(s.Text())
			return
		}
	}
	t.Fatal("exhausted loop")
}
