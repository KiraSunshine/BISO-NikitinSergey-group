package patterns

import "gitlab.com/nikitins506/pr3_name_nsa_13/internal/domain"

type CommandFactory struct{ manager *domain.TaskManager }

func NewCommandFactory(manager *domain.TaskManager) *CommandFactory {
	return &CommandFactory{manager: manager}
}
func (f *CommandFactory) CreateAdd(description string) Command {
	return NewAddCommand(f.manager, description)
}
func (f *CommandFactory) CreateList() Command         { return NewListCommand(f.manager) }
func (f *CommandFactory) CreateDone(id int) Command   { return NewDoneCommand(f.manager, id) }
func (f *CommandFactory) CreateDelete(id int) Command { return NewDeleteCommand(f.manager, id) }
