package domain

import (
	"errors"
	"testing"
)

type observerSpy struct{ events []string }

func (o *observerSpy) Update(event string, data any) { o.events = append(o.events, event) }

type repoMock struct {
	nextID    int
	tasks     map[int]Task
	addErr    error
	updateErr error
	deleteErr error
}

func newRepoMock() *repoMock { return &repoMock{nextID: 1, tasks: map[int]Task{}} }
func (r *repoMock) Add(task Task) (Task, error) {
	if r.addErr != nil {
		return Task{}, r.addErr
	}
	r.tasks[task.ID] = task
	return task, nil
}
func (r *repoMock) Update(task Task) error {
	if r.updateErr != nil {
		return r.updateErr
	}
	r.tasks[task.ID] = task
	return nil
}
func (r *repoMock) Delete(id int) error {
	if r.deleteErr != nil {
		return r.deleteErr
	}
	delete(r.tasks, id)
	return nil
}
func (r *repoMock) GetByID(id int) (Task, bool) { task, ok := r.tasks[id]; return task, ok }
func (r *repoMock) List() []Task {
	result := make([]Task, 0, len(r.tasks))
	for i := 1; i <= len(r.tasks)+1; i++ {
		if task, ok := r.tasks[i]; ok {
			result = append(result, task)
		}
	}
	return result
}
func (r *repoMock) NextID() int { id := r.nextID; r.nextID++; return id }

func TestTaskManagerAddTask(t *testing.T) {
	repo := newRepoMock()
	manager := NewTaskManager(repo)
	obs := &observerSpy{}
	manager.Subscribe(obs)

	id, err := manager.AddTask("Купить молоко")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 1 {
		t.Fatalf("expected ID 1, got %d", id)
	}
	if len(obs.events) != 1 || obs.events[0] != "task_added" {
		t.Fatalf("expected task_added notification, got %#v", obs.events)
	}
}

func TestTaskManagerAddTaskValidation(t *testing.T) {
	repo := newRepoMock()
	manager := NewTaskManager(repo)
	if _, err := manager.AddTask("   "); err == nil {
		t.Fatal("expected validation error for empty description")
	}
}

func TestTaskManagerMarkDoneTable(t *testing.T) {
	tests := []struct {
		name    string
		seed    []Task
		id      int
		wantErr bool
	}{
		{name: "existing task", seed: []Task{{ID: 1, Description: "test", Done: false}}, id: 1, wantErr: false},
		{name: "missing task", seed: []Task{}, id: 2, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := newRepoMock()
			for _, task := range tt.seed {
				repo.tasks[task.ID] = task
			}
			manager := NewTaskManager(repo)
			err := manager.MarkDone(tt.id)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error mismatch: got %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				updated, _ := repo.GetByID(tt.id)
				if !updated.Done {
					t.Fatal("task should be marked as done")
				}
			}
		})
	}
}

func TestTaskManagerDeleteTask(t *testing.T) {
	repo := newRepoMock()
	repo.tasks[1] = Task{ID: 1, Description: "Удалить меня"}
	manager := NewTaskManager(repo)

	if err := manager.DeleteTask(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := repo.GetByID(1); ok {
		t.Fatal("task was not deleted")
	}
}

func TestTaskManagerRepositoryErrors(t *testing.T) {
	repo := newRepoMock()
	repo.addErr = errors.New("add failed")
	manager := NewTaskManager(repo)
	if _, err := manager.AddTask("test"); err == nil {
		t.Fatal("expected repository add error")
	}
}
