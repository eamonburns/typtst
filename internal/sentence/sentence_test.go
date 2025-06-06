package sentence_test

import (
	"slices"
	"testing"

	"github.com/agent-e11/typtst/internal/sentence"
)

func TestSplit(t *testing.T) {
	s := []rune("This is a test.")
	expected := []sentence.Token{
		{
			T:      sentence.WordToken,
			String: "This",
		},
		{
			T:      sentence.SpaceToken,
			String: " ",
		},
		{
			T:      sentence.WordToken,
			String: "is",
		},
		{
			T:      sentence.SpaceToken,
			String: " ",
		},
		{
			T:      sentence.WordToken,
			String: "a",
		},
		{
			T:      sentence.SpaceToken,
			String: " ",
		},
		{
			T:      sentence.WordToken,
			String: "test",
		},
		{
			T:      sentence.PunctuationToken,
			String: ".",
		},
	}
	actual := sentence.Split(s)

	if !slices.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}

	s = []rune("It's called \"pizza\"?")
	expected = []sentence.Token{
		{
			T:      sentence.WordToken,
			String: "It",
		},
		{
			T:      sentence.PunctuationToken,
			String: "'",
		},
		{
			T:      sentence.WordToken,
			String: "s",
		},
		{
			T:      sentence.SpaceToken,
			String: " ",
		},
		{
			T:      sentence.WordToken,
			String: "called",
		},
		{
			T:      sentence.SpaceToken,
			String: " ",
		},
		{
			T:      sentence.PunctuationToken,
			String: "\"",
		},
		{
			T:      sentence.WordToken,
			String: "pizza",
		},
		{
			T:      sentence.PunctuationToken,
			String: "\"",
		},
		{
			T:      sentence.PunctuationToken,
			String: "?",
		},
	}

	actual = sentence.Split(s)

	if !slices.Equal(expected, actual) {
		t.Fatalf("Expected %v, got %v", expected, actual)
	}

	s = []rune("(Parentheses)")
	expected = []sentence.Token{
		{
			T:      sentence.PunctuationToken,
			String: "(",
		},
		{
			T:      sentence.WordToken,
			String: "Parentheses",
		},
		{
			T:      sentence.PunctuationToken,
			String: ")",
		},
	}
	actual = sentence.Split(s)

	if !slices.Equal(expected, actual) {
		t.Fatalf("\nExpected: %v\nGot:      %v", expected, actual)
	}
}
