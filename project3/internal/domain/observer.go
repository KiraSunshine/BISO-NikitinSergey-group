package domain

type Observer interface {
	Update(event string, data any)
}
