package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	// Разбираем аргументы командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "таймаут на подключение")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Использование: go-telnet [--timeout=10s] <host> <port>")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	// Подключаемся к серверу
	conn, err := net.DialTimeout("tcp", address, *timeout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка подключения: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Подключено к %s\n", address)

	// Создаем каналы для обработки данных и завершения
	done := make(chan struct{})

	// Чтение данных из сокета и вывод в STDOUT
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			if err != io.EOF { // Проверяем, что ошибка не EOF, так как EOF означает нормальное закрытие сокета
				fmt.Fprintf(os.Stderr, "Ошибка чтения из сокета: %v\n", err)
			}
		}
		close(done)
	}()

	// Чтение данных из STDIN и отправка их в сокет
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if !scanner.Scan() {
				if err := scanner.Err(); err != nil {
					if err != io.EOF {
						fmt.Fprintf(os.Stderr, "Ошибка чтения из STDIN: %v\n", err)
					}
				}
				break // Выходим из цикла, если достигнут EOF или есть ошибка
			}
			_, err := fmt.Fprintln(conn, scanner.Text())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ошибка записи в сокет: %v\n", err)
				break
			}
		}
		// Закрытие соединения при нажатии Ctrl+D
		conn.Close()
		close(done)
	}()

	// Ждем завершения одного из потоков
	<-done
	fmt.Println("Соединение закрыто")
}
