package habit

type Habit struct {
	ID         string
	Name       string
	Desc       string
	Target     int
	TicksCount int
}

func (t Habit) FilterValue() string {
	return t.Name
}

func (t Habit) Title() string {
	return t.Name
}

func (t Habit) Description() string {
	return t.Desc
}

var (
	Code = Habit{
		Name:       "coding",
		Desc:       "in Go",
		Target:     5,
		TicksCount: 2,
	}
	Read = Habit{
		Name:       "reading",
		Desc:       "Jane Austen",
		Target:     5,
		TicksCount: 3,
	}
	Walk = Habit{
		Name:       "walk",
		Desc:       "in the forest",
		Target:     7,
		TicksCount: 3,
	}
	Clarinet = Habit{
		Name:       "play the clarinet",
		Desc:       "",
		Target:     7,
		TicksCount: 0,
	}
)
