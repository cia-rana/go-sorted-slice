package sortedslice

import (
	"fmt"
	"sort"
)

// String is a sorted slice of string.
type String struct {
	values []string
}

// NewStrings creates a sorted slice of string.
func NewString() *String {
	return &String{[]string{}}
}

// NewStringWithStrings creates a sorted slice of string using the given slice of string.
func NewStringWithStrings(a []string) *String {
	return &String{duplicateStrings(a)}
}

// Insert inserts a value x into the sorted slice of string. If uniq is true and the sorted slice already has x, then Insert dosen't insert x.
func (s *String) Insert(x string, uniq bool) {
	t := s.values
	index := sort.SearchStrings(t, x)
	if !uniq || len(t) == index || t[index] != x {
		s.values = append(t[:index], append([]string{x}, t[index:]...)...)
		return
	}
}

// Delete deletes x that be contained in String.
func (s *String) Delete(x string) {
	t := s.values
	for {
		index := sort.SearchStrings(t, x)
		if len(t) == index || t[index] != x {
			break
		}
		t = append(t[:index], t[index+1:]...)
	}
	s.values = t
}

// Strings returns a sorted slice as a slice of string.
func (s String) Strings() []string {
	return duplicateStrings(s.values)
}

// Uniq removes duplicates.
func (s *String) Uniq() {
	m := make(map[string]struct{})

	for _, v := range s.values {
		m[v] = struct{}{}
	}

	b := make([]string, len(m), len(m))
	i := 0
	for k, _ := range m {
		b[i] = k
		i++
	}

	s.values = b
}

// String creates string made from sorted slice's values.
func (s String) String() string {
	return fmt.Sprint(s.values)
}

// At returns i th value.
func (s String) At(i int) string {
	return s.values[i]
}

// Len return length of values.
func (s String) Len() int {
	return len(s.values)
}

func duplicateStrings(a []string) []string {
	b := make([]string, len(a), len(a))
	for i, v := range a {
		b[i] = v
	}
	return b
}
