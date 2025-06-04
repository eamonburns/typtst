package sentence

import (
	"fmt"
	"log"
	"unicode"
)

type Sentence []Token

// NOTE: Should this return an error?

func (self Sentence) Render(cursorIdx int, errors map[int]bool, maxWidth int) RenderedSentence {
	log.Printf("> Sentence.Render(cursorIdx: %v, errors: %v, maxWidth: %v)", cursorIdx, errors, maxWidth)
	log.Printf("self: %v", self)
	linePre := TypedColor
	linePost := ResetColor
	lines := RenderedSentence{}
	line := linePre

	currentIdx := 0
	currentWidth := 0

	for _, t := range self {
		// NOTE: I might not even use tokenIdx

		if len(t.String) > maxWidth {
			// TODO: I might want to eventually split the token to force it to wrap
			log.Fatalf("Length of token %v is greater than the allowed max width %v", t, maxWidth)
		}

		startIdx := currentIdx
		endIdx := currentIdx + len(t.String)

		if currentWidth+len(t.String) > maxWidth {
			// This line is done, append it and start new line
			line += linePost
			lines = append(lines, line) // NOTE: Is this actually copied, or will they all be references to the same string?
			line = linePre
			currentWidth = 0
		}

		if cursorIdx < startIdx {
			//     token
			//     s----e
			// ^^^^

			// I only need to handle "Untyped" style
			// I might be able to have a flag like `afterCursor` that I set here,
			// and if it's true, I don't need to style it at all (until the end, where I reset the style)
			// another idea: if I put the "Untyped" style after the cursor,
			// I won't need to style it at all

			// I don't need to handle any styles, because the "Untyped" style was applied right after the cursor
			line += t.String

			currentIdx += len(t.String)
			currentWidth += len(t.String)
		} else if cursorIdx >= endIdx {
			// token
			// s----e
			//      ^^^^

			// I only need to handle "Typed" and "Error" styles
			for i, r := range []rune(t.String) {
				if errors[currentIdx+i] {
					// There was an error at the current rune
					line += ErrorColor
					line += string(r)
					line += TypedColor
				} else {
					line += string(r)
				}
			}

			currentIdx += len(t.String)
			currentWidth += len(t.String)
		} else {
			// token
			// s----e
			// ^^^^^

			// I need to handle "Typed", "Untyped", "Error", and "Cursor" styles
			log.Printf("Cursor within token: %v", t)
			log.Printf("Rune slice: %v", []rune(t.String))
			for i, r := range []rune(t.String) {
				log.Printf("i: %v, r: %c", i, r)
				if currentIdx+i == cursorIdx {
					// The current rune is at the cursor
					// Append the rune (Cursor)
					line += CursorColor
					line += string(r)
					// Append the rest of the token after the cursor (Untyped)
					line += UntypedColor
					line += string([]rune(t.String)[i+1:]) // FIXME: Make sure this can't index out-of-bounds

					// The rest of the token has been appended. No need to loop over it
					break
				} else if errors[currentIdx+i] {
					// There was an error at the current rune
					line += ErrorColor
					line += string(r)
					line += TypedColor
				} else {
					line += string(r)
				}
			}

			currentIdx += len(t.String)
			currentWidth += len(t.String)
		}
	}

	// Append the last line
	line += linePost
	lines = append(lines, line)

	return lines
}

type RenderedSentence []string

// FIXME: I don't think this works properly anymore (Because the `RenderedSentece` is a slice of strings now)
func (s RenderedSentence) AppendLetter(l Letter) RenderedSentence {
	switch l.T {
	case TypedLetter:
		return append(s, TypedColor+string(l.Letter)+ResetColor)
	case UntypedLetter:
		return append(s, UntypedColor+string(l.Letter)+ResetColor)
	case CursorLetter:
		return append(s, CursorColor+string(l.Letter)+ResetColor)
	case ErrorLetter:
		return append(s, ErrorColor+string(l.Letter)+ResetColor)
	}

	return s
}

