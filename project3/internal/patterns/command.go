package patterns

type Command interface {
	Execute() string
	Undo() string
}
