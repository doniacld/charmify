package form

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

const (
	name   = iota
	desc   = 1
	target = 2
)

const (
	hotPink  = lipgloss.Color("#FF06B7")
	darkGray = lipgloss.Color("#767676")
)

var (
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type Model struct {
	Inputs  []textinput.Model
	Focused int
	Err     error
}

// nameValidator functions to ensure valid input
func nameValidator(s string) error {
	// TODO implement me
	return nil
}

func descValidator(s string) error {
	// TODO implement me
	return nil
}

func targetValidator(s string) error {
	// All we need to do is check that it is a number
	_, err := strconv.ParseInt(s, 10, 64)
	return err
}

func InitialModel() Model {
	var inputs = make([]textinput.Model, 3)
	inputs[name] = textinput.New()
	inputs[name].Placeholder = "Walk"
	inputs[name].Focus()
	inputs[name].CharLimit = 20
	inputs[name].Width = 30
	inputs[name].Prompt = ""
	inputs[name].Validate = nameValidator

	inputs[desc] = textinput.New()
	inputs[desc].Placeholder = "for 30 mins "
	inputs[desc].CharLimit = 20
	inputs[desc].Width = 30
	inputs[desc].Prompt = ""
	inputs[desc].Validate = descValidator

	inputs[target] = textinput.New()
	inputs[target].Placeholder = "14"
	inputs[target].CharLimit = 3
	inputs[target].Width = 5
	inputs[target].Prompt = ""
	inputs[target].Validate = targetValidator

	return Model{
		Inputs:  inputs,
		Focused: 0,
		Err:     nil,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.Inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.Focused == len(m.Inputs)-1 {
				return m, completeCmd()
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, completeCmd()
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.Inputs {
			m.Inputs[i].Blur()
		}
		m.Inputs[m.Focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.Err = msg
		return m, nil
	}

	for i := range m.Inputs {
		m.Inputs[i], cmds[i] = m.Inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return fmt.Sprintf(
		` Habit to start:

 %s
 %s

 %s  %s
 %s  %s

 %s
`,
		inputStyle.Width(30).Render("Name"),
		m.Inputs[name].View(),
		inputStyle.Width(30).Render("Description"),
		inputStyle.Width(6).Render("Target"),
		m.Inputs[desc].View(),
		m.Inputs[target].View(),
		continueStyle.Render("Continue ->"),
	) + "\n"
}

// nextInput focuses the next input field
func (m *Model) nextInput() {
	m.Focused = (m.Focused + 1) % len(m.Inputs)
}

// prevInput focuses the previous input field
func (m *Model) prevInput() {
	m.Focused--
	// Wrap around
	if m.Focused < 0 {
		m.Focused = len(m.Inputs) - 1
	}
}

// CompleteMsg is a message sent when form is filled.
type CompleteMsg struct{}

// completeCmd is a command to send a CompleteMsg.
func completeCmd() tea.Cmd {
	return func() tea.Msg {
		return CompleteMsg{}
	}
}
