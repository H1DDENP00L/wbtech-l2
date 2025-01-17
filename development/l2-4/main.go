package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseFile(f *os.File) ([]string, error) {

	var lines []string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}
	return lines, nil
}

func writeSortedFile(lines []string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return fmt.Errorf("ошибка записи в файл: %w", err)
		}
	}
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("ошибка сброса буфера: %w", err)
	}
	return nil
}

func extractColumns(line string, column int) string {
	fields := strings.Fields(line)
	if column <= 0 || column >= len(fields) {
		return ""
	}
	return fields[column-1]
}

func trimSpace(line string) string {
	return strings.TrimFunc(line, func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\n' || r == '\r'
	})
}

func main() {
	column := flag.Int("k", 0, "Номер колонки для сортировки")
	numeric := flag.Bool("n", false, "Числовая сортировка")
	reverse := flag.Bool("r", false, "Сортировка в обратном порядке")
	unique := flag.Bool("u", false, "Уникальная сортировка")
	ignoreSpace := flag.Bool("b", false, "Игнорировать хвостовые пробелы")

	output := flag.String("o", "sorted_output.txt", "Имя нового отсортированного файла")

	flag.Parse()

	// Получение пути к файлу из аргументов
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Не указан путь к файлу. Использование: ./programm -flag1 -flag2... path/to/file.txt")
	}

	filePath := args[0]

	// Открытие файла
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Не удается открыть файл: %v\n", err)
	}
	defer file.Close()

	lines, err := parseFile(file)
	if err != nil {
		log.Fatalf("Ошибка парсинга файла: %w\n", err)
	}

	if *ignoreSpace {
		for i := range lines {
			lines[i] = trimSpace(lines[i])
		}
	}
	sort.Slice(lines, func(i, j int) bool {
		line1 := lines[i]
		line2 := lines[j]

		col1 := extractColumns(line1, *column)
		col2 := extractColumns(line2, *column)

		if *numeric {
			val1, err1 := strconv.ParseFloat(col1, 64)
			val2, err2 := strconv.ParseFloat(col2, 64)

			if err1 != nil || err2 != nil {
				if *reverse {
					return col1 > col2
				}
				return col1 < col2
			}
			if *reverse {
				return val1 > val2
			}
			return val1 < val2
		}
		if *reverse {
			return col1 > col2
		}
		return col1 < col2
	})

	if *unique {
		seen := make(map[string]bool)
		var uniqueLines []string
		for _, line := range lines {
			if !seen[line] {
				seen[line] = true
				uniqueLines = append(uniqueLines, line)
			}
		}
		lines = uniqueLines
	}

	err = writeSortedFile(lines, *output)
	if err != nil {
		log.Fatalf("Ошибка записи в файл: %v", err)
	}
	fmt.Println("Файл успешно отсортирован и сохранен.")

}
