package patterns

import "fmt"

/*
	4. Паттерн "Команда" (Command)

	Назначение:
	Паттерн "Команда" инкапсулирует запрос как объект, позволяя параметризировать клиентов
	с различными запросами, ставить запросы в очередь, вести их журнал, а также поддерживать
	отмену операций.

	Применимость:
	- Когда требуется параметризовать объекты действиями (например, кнопками в GUI).
	- Когда нужно поддерживать отмену операций.
	- Когда требуется ставить запросы в очередь.
	- Когда необходимо логировать запросы.
	- Когда требуется выполнять серии действий, или макросы.

	Плюсы:
	- Разделение ответственности: Разделяет отправителя запроса и получателя запроса.
	- Расширяемость: Легко добавлять новые команды.
	- Параметризация: Позволяет параметризировать запросы, включая и операции, и получателей.
	- Поддержка отмены операций:  Легко реализовать undo и redo.
	- Журналирование: Запросы могут быть записаны в журнал.

	Минусы:
	- Усложнение кода: Введение дополнительных классов может усложнить простой код.
	- Увеличение количества классов: Для каждой операции необходима отдельная команда.

*/

// Интерфейс для команды
type Command interface {
	Execute()
	Undo()
	Log() string
}

// Получатель
type TaskManager struct {
	tasks []string
	log   []string
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make([]string, 0),
		log:   make([]string, 0),
	}
}

func (tm *TaskManager) AddTask(task string) {
	tm.tasks = append(tm.tasks, task)
	tm.log = append(tm.log, fmt.Sprintf("Added task: %s", task))
	fmt.Printf("Task added: %s\n", task)
}

func (tm *TaskManager) RemoveTask(task string) {
	for i, t := range tm.tasks {
		if t == task {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			tm.log = append(tm.log, fmt.Sprintf("Removed task: %s", task))
			fmt.Printf("Task removed: %s\n", task)
			return
		}
	}
	fmt.Printf("Task '%s' not found for removal.\n", task)
}

func (tm *TaskManager) GetTasks() []string {
	return tm.tasks
}

// Конкретная команда - добавление задачи
type AddTaskCommand struct {
	task string
	tm   *TaskManager
}

func NewAddTaskCommand(task string, tm *TaskManager) *AddTaskCommand {
	return &AddTaskCommand{task: task, tm: tm}
}

func (c *AddTaskCommand) Execute() {
	c.tm.AddTask(c.task)
}

func (c *AddTaskCommand) Undo() {
	c.tm.RemoveTask(c.task)
}
func (c *AddTaskCommand) Log() string {
	return fmt.Sprintf("Add Task Command: %s", c.task)
}

// Конкретная команда - удаление задачи
type RemoveTaskCommand struct {
	task string
	tm   *TaskManager
}

func NewRemoveTaskCommand(task string, tm *TaskManager) *RemoveTaskCommand {
	return &RemoveTaskCommand{task: task, tm: tm}
}

func (c *RemoveTaskCommand) Execute() {
	c.tm.RemoveTask(c.task)
}
func (c *RemoveTaskCommand) Undo() {
	c.tm.AddTask(c.task)
}

func (c *RemoveTaskCommand) Log() string {
	return fmt.Sprintf("Remove Task Command: %s", c.task)
}

// Вызывающий (Invoker)
type CommandExecutor struct {
	commands []Command
	history  []Command
}

func NewCommandExecutor() *CommandExecutor {
	return &CommandExecutor{
		commands: make([]Command, 0),
		history:  make([]Command, 0),
	}
}

func (ce *CommandExecutor) AddCommand(command Command) {
	ce.commands = append(ce.commands, command)
}

func (ce *CommandExecutor) ExecuteCommands() {
	for _, command := range ce.commands {
		command.Execute()
		ce.history = append(ce.history, command) // Add to history only if executed
	}
	ce.commands = make([]Command, 0)

}

func (ce *CommandExecutor) UndoLastCommand() {
	if len(ce.history) > 0 {
		lastCommand := ce.history[len(ce.history)-1]
		lastCommand.Undo()
		ce.history = ce.history[:len(ce.history)-1]
		fmt.Printf("Undo command : %s \n", lastCommand.Log())
	} else {
		fmt.Println("No commands to undo.")
	}

}

func (ce *CommandExecutor) ShowCommandsHistory() {
	fmt.Println("\nCommands History:")
	for _, command := range ce.history {
		fmt.Printf(" - %s\n", command.Log())
	}

}

// Пример использования
func ProcessTasks() {
	taskManager := NewTaskManager()
	commandExecutor := NewCommandExecutor()

	addTask1 := NewAddTaskCommand("Implement task command pattern", taskManager)
	addTask2 := NewAddTaskCommand("Refactor the main class", taskManager)

	removeTask1 := NewRemoveTaskCommand("Refactor the main class", taskManager)

	commandExecutor.AddCommand(addTask1)
	commandExecutor.AddCommand(addTask2)

	commandExecutor.ExecuteCommands()

	fmt.Println("Current tasks:", taskManager.GetTasks())

	commandExecutor.AddCommand(removeTask1)
	commandExecutor.ExecuteCommands()

	commandExecutor.ShowCommandsHistory()

	commandExecutor.UndoLastCommand()

	fmt.Println("Current tasks after undo:", taskManager.GetTasks())
	commandExecutor.ShowCommandsHistory()

	commandExecutor.UndoLastCommand()
	fmt.Println("Current tasks after second undo:", taskManager.GetTasks())
	commandExecutor.ShowCommandsHistory()

	commandExecutor.UndoLastCommand()

}
