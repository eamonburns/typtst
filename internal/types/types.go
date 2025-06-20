package types

import "fmt"

type TypingError struct {
	Expected rune
	Actual   rune
}

func (self TypingError) Format(s fmt.State, _ rune) {
	fmt.Fprintf(s, "TypingError{Expected: '%c', Actual: '%c'}", self.Expected, self.Actual)
}
