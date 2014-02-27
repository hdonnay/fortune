package fortune

import (
	"os"
	"strings"
	"testing"
)

var test = `a
%
b
%
asdfasdf`

func TestFortune(t *testing.T) {
	r := strings.NewReader(test)
	s := NewScanner(r)
	for s.Scan() {
		t.Log(s.Text())
	}
}

func TestZeroLen(t *testing.T) {
	s := NewScanner(strings.NewReader(""))
	for s.Scan() {
		t.Log(s.Text())
	}
}

func TestNoSep(t *testing.T) {
	s := NewScanner(strings.NewReader("blahblahblah"))
	for s.Scan() {
		t.Log(s.Text())
	}
}

func TestRandomUse(t *testing.T) {
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
}
