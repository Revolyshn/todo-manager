package task

type Task struct {
	Description string
	Completed   bool
}

func NewTask(description string) Task {
	return Task{
		Description: description,
		Completed:   false,
	}
}