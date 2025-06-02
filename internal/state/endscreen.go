package state

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type endScreenModel struct {
	ErrorCount int
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
	return fmt.Sprintf("\n  Done, with %v errors\n\nPress any key to quit", self.ErrorCount)
}
