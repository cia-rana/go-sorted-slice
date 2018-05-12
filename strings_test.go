package sortedslice

import (
	"reflect"
	"sort"
	"testing"
)

var a = []string{"l", "m", "n", "ll", "lm", "ln", "ml", "mm", "mn", "nl", "nm", "nn", "a", "k", "o", "z", "n", "m", "l"}

func TestInsertStrings(t *testing.T) {
	var b []string
	var s *String

	s = NewString()
	b = duplicateStrings(a)
	for _, v := range b {
		s.Insert(v, true)
		t.Log(s)
	}

	b = uniqStrings(b)
	sort.Strings(b)
	if !reflect.DeepEqual(b, s.values) {
		t.Fatalf("expected: %s\nactual: %s\n", b, s.values)
	}

	s = NewString()
	b = duplicateStrings(a)
	for _, v := range b {
		s.Insert(v, false)
	}

	sort.Strings(b)
	if !reflect.DeepEqual(b, s.values) {
		t.Fatalf("expected: %s\nacual: %s\n", b, s.values)
	}
}

func uniqStrings(a []string) []string {
	m := make(map[string]struct{})
	for _, v := range a {
		m[v] = struct{}{}
	}

	b := make([]string, len(m), len(m))
	i := 0
	for k, _ := range m {
		b[i] = k
		i++
	}

	return b
}
