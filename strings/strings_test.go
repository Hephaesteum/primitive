package strings

import (
	"math"
	"math/rand"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		input, want string
	}{
		{"", ""},
		{"a", "a"},
		{"hello", "olleh"},
		{"こんにちは", "はちにんこ"},
	}

	for _, c := range cases {
		got := Reverse(c.input)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.input, got, c.want)
		}
	}
}

func TestFuzzyDistance(t *testing.T) {
	cases := []struct {
		a, b string
		want float64
	}{
		{"", "", 0.0},
		{"a", "a", 1.0},
		{"hello", "olleh", 0.2},
		{"abc", "xyz", 0.0},
		{"kitten", "sitting", 0.5714285714285714},
	}

	for _, c := range cases {
		got, _ := FuzzyDistance(c.a, c.b)
		if got != c.want {
			t.Errorf("FuzzyDistance(%q, %q) == %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func TestLengthCoord(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     string
		expected float64
	}{
		{
			name:     "short vs long",
			a:        "abc",
			b:        "xyz",
			expected: 1.0,
		},
		{
			name:     "same length",
			a:        "abcd",
			b:        "efgh",
			expected: 1.0,
		},
		{
			name:     "long vs short",
			a:        "xy",
			b:        "opqrst",
			expected: 0.7615941559557649,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := &tc.a
			b := &tc.b

			got := LengthCoord(a, b)

			if math.Abs(got-tc.expected) > 0.001 {
				t.Errorf("LengthCoord(%q, %q) == %v, want %v", *a, *b, got, tc.expected)
			}
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	// Create a random string of length n for benchmarking.
	n := 100000
	str := make([]rune, n)
	for i := 0; i < n; i++ {
		str[i] = rune(rand.Intn(26) + 'a')
	}
	input := string(str)

	// Run the function n times and record the average time per call.
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Reverse(input)
	}
}

func BenchmarkFuzzyDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FuzzyDistance("kitten", "sitting")
	}
}

func BenchmarkLengthCoord(b *testing.B) {
	x := "abc"
	y := "xyz"
	for i := 0; i < b.N; i++ {
		LengthCoord(&x, &y)
	}
}
