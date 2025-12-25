package todo

type Todo struct {
	Title     string
	Completed bool
}

func NewTodo(title string) Todo {
	return Todo{
		Title:     title,
		Completed: false,
	}
}

func (t *Todo) MarkDone() {
	t.Completed = true
}

func (t *Todo) Toggle() {
	t.Completed = !t.Completed
}
