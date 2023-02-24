package strings

import "testing"

func TestLevenshtein(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"kitten", "sitting", 3},
		{"", "", 0},
		{"abc", "def", 3},
		{"abcdefg", "abcdeg", 1},
	}

	for _, c := range cases {
		got := Levenshtein(c.a, c.b)
		if got != c.want {
			t.Errorf("Levenshtein(%q, %q) == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func BenchmarkLevenshtein(b *testing.B) {
	a := "kitten"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Levenshtein(a, a)
	}
}
