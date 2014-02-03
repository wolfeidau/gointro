package strings

import (
	"testing"
)

func TestReverse(t *testing.T) {
	const in, out = "Hello, 鸡炒饭", "饭炒鸡 ,olleH"
	if x := Reverse(in); x != out {
		t.Errorf("Reverse(%s) = %s, expected %s", in, x, out)
	}
}
