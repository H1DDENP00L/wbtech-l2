package patterns

/*
	7. Паттерн "Стратегия" (Strategy)

	Назначение:
	Паттерн "Стратегия" определяет семейство алгоритмов,
	инкапсулирует каждый из них и делает их взаимозаменяемыми.
	Паттерн "Стратегия" позволяет выбирать алгоритм во время выполнения,
	не изменяя код клиента, использующий алгоритм.

	Применимость:
	- Когда нужно выбирать алгоритм во время выполнения.
	- Когда класс имеет много вариантов поведения (алгоритмов) которые нужно использовать по разным условиям.
	- Когда нужно инкапсулировать алгоритмы, отделяя их от классов, которые их используют.
	- Когда нужно избегать длинных условных операторов (if/else или switch)
		выбирающих тип поведения.

	Плюсы:
	- Гибкость:  Позволяет выбирать алгоритм во время выполнения.
	- Разделение ответственности: Разделяет алгоритмы от классов, использующих их.
	- Расширяемость: Позволяет легко добавлять новые алгоритмы.
	- Соответствие Open/Closed Principle: Легко добавлять новые стратегии, не меняя
	существующий код.
	- Избежание дублирования: Позволяет избежать дублирования логики в разных классах.

	Минусы:
	- Усложнение: Введение дополнительных интерфейсов и классов для стратегий может
	усложнить код в некоторых простых ситуациях.
	- Дополнительные классы: Применение этого паттерна ведет к увеличению числа классов.

*/

import (
	"fmt"
	"sort"
)

// Интерфейс стратегии
type SortStrategy interface {
	Sort(data []int)
	GetName() string
}

// Конкретные стратегии сортировки
// Сортировка пузырьком
type BubbleSort struct{}

func (bs *BubbleSort) Sort(data []int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

func (*BubbleSort) GetName() string {
	return "BubbleSort"
}

// Сортировка выбором
type SelectionSort struct{}

func (ss *SelectionSort) Sort(data []int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if data[j] < data[minIndex] {
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]

	}
}
func (*SelectionSort) GetName() string {
	return "SelectionSort"
}

// Сортировка встроеной в go
type GoSort struct{}

func (gs *GoSort) Sort(data []int) {
	sort.Ints(data)
}
func (*GoSort) GetName() string {
	return "GoSort"
}

// Контекст
type Sorter struct {
	strategy SortStrategy
}

func NewSorter(strategy SortStrategy) *Sorter {
	return &Sorter{strategy: strategy}
}

func (s *Sorter) SortData(data []int) {
	fmt.Println("Sorting using strategy : ", s.strategy.GetName())
	s.strategy.Sort(data)

}
func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

// Клиентский код
func SortExample() {

	data := []int{5, 2, 8, 1, 9, 4}

	// создание контекста и установка стратегий
	sorter := NewSorter(&BubbleSort{}) // Используем сортировку пузырьком по дефолту
	sorter.SortData(data)
	fmt.Println("Bubble Sort:", data)
	fmt.Println()

	sorter.SetStrategy(&SelectionSort{}) // Используем сортировку выбором
	data = []int{5, 2, 8, 1, 9, 4}
	sorter.SortData(data)
	fmt.Println("Selection sort: ", data)
	fmt.Println()

	sorter.SetStrategy(&GoSort{})
	data = []int{5, 2, 8, 1, 9, 4}
	sorter.SortData(data)
	fmt.Println("Go sort: ", data)
	fmt.Println()
}
