package utils

import "testing"

func TestStripChars(t *testing.T) {
	testCases := []struct {
		label    string
		str      string
		chars    string
		expected string
	}{
		{
			label:    "Empty input string, empty chars",
			str:      "",
			chars:    "",
			expected: "",
		},
		{
			label:    "Empty input string, non-empty chars",
			str:      "",
			chars:    "12345",
			expected: "",
		},
		{
			label:    "Non-empty input string, empty chars",
			str:      "This is a test.",
			chars:    "",
			expected: "This is a test.",
		},
		{
			label:    "Regular use case",
			str:      "Does this work!? @Or doesn't it?!",
			chars:    "!@",
			expected: "Does this work? Or doesn't it?",
		},
	}

	for _, tc := range testCases {
		result := StripChars(tc.str, tc.chars)
		if result != tc.expected {
			t.Fatalf("Test \"%s\" expected \"%s\". Got \"%s\".", tc.label, tc.expected, result)
		}
	}
}

func TestMin(t *testing.T) {
	testCases := []struct {
		label    string
		x        int
		y        int
		expected int
	}{
		{
			label:    "x lower",
			x:        0,
			y:        4,
			expected: 0,
		},
		{
			label:    "y lower",
			x:        5,
			y:        2,
			expected: 2,
		},
		{
			label:    "Both equal",
			x:        6,
			y:        6,
			expected: 6,
		},
	}

	for _, tc := range testCases {
		result := Min(tc.x, tc.y)
		if result != tc.expected {
			t.Fatalf("Test \"%s\" expected %d. Got %d.", tc.label, tc.expected, result)
		}
	}
}
