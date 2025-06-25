package state

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eamonburns/typtst/internal/types"
)

type endScreenModel struct {
	Errors map[int]types.TypingError
}

func (self endScreenModel) Init() tea.Cmd {
	return nil
}

func (self endScreenModel) Update(appModel AppModel, msg tea.Msg) (pageType, tea.Cmd) {
	log.Printf("> endScreenModel.Update(appModel: %v, msg: %v)", appModel, msg)
	switch msg.(type) {
	case tea.KeyMsg:
		return EndScreenPage, tea.Quit
	}
	return EndScreenPage, nil
}

func (self endScreenModel) View(appModel AppModel) string {
	log.Printf("> endScreenModel.View()")
	str := ""
	str += fmt.Sprintf("\n  Done, with %v errors\n", len(self.Errors))
	for i, e := range self.Errors {
		str += fmt.Sprintf("  i: %v, e: %v\n", i, e)
	}

	str += "\n\nPress any key to quit"
	return str
}
