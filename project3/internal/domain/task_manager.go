package domain

import (
	"fmt"
	"strings"
)

type TaskManager struct {
	repo      TaskRepository
	observers []Observer
}

func NewTaskManager(repo TaskRepository) *TaskManager {
	return &TaskManager{repo: repo, observers: make([]Observer, 0)}
}

func (tm *TaskManager) Subscribe(o Observer) {
	tm.observers = append(tm.observers, o)
}

func (tm *TaskManager) Notify(event string, data any) {
	for _, observer := range tm.observers {
		observer.Update(event, data)
	}
}

func (tm *TaskManager) AddTask(description string) (int, error) {
	description = strings.TrimSpace(description)
	if description == "" {
		return 0, fmt.Errorf("описание задачи не может быть пустым")
	}

	task := Task{ID: tm.repo.NextID(), Description: description, Done: false}
	created, err := tm.repo.Add(task)
	if err != nil {
		return 0, err
	}

	tm.Notify("task_added", created)
	return created.ID, nil
}

func (tm *TaskManager) MarkDone(id int) error {
	task, ok := tm.repo.GetByID(id)
	if !ok {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}

	task.Done = true
	if err := tm.repo.Update(task); err != nil {
		return err
	}

	tm.Notify("task_done", task)
	return nil
}

func (tm *TaskManager) DeleteTask(id int) error {
	task, ok := tm.repo.GetByID(id)
	if !ok {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}

	if err := tm.repo.Delete(id); err != nil {
		return err
	}

	tm.Notify("task_deleted", task)
	return nil
}

func (tm *TaskManager) GetTasks() []Task {
	return tm.repo.List()
}
