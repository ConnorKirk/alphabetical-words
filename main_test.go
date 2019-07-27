package main

import "testing"

func TestIsAlphabetical(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"abcdefg", true},
		{"bacde", false},
		{"aabc", true},
	}

	for _, tc := range tests {
		got := isAlphabetical(tc.word)
		if got != tc.want {
			t.Errorf("got=%v; want=%v", got, tc.want)
		}
	}
}