// FIXME: I don't think this works properly anymore (Because the `RenderedSentece` is a slice of strings now)
func (s RenderedSentence) AppendString(str string) RenderedSentence {
	return append(s, str)
}

// FIXME: I don't think this works properly anymore (Because the `RenderedSentece` is a slice of strings now)
func (s RenderedSentence) AppendResetColor() RenderedSentence {
	return append(s, ResetColor)
}

type letterType int

// TODO: Make a specific `style` package
// IDEA: Make a "testing" style set, where the strings are just `TypedColor = "{Typed}", ResetColor = "{Reset}"` etc., for easier testing/debugging

const (
	TypedColor   string = "\x1b[38;5;245m" + "\x1b[49m"
	UntypedColor string = "\x1b[38;5;255m" + "\x1b[49m"
	CursorColor  string = "\x1b[38;5;232m" + "\x1b[48;5;231m"
	ErrorColor   string = "\x1b[38;5;160m" + "\x1b[49m"
	ResetColor   string = "\x1b[39m" + "\x1b[49m"
)

const (
	// A letter that was typed correctly
	TypedLetter letterType = iota
	// A letter that hasn't been typed
	UntypedLetter
	// A letter that is under the cursor
	CursorLetter
	// A letter that was typed incorrectly
	ErrorLetter
)

type Letter struct {
	T      letterType
	Letter rune
}

type tokenType int

const (
	UnknownToken tokenType = iota
	WordToken
	PunctuationToken
	WhitespaceToken
)

type Token struct {
	T tokenType
	// NOTE: I am converting this to a []rune all the time. Should that just be its type?
	String string
}

func (self Token) Format(s fmt.State, verb rune) {
	var typeString string
	switch self.T {
	case UnknownToken:
		typeString = "Unknown"
	case WordToken:
		typeString = "Word"
	case PunctuationToken:
		typeString = "Punctuation"
	case WhitespaceToken:
		typeString = "Whitespace"
	default:
		typeString = "Error"
	}
	fmt.Fprintf(s, "{%s \"%s\"}", typeString, self.String)
}

func Split(runes []rune) Sentence {
	tokens := []Token{}
	currentTokenStart := 0
	// NOTE: This will probably fail because of the first iteration
	currentTokenType := UnknownToken

	for i, r := range runes {
		log.Printf("i: %v, r: '%c'", i, r)
		if unicode.IsSpace(r) {
			// The rune is whitespace
			log.Printf("Is whitespace")

			if currentTokenType != WhitespaceToken && i != 0 {
				tokens = append(tokens, Token{
					T:      currentTokenType,
					String: string(runes[currentTokenStart:i]),
				})
				currentTokenStart = i
			}
			currentTokenType = WhitespaceToken
		} else if unicode.In(r, unicode.Letter, unicode.Digit) {
			// The rune is a letter or digit
			log.Printf("Is letter or digit")

			if currentTokenType != WordToken && i != 0 {
				tokens = append(tokens, Token{
					T:      currentTokenType,
					String: string(runes[currentTokenStart:i]),
				})
				currentTokenStart = i
			}
			currentTokenType = WordToken
		} else if unicode.IsPunct(r) {
			// The rune is punctuation
			log.Printf("Is punctuation")

			if i != 0 {
				tokens = append(tokens, Token{
					T:      currentTokenType,
					String: string(runes[currentTokenStart:i]),
				})
				currentTokenStart = i
			}
			currentTokenType = PunctuationToken
		} else {
			// I don't know what the token is
			log.Printf("Is unknown")

			if currentTokenType != UnknownToken && i != 0 {
				tokens = append(tokens, Token{
					T:      currentTokenType,
					String: string(runes[currentTokenStart:i]),
				})
				currentTokenStart = i
			}
			currentTokenType = UnknownToken
		}
	}

	// Append the last token
	tokens = append(tokens, Token{
		T:      currentTokenType,
		String: string(runes[currentTokenStart:]),
	})

	return tokens
}
