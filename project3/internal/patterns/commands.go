package patterns

import (
	"fmt"
	"strings"

	"gitlab.com/nikitins506/pr3_name_nsa_13/internal/domain"
)

type AddCommand struct {
	manager     *domain.TaskManager
	description string
}

func NewAddCommand(manager *domain.TaskManager, description string) *AddCommand {
	return &AddCommand{manager: manager, description: description}
}
func (c *AddCommand) Execute() string {
	id, err := c.manager.AddTask(c.description)
	if err != nil {
		return fmt.Sprintf("Ошибка: %v", err)
	}
	return fmt.Sprintf("Задача добавлена с ID %d", id)
}
func (c *AddCommand) Undo() string {
	tasks := c.manager.GetTasks()
	if len(tasks) == 0 {
		return "Нечего отменять"
	}
	last := tasks[len(tasks)-1]
	if err := c.manager.DeleteTask(last.ID); err != nil {
		return fmt.Sprintf("Ошибка отмены: %v", err)
	}
	return fmt.Sprintf("Отменено добавление задачи %d", last.ID)
}

type ListCommand struct{ manager *domain.TaskManager }

func NewListCommand(manager *domain.TaskManager) *ListCommand { return &ListCommand{manager: manager} }
func (c *ListCommand) Execute() string {
	tasks := c.manager.GetTasks()
	if len(tasks) == 0 {
		return "Список задач пуст"
	}
	lines := []string{"Список задач:"}
	for _, task := range tasks {
		lines = append(lines, task.String())
	}
	return strings.Join(lines, "\n")
}
func (c *ListCommand) Undo() string {
	return "Для команды просмотра отмена не требуется"
}

type DoneCommand struct {
	manager *domain.TaskManager
	id      int
}

func NewDoneCommand(manager *domain.TaskManager, id int) *DoneCommand {
	return &DoneCommand{manager: manager, id: id}
}
func (c *DoneCommand) Execute() string {
	if err := c.manager.MarkDone(c.id); err != nil {
		return fmt.Sprintf("Ошибка: %v", err)
	}
	return fmt.Sprintf("Задача %d отмечена как выполненная", c.id)
}
func (c *DoneCommand) Undo() string {
	return "Undo для done в демонстрационной версии не реализован"
}

type DeleteCommand struct {
	manager *domain.TaskManager
	id      int
}

func NewDeleteCommand(manager *domain.TaskManager, id int) *DeleteCommand {
	return &DeleteCommand{manager: manager, id: id}
}
func (c *DeleteCommand) Execute() string {
	if err := c.manager.DeleteTask(c.id); err != nil {
		return fmt.Sprintf("Ошибка: %v", err)
	}
	return fmt.Sprintf("Задача %d удалена", c.id)
}
func (c *DeleteCommand) Undo() string {
	return "Undo для delete в демонстрационной версии не реализован"
}
