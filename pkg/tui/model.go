package tui

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/doniacld/charmify/habit-tracker-service/api"
	"github.com/doniacld/charmify/pkg/client"
	"github.com/doniacld/charmify/pkg/form"
	"github.com/doniacld/charmify/pkg/habit"
)

// Model holds necessary charm models and more to bring
// the habit tracker to life!
type Model struct {
	grpcClient api.HabitsClient
	habits     list.Model
	progress   progress.Model
	animation  Animation
	form       form.Model
	state      session
}

type session int

const (
	nav    session = 0
	create         = 1
)

// New creates a new Model.
func New(cli api.HabitsClient) *Model {
	// returns an initiated model
	m := &Model{
		grpcClient: cli,
		form:       form.InitialModel(),
		state:      nav,
	}

	// Adding a gradient motivates to get that extra tick,
	// sometimes maybe just to see the next colour.
	m.progress = progress.New(progress.WithDefaultGradient())

	ctx := context.Background()

	const (
		// These values will be overridden when the engine initially calls Update for the first time.
		// Check tea.WindowSizeMsg for more details.
		defaultWidth  = 80
		defaultHeight = 60
	)

	m.habits = list.New([]list.Item{}, list.NewDefaultDelegate(), defaultWidth, defaultHeight)
	m.habits.Title = "My list of habits"
	m.habits.SetStatusBarItemName("habit", "habits")
	m.habits.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			Keymap.Create,
			Keymap.Tick,
		}
	}

	// retrieve the habits from the server
	habits, err := client.ListHabits(ctx, m.grpcClient)
	if err != nil {
		log.Println(err)
	}

	// for demo purposes
	if len(habits) == 0 {
		var defaultHabits = []habit.Habit{habit.Code, habit.Read, habit.Walk}
		_, err = client.AddHabits(ctx, m.grpcClient, defaultHabits)
		if err != nil {
			log.Println(err)
		}
		habits = defaultHabits
	}

	m.habits.SetItems(habitsToItems(habits))

	return m
}

// resize set up the model with a list of default habits for any developer humankind
// and a corresponding progress bar animation.
func (m *Model) resize(width, height int) {
	m.habits.SetWidth(width / 2)
	m.habits.SetHeight(height)
	m.habits.ResetSelected()

	// create the progress bar animation
	m.animation.New()
}

func habitsToItems(habits []habit.Habit) []list.Item {
	items := make([]list.Item, 0)
	for _, h := range habits {
		items = append(items, h)
	}

	return items
}

// Init initializes Model behaviour. Here we start to tick to
// launch the progress bar animation
func (m Model) Init() tea.Cmd {
	return clockCmd()
}

// Update captures the new tea.Msg and runs actions depending on the new message.
// msg can be a window resize, a keystroke, or a tick.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)
	var cmd tea.Cmd

	// let's see what we got in this msg!
	switch msg := msg.(type) {
	case tea.WindowSizeMsg: // a simple resize, just relaunch the initialisation
		m.resize(msg.Width, msg.Height)

	case tea.KeyMsg:
		cmd = m.keyBindings(msg)
		cmds = append(cmds, cmd)

	case clockMsg:
		// A tick! Let's see first if the progress bar animation happened already...
		if m.animation.Complete() || m.habits.SelectedItem() == nil {
			break
		}

		// ... if not, generates the command.
		cmd = m.progressCmd()
		m.animation.Update(m.habits.Index(), true)
		return m, tea.Batch(clockCmd(), cmd)

	case progress.FrameMsg: // FrameMsg is sent when the progress bar wants to animate itself
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case form.CompleteMsg:
		cmd, err := m.addHabit()
		if err != nil {
			log.Println(err)
		}
		cmds = append(cmds, cmd)

		m.state = nav
		m.form = form.InitialModel()
	}

	// Return before refreshing the habits
	if m.state == create {
		return m, tea.Batch(cmds...)
	}

	cmd = m.refreshHabits(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) refreshHabits(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd

	// in any case we want to update the list of habits
	m.habits, cmd = m.habits.Update(msg)
	// we want to update the animation values only if we changed
	// the focus on another item of the list
	if m.animation.Selected() != m.habits.Index() {
		m.animation.Update(m.habits.Index(), false)
		return tea.Batch(clockCmd(), cmd)
	}

	return cmd
}

func (m *Model) addHabit() (tea.Cmd, error) {
	var cmd tea.Cmd

	name := m.form.Inputs[0].Value()
	desc := m.form.Inputs[1].Value()
	target, err := strconv.Atoi(m.form.Inputs[2].Value())
	if err != nil {
		return cmd, err
	}

	h, err := client.Add(context.Background(), m.grpcClient, name, desc, target)
	if err != nil {
		return cmd, err
	}

	cmd = m.habits.InsertItem(len(m.habits.Items()), h)

	return cmd, nil
}

// keyBindings returns a command depending on the keystroke.
func (m *Model) keyBindings(msg tea.KeyMsg) tea.Cmd {
	var cmd tea.Cmd

	if m.state == create {
		_, cmd = m.form.Update(msg)
		return cmd
	}

	switch {
	case key.Matches(msg, Keymap.Create):
		return m.creationStateCmd(msg)
	case key.Matches(msg, Keymap.Tick):
		return m.incrTicksCmd(msg)
	case key.Matches(msg, Keymap.Quit):
		return tea.Quit
	}

	return cmd
}

// View renders the view of the Model in a charming way.
func (m Model) View() string {
	var rightProgress, progressView, formView string

	// display the list of habits
	habitsView := m.habits.View()

	if m.habits.SelectedItem() == nil {
		return habitsView
	}

	if m.habits.SettingFilter() {
		return habitsView
	}

	// if we have selected an item, display its associated info and the progress bar animation
	current := m.habits.SelectedItem().(habit.Habit)
	rightProgress = fmt.Sprintf("Details: \nAchieved %d/%d", current.TicksCount, current.Target)
	progressView = m.progress.View()

	if m.state == create {
		formView = m.form.View()
	}

	// lipgloss library furnishes nice and useful methods to join views
	details := lipgloss.JoinVertical(lipgloss.Top, rightProgress, progressView, formView)
	return lipgloss.JoinHorizontal(lipgloss.Left, habitsView, details)
}
