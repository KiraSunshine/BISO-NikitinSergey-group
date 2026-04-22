package domain

type TaskRepository interface {
	Add(task Task) (Task, error)
	Update(task Task) error
	Delete(id int) error
	GetByID(id int) (Task, bool)
	List() []Task
	NextID() int
}
