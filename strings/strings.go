package strings

import (
	"errors"
	"math"
	"unicode"
)

// Reverse returns a new string with characters of the input string `a` in
// reverse order. Original string is not modified. A slice of runes holds the
// characters of the input string, and swaps corresponding characters at
// opposite ends of slice until reaching the middle of the slice. The returned
// string is a new string created from the slice of runes.
func Reverse(a string) string {
	// Convert input string to slice of runes for efficient manipulation.
	runes := []rune(a)

	// Swap corresponding runes at opposite ends of the slice until middle.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Create new string from reversed slice of runes.
	return string(runes)
}

// FuzzyDistance computes the fuzzy distance between two input strings `a` and
// `b` using the Levenshtein distance algorithm. The fuzzy distance is defined
// as the proportion of characters in the longer string that are not part of the
// Levenshtein edit distance.
func FuzzyDistance(a, b string) (float64, error) {
	if IsBlank(a) && IsBlank(b) {
		return 0.0, errors.New("unable to compare two empty strings")
	}
	// Compute the maximum length of the input strings.
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	// Compute the Levenshtein edit distance between the input strings.
	edit := Levenshtein(a, b)

	// Compute the fuzzy distance as the proportion of characters in the longer
	// string that are not part of the edit distance.
	return float64(maxLen-edit) / float64(maxLen), nil
}

// LengthCoord takes two string pointers and computes a length coordinate that
// captures the difference in length between the two strings.
//
// If the two strings have the same length the length coordinate is set to 1.0
// indicating that the two strings are equal length and there is no difference
// in length to consider.
//
// If two strings have different lengths the function computes the length ratio
// between the shorter and longer strings. The length ratio is then multiplied
// by 3 to amplify the effect of the difference in length, and the hyperbolic
// tangent function is applied to obtain the final length coordinate.
func LengthCoord(a, b *string) float64 {
	// If strings same length, length coordinate is 1.0.
	if len(*a) == len(*b) {
		return 1.0
	}

	// Compute length ratio between shorter and longer strings.
	var shortLen, longLen int
	if len(*a) > len(*b) {
		shortLen = len(*b)
		longLen = len(*a)
	} else {
		shortLen = len(*a)
		longLen = len(*b)
	}
	lenRatio := float64(shortLen) / float64(longLen)

	// Compute length coordinate using hyperbolic tangent function.
	// Multiply length ratio by 3 to amplify the effect of difference in length.
	return math.Tanh(lenRatio * 3)
}

// IsBlank returns true if a given string contains only whitespace characters
// (spaces, tabs, newlines, etc.) or is empty.
func IsBlank(s string) bool {
	// Loop over each rune in the string.
	for _, r := range s {
		// If the rune is not a whitespace character, return false.
		if !unicode.IsSpace(r) {
			return false
		}
	}
	// If we made it through the whole loop without returning, the string is blank.
	return true
}
