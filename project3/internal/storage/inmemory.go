package storage

import (
	"sort"

	"gitlab.com/nikitins506/pr3_name_nsa_13/internal/domain"
)

type InMemoryRepository struct {
	tasks  map[int]domain.Task
	nextID int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{tasks: make(map[int]domain.Task), nextID: 1}
}
func (r *InMemoryRepository) Add(task domain.Task) (domain.Task, error) {
	r.tasks[task.ID] = task
	return task, nil
}
func (r *InMemoryRepository) Update(task domain.Task) error { r.tasks[task.ID] = task; return nil }
func (r *InMemoryRepository) Delete(id int) error           { delete(r.tasks, id); return nil }
func (r *InMemoryRepository) GetByID(id int) (domain.Task, bool) {
	task, ok := r.tasks[id]
	return task, ok
}
func (r *InMemoryRepository) NextID() int { id := r.nextID; r.nextID++; return id }
func (r *InMemoryRepository) List() []domain.Task {
	ids := make([]int, 0, len(r.tasks))
	for id := range r.tasks {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	result := make([]domain.Task, 0, len(ids))
	for _, id := range ids {
		result = append(result, r.tasks[id])
	}
	return result
}
