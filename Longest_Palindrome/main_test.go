package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase(t *testing.T) {
	testcase := map[string]string{
		"abaxyzzyxf":                        "xyzzyx",
		"a":                                 "a",
		`it's highnoon`:                     "noon",
		`noon high it is`:                   "noon",
		"abccbait's highnoon":               "abccba",
		"abcdefgfedcbazzzzzzzzzzzzzzzzzzzz": "zzzzzzzzzzzzzzzzzzzz",
		"abcdefgfedcba":                     "abcdefgfedcba",
		"abcdefghfedcbaa":                   "aa",
		"abcdefggfedcba":                    "abcdefggfedcba",
		"zzzzzzz2345abbbba5432zzbbababa":    "zz2345abbbba5432zz",
		"z234a5abbbba54a32z":                "5abbbba5",
		"z234a5abbba54a32z":                 "5abbba5",
		"ab12365456321bb":                   "b12365456321b",
	}

	for i := range testcase {
		t.Run(i, func(t *testing.T) {
			output := LongestPalindromicSubstring(i)
			assert.Equal(t, testcase[i], output)
		})
	}
}
