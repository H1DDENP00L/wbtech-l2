package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

/*
	Задание:
	Создать программу, печатающую точное время с использованием NTP -библиотеки.
	Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp.
	Написать программу, печатающую текущее время / точное время с использованием этой библиотеки.

	Требования
	Программа должна быть оформлена как go module.

	Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS.
*/

// DisplayCurrentTime - функция выводящая на экран текущее точное время с обработкой ошибок
func DisplayCurrentTime() error {
	ntpTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Printf("Не удается получить время с сервера, %v\n", err)
		return err
	}
	fmt.Printf("Текущее точное время: %v\n", ntpTime)

	return nil
}

func main() {
	err := DisplayCurrentTime()
	if err != nil {
		log.Printf("Ошибка: %v\n", err)
		// Выход с ненулевым кодом
		os.Exit(1)
	}
}
