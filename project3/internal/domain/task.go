package domain

import "fmt"

type Task struct {
	ID          int
	Description string
	Done        bool
}

func (t Task) String() string {
	status := " "
	if t.Done {
		status = "x"
	}
	return fmt.Sprintf("%d. [%s] %s", t.ID, status, t.Description)
}
