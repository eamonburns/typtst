package sentence

type RenderedSentence string

func (s RenderedSentence) AppendLetter(l Letter) RenderedSentence {
	switch l.T {
	case LetterTyped:
		return s + RenderedSentence(TypedColor) + RenderedSentence([]rune{l.Letter}) + RenderedSentence(ResetColor)
	case LetterUntyped:
		return s + RenderedSentence(UntypedColor) + RenderedSentence([]rune{l.Letter}) + RenderedSentence(ResetColor)
	case LetterCursor:
		return s + RenderedSentence(CursorColor) + RenderedSentence([]rune{l.Letter}) + RenderedSentence(ResetColor)
	case LetterError:
		return s + RenderedSentence(ErrorColor) + RenderedSentence([]rune{l.Letter}) + RenderedSentence(ResetColor)
	}

	return s
}

func (s RenderedSentence) AppendString(str string) RenderedSentence {
	return s + RenderedSentence(str)
}

func (s RenderedSentence) AppendResetColor() RenderedSentence {
	return s + RenderedSentence(ResetColor)
}

type letterType int

const (
	TypedColor   string = "\x1b[38;5;245m"
	UntypedColor string = "\x1b[38;5;255m"
	CursorColor  string = "\x1b[38;5;232m\x1b[48;5;231m"
	ErrorColor   string = "\x1b[38;5;160m"
	ResetColor   string = "\x1b[39m\x1b[49m"
)

const (
	// A letter that was typed correctly
	LetterTyped letterType = iota
	// A letter that hasn't been typed
	LetterUntyped
	// A letter that is under the cursor
	LetterCursor
	// A letter that was typed incorrectly
	LetterError
)

type Letter struct {
	T      letterType
	Letter rune
}

type tokenType int

const (
	WordToken tokenType = iota
	PunctuationToken
	WhitespaceToken
)

type Token struct {
	t      tokenType
	string string
}

func Split(s string) []Token {
	return []Token{}
}
