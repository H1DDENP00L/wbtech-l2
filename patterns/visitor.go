package patterns

import "fmt"

/*
	3. Паттерн "Посетитель" (Visitor)

	Назначение:
	Паттерн "Посетитель" позволяет добавлять новые операции к иерархии классов
	без изменения самих классов. Это достигается путем отделения алгоритма от
	структуры объектов, к которым он применяется.

	Применимость:
	- Когда нужно выполнять различные операции над объектами разного типа, но
	с сохранением их интерфейсов.
	- Когда нужно добавлять новые операции к иерархии классов, не изменяя их.
	- Когда нужно избегать раздувания классов методами.
	- Когда нужно разделить сложные алгоритмы на более мелкие,
	посещающие конкретные объекты.

	Плюсы:
	- Расширяемость: Легко добавлять новые операции, добавляя новых посетителей,
	без изменения классов, которые они посещают.
	- Разделение ответственности: Алгоритмы и структуры данных разделены.
	- Open/Closed Principle: Можно добавлять новые функциональности, не изменяя
	существующий код.

	Минусы:
	- Усложнение: Внедрение паттерна может привести к увеличению количества классов
	и усложнению кода.
	- Сложности при изменении иерархии: Изменение иерархии элементов может
	потребовать изменения интерфейсов посетителей.
	- Связанность: Посетитель привязан к конкретной иерархии классов.
*/

// Интерфейс для элементов (расходов)
type Expense interface {
	Accept(visitor Visitor)
	GetType() string
}

// Конкретные элементы (типы расходов)
type FoodExpense struct {
	Amount float64
}

func (f *FoodExpense) Accept(v Visitor) {
	v.VisitFoodExpense(f)
}

func (f *FoodExpense) GetType() string {
	return "Food"
}

type TransportExpense struct {
	Amount   float64
	Distance float64
}

func (t *TransportExpense) Accept(v Visitor) {
	v.VisitTransportExpense(t)
}
func (t *TransportExpense) GetType() string {
	return "Transport"
}

type EntertainmentExpense struct {
	Amount float64
}

func (e *EntertainmentExpense) Accept(v Visitor) {
	v.VisitEntertainmentExpense(e)
}

func (e *EntertainmentExpense) GetType() string {
	return "Entertainment"
}

// Интерфейс посетителя
type Visitor interface {
	VisitFoodExpense(expense *FoodExpense)
	VisitTransportExpense(expense *TransportExpense)
	VisitEntertainmentExpense(expense *EntertainmentExpense)
}

// Конкретные посетители
// Посетитель для расчета общей суммы
type TotalAmountVisitor struct {
	Total float64
}

func (v *TotalAmountVisitor) VisitFoodExpense(expense *FoodExpense) {
	v.Total += expense.Amount
}

func (v *TotalAmountVisitor) VisitTransportExpense(expense *TransportExpense) {
	v.Total += expense.Amount
}

func (v *TotalAmountVisitor) VisitEntertainmentExpense(expense *EntertainmentExpense) {
	v.Total += expense.Amount
}

// Посетитель для вывода отчета
type ReportVisitor struct {
	Report string
}

func (v *ReportVisitor) VisitFoodExpense(expense *FoodExpense) {
	v.Report += fmt.Sprintf("- Food Expense: %.2f\n", expense.Amount)
}

func (v *ReportVisitor) VisitTransportExpense(expense *TransportExpense) {
	v.Report += fmt.Sprintf("- Transport Expense: %.2f (Distance: %.2f)\n", expense.Amount, expense.Distance)
}
func (v *ReportVisitor) VisitEntertainmentExpense(expense *EntertainmentExpense) {
	v.Report += fmt.Sprintf("- Entertainment Expense: %.2f\n", expense.Amount)
}

// Другой пример применения паттерна - сбор статистики
// Посетитель для сбора статистики
type StatisticsVisitor struct {
	Counts        map[string]int
	totalDistance float64
}

func NewStatisticsVisitor() *StatisticsVisitor {
	return &StatisticsVisitor{Counts: make(map[string]int)}
}

func (v *StatisticsVisitor) VisitFoodExpense(expense *FoodExpense) {
	v.Counts["Food"]++
}
func (v *StatisticsVisitor) VisitTransportExpense(expense *TransportExpense) {
	v.Counts["Transport"]++
	v.totalDistance += expense.Distance
}

func (v *StatisticsVisitor) VisitEntertainmentExpense(expense *EntertainmentExpense) {
	v.Counts["Entertainment"]++
}

// Пример использования паттерна
func ProcessExpenses() {
	expenses := []Expense{
		&FoodExpense{Amount: 15.50},
		&TransportExpense{Amount: 25.00, Distance: 12.5},
		&EntertainmentExpense{Amount: 30.00},
		&FoodExpense{Amount: 12.75},
		&TransportExpense{Amount: 40.00, Distance: 20},
	}

	// Использование посетителя для расчета общей суммы
	totalAmountVisitor := &TotalAmountVisitor{}
	for _, expense := range expenses {
		expense.Accept(totalAmountVisitor)
	}
	fmt.Printf("Total expenses: %.2f\n", totalAmountVisitor.Total)

	// Использование посетителя для создания отчета
	reportVisitor := &ReportVisitor{}
	for _, expense := range expenses {
		expense.Accept(reportVisitor)
	}

	fmt.Println("\nExpenses Report:")
	fmt.Print(reportVisitor.Report)

	// Использование посетителя для сбора статистики
	statsVisitor := NewStatisticsVisitor()
	for _, expense := range expenses {
		expense.Accept(statsVisitor)
	}

	fmt.Println("\nExpense Statistics:")
	fmt.Println("Counts:", statsVisitor.Counts)
	fmt.Printf("Total distance travelled: %.2f\n", statsVisitor.totalDistance)

}
