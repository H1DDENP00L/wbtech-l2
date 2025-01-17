package main

import (
	"fmt"
	"sort"
	"strings"
)

// Функция для сортировки букв в слове
func sortRunes(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

// Поиск множеств анаграмм
func findAnagrams(dictionary []string) map[string][]string {
	anagramGroups := make(map[string][]string)

	// Пройдем по всем словам из словаря
	for _, word := range dictionary {
		normalized := strings.ToLower(word) // Приведение к нижнему регистру
		sortedKey := sortRunes(normalized)  // Создаем ключ путем сортировки букв

		anagramGroups[sortedKey] = append(anagramGroups[sortedKey], normalized)
	}

	// Формирование результирующей мапы с учетом условий
	result := make(map[string][]string)
	for _, group := range anagramGroups {
		if len(group) > 1 {
			sort.Strings(group)      // Отсортируем группу анаграмм
			result[group[0]] = group // Используем первое слово в качестве ключа
		}
	}

	return result
}

func main() {
	// Пример входных данных
	words := []string{"кабан", "пятак", "тяпка", "банка", "столик", "листок", "пятка", "слиток", "молоток"}

	// Поиск анаграмм
	anagrams := findAnagrams(words)

	// Вывод результата
	for key, group := range anagrams {
		fmt.Printf("Ключ: %s, анаграммы: %v\n", key, group)
	}
}
