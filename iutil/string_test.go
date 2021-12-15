package iutil

import "testing"

func TestTrim(t *testing.T) {
	s := "hello\nworld\t! !"
	t.Log(Trim(s))
}
