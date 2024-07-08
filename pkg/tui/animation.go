package tui

type Animation struct {
	selected int
	complete bool
}

func (a *Animation) Update(s int, complete bool) {
	a.selected = s
	a.complete = complete
}

func (a *Animation) New() {
	a.selected = -1
	a.complete = false
}

func (a *Animation) Selected() int {
	return a.selected
}

func (a *Animation) Complete() bool {
	return a.complete
}
