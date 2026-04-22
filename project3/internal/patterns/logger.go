package patterns

import "fmt"

type ConsoleLogger struct{}

func (l *ConsoleLogger) Update(event string, data any) { fmt.Printf("[LOG] %s: %v\n", event, data) }
