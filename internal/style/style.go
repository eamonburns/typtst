package style

type styleType int

const (
	//
	ResetStyle styleType = iota
	// A letter that was typed correctly
	TypedStyle
	// A letter that hasn't been typed
	UntypedStyle
	// A letter that is under the cursor
	CursorStyle
	// A letter that was typed incorrectly
	ErrorStyle
)

type styleMap map[styleType]string

var currentStyle styleMap = map[styleType]string{
	ResetStyle:   "\x1b[39m" + "\x1b[49m",
	TypedStyle:   "\x1b[38;5;255m" + "\x1b[49m",
	UntypedStyle: "\x1b[38;5;245m" + "\x1b[49m",
	CursorStyle:  "\x1b[38;5;232m" + "\x1b[48;5;231m",
	ErrorStyle:   "\x1b[38;5;160m" + "\x1b[49m",
}

func Get(t styleType) string {
	return currentStyle[t]
}
