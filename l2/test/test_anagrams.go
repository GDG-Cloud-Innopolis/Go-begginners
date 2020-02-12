package test

import (
	"fmt"
	"testing"
)

// Anagrams is a function to test a isAnagram function
func Anagrams(IsAnagrams func(input, target string) bool) {
	testAnagrams := func(t *testing.T) {
		tests := []struct {
			first  string
			second string
			want   bool
		}{
			{
				first:  "Hello",
				second: "olleH",
				want:   true,
			},
			{
				first:  "qiu",
				second: "iuq",
				want:   true,
			},
			{
				first:  "zprl",
				second: "zprc",
				want:   false,
			},
		}
		for _, tt := range tests {
			name := fmt.Sprintf("Checking anagrams for %s, %s", tt.first, tt.second)
			t.Run(name, func(t *testing.T) {
				t.Log()
				got := IsAnagrams(tt.second, tt.first)
				if tt.want != got {
					t.Errorf("Expected: %v, got: %v", tt.want, got)
				}
			})
		}
	}
	testSuite := []testing.InternalTest{
		{
			Name: "TestAnagrams",
			F:    testAnagrams,
		},
	}

	testing.Main(matchString, testSuite, nil, nil)
}
