package state

import (
	"log"
	"strings"

	"github.com/agent-e11/typtst/internal/sentence"
	tea "github.com/charmbracelet/bubbletea"
)

type testScreenModel struct {
	sentence []rune
	cursor   int
	errors   map[int]bool
}

func newTestScreen(s string) testScreenModel {
	return testScreenModel{
		sentence: []rune(strings.Join(strings.Fields(s), " ")),
		cursor:   0,
		errors:   make(map[int]bool),
	}
}

func (self *testScreenModel) Update(appModel *AppModel, msg tea.Msg) (pageType, tea.Cmd) {
	log.Printf("> testScreenModel.Update(appModel: %v, msg: %v)", appModel, msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		log.Printf("Keypress msg: %v", msg)
		switch msg.String() {
		case "ctrl+c", "esc":
			return TestScreenPage, tea.Quit
		case "backspace":
			self.cursor = max(0, self.cursor-1)
			delete(self.errors, self.cursor) // Remove any error that may be at the new cursor position
			return TestScreenPage, nil
		default:
			if len(msg.Runes) != 1 {
				log.Printf("Not a single rune. `break` (len: %v, runes: %v)", len(msg.Runes), msg.Runes)
				break
			}

			key := msg.Runes[0]

			// TODO: If `space` is pressed, jump after next whitespace token
			// FIXME: In some cases, this can cause an out-of-bounds exception. I don't know how
			if key != []rune(self.sentence)[self.cursor] {
				log.Printf("key (%c) != current (%c)", key, []rune(self.sentence)[self.cursor])
				self.errors[self.cursor] = true
			}
			self.cursor += 1
		}
	}

	if self.cursor >= len(self.sentence) {
		log.Printf("Error set: %v, Len: %v", self.errors, len(self.errors))
		errorCount := len(self.errors)
		//fmt.Printf("Done, with %v errors", errorCount)
		appModel.EndScreenState.ErrorCount = errorCount
		return EndScreenPage, nil
	}

	return TestScreenPage, nil
}

func (self testScreenModel) View(appModel AppModel) string {
	log.Printf("> testScreenModel.View(appModel: %v)", appModel)

	// TODO: Make this configurable
	const padding = 10
	s := sentence.Split(self.sentence)

	if appModel.WindowWidth < 1 {
		// Window size hasn't been retrieved yet
		// Don't render anything
		return ""
	}
	horizontalPadding := appModel.WindowWidth / 4
	maxLineWidth := appModel.WindowWidth - (2 * horizontalPadding)

	lines := s.Render(self.cursor, self.errors, maxLineWidth)
	log.Printf("lines: %q", lines)

	// TODO: Vertically-center text using appModel.WindowHeight and len(lines)
	if len(lines) > appModel.WindowHeight {
		log.Fatalf("number of rendered lines (%v) is greater than height of window (%v)", len(lines), appModel.WindowHeight)
	}
	str := strings.Repeat("\n", (appModel.WindowHeight-len(lines))/2)
	for _, line := range lines {
		str += strings.Repeat(" ", horizontalPadding)
		str += line
		str += "\n"
	}
	return str
}
