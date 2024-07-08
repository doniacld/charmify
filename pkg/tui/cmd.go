package tui

import (
	"context"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/doniacld/charmify/pkg/client"
	"github.com/doniacld/charmify/pkg/habit"
)

func (m *Model) creationStateCmd(msg tea.KeyMsg) tea.Cmd {
	m.state = create
	return m.createCmd(msg)
}

func (m *Model) createCmd(_ tea.KeyMsg) tea.Cmd {
	cmds := make([]tea.Cmd, 0)
	var cmd tea.Cmd

	cmd = m.form.Init()
	cmds = append(cmds, cmd)

	_, cmd = m.form.Update("")
	cmds = append(cmds, cmd)

	cmd = m.refreshHabits("")
	cmds = append(cmds, cmd)

	return tea.Sequence(cmds...)
}

// incrTicksCmd increments the number of ticks for a habit
// and, relaunch the progress bar animation with the clockCmd.
func (m *Model) incrTicksCmd(msg tea.KeyMsg) tea.Cmd {
	h := m.habits.SelectedItem().(habit.Habit)

	newHabit, err := client.TickHabit(context.Background(), m.grpcClient, h.ID)
	if err != nil {
		log.Println("failed to tick habit")
	}

	m.habits.Items()[m.habits.Index()] = newHabit
	_, cmd := m.habits.Update(msg)
	m.animation.New()

	return tea.Batch(clockCmd(), cmd)
}

type clockMsg time.Time

// clockCmd returns a tick every second.
func clockCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*40, func(t time.Time) tea.Msg {
		return clockMsg(t)
	})
}

// progressCmd returns the cmd command for the progress bar animation.
// Its sets the percent to display to the number of ticks achieved on 100%.
// /!\ I do not know how to display more than 100% for the moment
func (m *Model) progressCmd() tea.Cmd {
	h := m.habits.SelectedItem().(habit.Habit)

	var percent float64
	if h.TicksCount == 0 {
		percent = 0
	} else {
		percent = float64(h.TicksCount) / float64(h.Target)
	}

	return m.progress.SetPercent(percent)
}
