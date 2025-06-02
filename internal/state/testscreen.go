package state

import (
	"log"

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
		sentence: []rune(s),
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

			// FIXME: In some cases, this can cause an out-of-bounds exception
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
	log.Printf("> testScreenModel.View()")
	s := sentence.RenderedSentence("\n\n  ")
	for i, r := range []rune(self.sentence) {
		l := sentence.Letter{
			Letter: r,
		}
		if i == self.cursor {
			l.T = sentence.LetterCursor
		} else if i > self.cursor {
			l.T = sentence.LetterUntyped
		} else if self.errors[i] {
			l.T = sentence.LetterError
		}
		s = s.AppendLetter(l)
	}
	s = s.AppendResetColor()
	s = s.AppendString("\n\n")
	return string(s)
}
