package state

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/eamonburns/typtst/internal/sentence"
)

type pageType int

const (
	TestScreenPage pageType = iota
	EndScreenPage
)

type AppModel struct {
	TestScreenState testScreenModel
	EndScreenState  endScreenModel
	CurrentPage     pageType
	WindowHeight    int
	WindowWidth     int
}

func Init() AppModel {
	return AppModel{
		// Initialize the test screen with 50 random words
		//TestScreenState: newTestScreen(sentence.GenerateRandom(50)),
		TestScreenState: newTestScreen(sentence.GenerateRandom(10)),
		CurrentPage:     TestScreenPage,
	}
}

func (self AppModel) Init() tea.Cmd {
	log.Printf("> AppModel.Init()")
	return nil
}

func (self AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	log.Printf("> AppModel.Update(msg: %v)", msg)
	var nextPage pageType
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		self.WindowHeight = msg.Height
		self.WindowWidth = msg.Width
		log.Printf("(AppModel.Update) New window size: height: %v, width: %v", self.WindowHeight, self.WindowWidth)
	default:
		switch self.CurrentPage {
		case TestScreenPage:
			nextPage, cmd = self.TestScreenState.Update(&self, msg)
		case EndScreenPage:
			nextPage, cmd = self.EndScreenState.Update(self, msg)
		default:
			panic("Unkown state (AppModel.Update)")
		}
	}

	self.CurrentPage = nextPage

	// TODO: Should I limit the commands *State.Update methods can return?
	return self, cmd
}

func (self AppModel) View() string {
	log.Printf("> AppModel.View()")
	switch self.CurrentPage {
	case TestScreenPage:
		return self.TestScreenState.View(self)
	case EndScreenPage:
		return self.EndScreenState.View(self)
	default:
		panic("Unknown state (in View)")
	}
}
