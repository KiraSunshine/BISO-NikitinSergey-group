package main

import (
	"flag"
	"fmt"
	"os"

	"gitlab.com/nikitins506/pr3_name_nsa_13/internal/domain"
	"gitlab.com/nikitins506/pr3_name_nsa_13/internal/patterns"
	"gitlab.com/nikitins506/pr3_name_nsa_13/internal/storage"
)

func main() {
	addFlag := flag.String("add", "", "добавить задачу")
	listFlag := flag.Bool("list", false, "показать список задач")
	doneFlag := flag.Int("done", 0, "отметить задачу выполненной по ID")
	deleteFlag := flag.Int("delete", 0, "удалить задачу по ID")
	demoFlag := flag.Bool("demo", false, "запустить демонстрационный сценарий")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Использование: %s [опции]\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Опции:")
		flag.PrintDefaults()
	}
	flag.Parse()

	manager := domain.NewTaskManager(storage.NewInMemoryRepository())
	manager.Subscribe(&patterns.ConsoleLogger{})
	factory := patterns.NewCommandFactory(manager)

	if *demoFlag {
		runDemo(factory)
		return
	}

	var cmd patterns.Command
	switch {
	case *addFlag != "":
		cmd = factory.CreateAdd(*addFlag)
	case *listFlag:
		seedDemoData(manager)
		cmd = factory.CreateList()
	case *doneFlag > 0:
		seedDemoData(manager)
		cmd = factory.CreateDone(*doneFlag)
	case *deleteFlag > 0:
		seedDemoData(manager)
		cmd = factory.CreateDelete(*deleteFlag)
	default:
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(cmd.Execute())
}

func seedDemoData(manager *domain.TaskManager) {
	if len(manager.GetTasks()) > 0 {
		return
	}
	_, _ = manager.AddTask("Подготовить проект")
	_, _ = manager.AddTask("Сделать UML")
	_, _ = manager.AddTask("Настроить CI")
}

func runDemo(factory *patterns.CommandFactory) {
	fmt.Println("Маркер проекта: name_nsa_13")
	fmt.Println(factory.CreateAdd("Подготовить пояснительную записку").Execute())
	fmt.Println(factory.CreateAdd("Проверить тесты").Execute())
	fmt.Println(factory.CreateDone(1).Execute())
	fmt.Println(factory.CreateList().Execute())
}
