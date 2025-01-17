package patterns

import "fmt"

/*
	2. Паттерн "Строитель" (Builder)

	Назначение:
	Паттерн "Строитель" отделяет конструирование сложного объекта
	от его представления, что позволяет создавать различные вариации
	этого объекта, используя один и тот же процесс конструирования.
	Строитель позволяет избежать "телескопических конструкторов"
	(конструкторов с большим количеством параметров).

	Применимость:
	- Когда процесс создания объекта включает множество шагов или параметров.
	- Когда необходимо создавать различные вариации объекта с помощью
	одного и того же процесса.
	- Когда требуется избежать "телескопических конструкторов".
	- Когда требуется гарантировать неизменность (immutable) создаваемых объектов.

	Плюсы:
	- Улучшение читаемости и сопровождаемости: Код становится более понятным,
	процесс конструирования разделен на отдельные шаги.
	- Гибкость: Можно создавать различные вариации объектов, меняя последовательность
	вызываемых методов строителя.
	- Контроль: Позволяет контролировать процесс создания объекта, избегая
	ошибок инициализации.
	- Immutable объекты: Строитель позволяет легко создавать неизменяемые
	объекты.
	- Улучшает чистоту кода: Код клиента, использующего паттерн становится проще.

	Минусы:
	- Усложнение кода: Требуется написание дополнительного класса строителя,
	что может показаться избыточным для простых объектов.
	- Увеличение количества классов: В проекте появляются новые классы,
	что может усложнить структуру проекта.
*/

// Структура машины (продукт)
type Car struct {
	Model    string
	Engine   string
	Color    string
	NumDoors int
	HasGPS   bool
}

// String метод для печати данных машины
func (c Car) String() string {
	return fmt.Sprintf("Model: %s, Engine: %s, Color: %s, Doors: %d, GPS: %t",
		c.Model, c.Engine, c.Color, c.NumDoors, c.HasGPS)
}

// Интерфейс Строителя
type CarBuilder interface {
	SetModel(model string) CarBuilder
	SetEngine(engine string) CarBuilder
	SetColor(color string) CarBuilder
	SetNumDoors(doors int) CarBuilder
	SetGPS(hasGPS bool) CarBuilder
	Build() Car
}

// Конкретный Строитель
type ConcreteCarBuilder struct {
	model    string
	engine   string
	color    string
	numDoors int
	hasGPS   bool
}

func NewCarBuilder() *ConcreteCarBuilder {
	return &ConcreteCarBuilder{}
}

func (b *ConcreteCarBuilder) SetModel(model string) CarBuilder {
	b.model = model
	return b
}

func (b *ConcreteCarBuilder) SetEngine(engine string) CarBuilder {
	b.engine = engine
	return b
}

func (b *ConcreteCarBuilder) SetColor(color string) CarBuilder {
	b.color = color
	return b
}

func (b *ConcreteCarBuilder) SetNumDoors(doors int) CarBuilder {
	b.numDoors = doors
	return b
}

func (b *ConcreteCarBuilder) SetGPS(hasGPS bool) CarBuilder {
	b.hasGPS = hasGPS
	return b
}

func (b *ConcreteCarBuilder) Build() Car {
	return Car{
		Model:    b.model,
		Engine:   b.engine,
		Color:    b.color,
		NumDoors: b.numDoors,
		HasGPS:   b.hasGPS,
	}
}
