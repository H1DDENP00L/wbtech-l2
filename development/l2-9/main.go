package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
	Задание
	Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
	- cd (args) - смена директории (в качестве аргумента могут быть то-то и то)
	- pwd - показать путь до текущего каталога
	- echo (args) - вывод аргумента в STDOUT
	- kill (args) - "убить" процесс, переданный в качестве аргумента (пример: такой-то пример)
	- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

	Так же требуется поддерживать функционал fork/exec-команд
	Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

	Шелл - это обычная консольная программа, которая будучи запущенной,
	в интерактивном сеансе выводит некое приглашение в STDOUT и ожидает ввода пользователя через STDIN.
	Дождавшись ввода, обрабатывает команду согласно своей логике и при необходимости выводит результат на экран.
	Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например quit)
*/

// Основная точка входа программы
func main() {
	fmt.Println("Type 'quit' to exit.")

	for {
		fmt.Print("GoShell> ")

		// Читаем пользовательский ввод
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		userInput := scanner.Text()

		// Завершение работы при вводе команды 'quit'
		if userInput == "quit" {
			break
		}

		// Обработка конвейера команд через пайпы
		commands := strings.Split(userInput, " | ")

		for _, cmd := range commands {
			// Разделение строки на команду и аргументы
			parts := strings.Fields(cmd)
			if len(parts) == 0 {
				continue
			}

			// Выполнение встроенных и внешних команд
			switch parts[0] {
			case "cd":
				executeCD(parts)
			case "pwd":
				executePWD()
			case "echo":
				executeEcho(parts)
			case "kill":
				executeKill(parts)
			case "ps":
				executePS()
			default:
				executeExternal(parts)
			}
		}
	}
}

// Меняет текущую директорию
func executeCD(args []string) {
	if len(args) < 2 {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "cd error: %v\n", err)
			return
		}
		if err := os.Chdir(homeDir); err != nil {
			fmt.Fprintf(os.Stderr, "cd error: %v\n", err)
		}
	} else {
		if err := os.Chdir(args[1]); err != nil {
			fmt.Fprintf(os.Stderr, "cd error: %v\n", err)
		}
	}
}

// Показывает текущую директорию
func executePWD() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "pwd error: %v\n", err)
		return
	}
	fmt.Println(currentDir)
}

// Выводит текст на экран
func executeEcho(args []string) {
	output := strings.Join(args[1:], " ")
	fmt.Println(output)
}

// Убивает процесс по PID
func executeKill(args []string) {
	if len(args) < 2 {
		fmt.Println("kill error: missing process ID")
		return
	}
	pid, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "kill error: invalid process ID\n")
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "kill error: %v\n", err)
		return
	}
	if err := process.Kill(); err != nil {
		fmt.Fprintf(os.Stderr, "kill error: %v\n", err)
	}
}

// Выводит список процессов
func executePS() {
	cmd := exec.Command("ps")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "ps error: %v\n", err)
	}
}

// Выполняет внешнюю команду
func executeExternal(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "command error: %v\n", err)
	}
}
