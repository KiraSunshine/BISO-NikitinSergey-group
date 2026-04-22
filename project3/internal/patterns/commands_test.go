package patterns

import (
	"strings"
	"testing"

	"gitlab.com/nikitins506/pr3_name_nsa_13/internal/domain"
	"gitlab.com/nikitins506/pr3_name_nsa_13/internal/storage"
)

func setupManager() *domain.TaskManager {
	return domain.NewTaskManager(storage.NewInMemoryRepository())
}

func TestAddCommandExecute(t *testing.T) {
	manager := setupManager()
	cmd := NewAddCommand(manager, "Подготовить отчёт")
	got := cmd.Execute()
	if !strings.Contains(got, "Задача добавлена с ID 1") {
		t.Fatalf("unexpected result: %s", got)
	}
}

func TestListCommandExecute(t *testing.T) {
	manager := setupManager()
	_, _ = manager.AddTask("Одна задача")
	cmd := NewListCommand(manager)
	got := cmd.Execute()
	if !strings.Contains(got, "Список задач") || !strings.Contains(got, "Одна задача") {
		t.Fatalf("unexpected list output: %s", got)
	}
}

func TestDoneCommandExecute(t *testing.T) {
	manager := setupManager()
	_, _ = manager.AddTask("Закрыть задачу")
	cmd := NewDoneCommand(manager, 1)
	got := cmd.Execute()
	if !strings.Contains(got, "отмечена как выполненная") {
		t.Fatalf("unexpected result: %s", got)
	}
}

func TestCommandFactory(t *testing.T) {
	manager := setupManager()
	factory := NewCommandFactory(manager)
	cmds := []Command{factory.CreateAdd("test"), factory.CreateList(), factory.CreateDone(1), factory.CreateDelete(1)}
	for i, cmd := range cmds {
		if cmd == nil {
			t.Fatalf("factory returned nil command at index %d", i)
		}
	}
}
