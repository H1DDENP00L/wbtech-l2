package patterns

/*
	5. Паттерн "Цепочка вызовов" (Chain of Responsibility)

	Назначение:
	Паттерн "Цепочка вызовов" позволяет избежать жесткой привязки отправителя запроса к
	его получателю, давая возможность нескольким объектам обработать запрос.
	Объекты-обработчики выстраиваются в цепочку, и запрос поочередно передается
	каждому обработчику до тех пор, пока какой-либо из них не сможет обработать запрос.

	Применимость:
	- Когда несколько объектов могут обрабатывать один и тот же запрос.
	- Когда нужно не привязывать конкретного получателя к запросу.
	- Когда нужно обрабатывать запросы в разном порядке в разное время.
	- Когда требуется, чтобы обработчик мог передавать запрос дальше по цепочке.
	- Когда нужно разделить логику обработки запроса на части

	Плюсы:
	- Снижение связанности: Отправитель запроса не знает, кто конкретно обработает запрос.
	- Гибкость: Легко добавлять, удалять и изменять порядок обработчиков в цепочке.
	- Расширяемость: Легко добавлять новые типы обработчиков.
	- Удобство обработки: Позволяет декомпозировать логику обработки запросов.

	Минусы:
	- Возможные трудности при отладке: Может быть сложно отслеживать, какой именно
	обработчик обработал запрос.
	- Гарантии обработки: Запрос может остаться необработанным, если ни один обработчик
	не смог его обработать.
	- Возможное снижение производительности: Длинная цепочка может привести к
	увеличению времени обработки запроса.

*/

import "fmt"

// Интерфейс обработчика
type SupportHandler interface {
	SetNext(handler SupportHandler)
	HandleRequest(request *SupportRequest)
}

// Конкретные типы запросов
type SupportRequest struct {
	RequestType string
	Description string
	IsResolved  bool
}

// Абстрактный класс обработчика
type AbstractSupportHandler struct {
	nextHandler SupportHandler
}

func (ah *AbstractSupportHandler) SetNext(handler SupportHandler) {
	ah.nextHandler = handler
}

func (ah *AbstractSupportHandler) HandleRequest(request *SupportRequest) {
	if ah.nextHandler != nil {
		ah.nextHandler.HandleRequest(request)
	}
}

// Конкретные обработчики
// Обработчик первого уровня
type FirstLevelSupport struct {
	AbstractSupportHandler
}

func NewFirstLevelSupport() *FirstLevelSupport {
	return &FirstLevelSupport{}
}

func (h *FirstLevelSupport) HandleRequest(request *SupportRequest) {
	if request.RequestType == "Basic" {
		fmt.Println("First level support handling the request : ", request.Description)
		request.IsResolved = true
		return
	}
	h.AbstractSupportHandler.HandleRequest(request)
}

// Обработчик второго уровня
type SecondLevelSupport struct {
	AbstractSupportHandler
}

func NewSecondLevelSupport() *SecondLevelSupport {
	return &SecondLevelSupport{}
}
func (h *SecondLevelSupport) HandleRequest(request *SupportRequest) {
	if request.RequestType == "Advanced" {
		fmt.Println("Second level support handling the request : ", request.Description)
		request.IsResolved = true
		return
	}
	h.AbstractSupportHandler.HandleRequest(request)
}

// Обработчик третьего уровня (неизвестный)
type EscalationSupport struct {
	AbstractSupportHandler
}

func NewEscalationSupport() *EscalationSupport {
	return &EscalationSupport{}
}

func (h *EscalationSupport) HandleRequest(request *SupportRequest) {
	fmt.Println("Escalation team, handling the request :", request.Description)
	request.IsResolved = true
	h.AbstractSupportHandler.HandleRequest(request)

}

// Клиентский код
func ProcessSupportRequest() {
	// создание цепочки обработки
	firstLevel := NewFirstLevelSupport()
	secondLevel := NewSecondLevelSupport()
	escalation := NewEscalationSupport()

	firstLevel.SetNext(secondLevel)
	secondLevel.SetNext(escalation)

	// запросы поддержки
	requests := []*SupportRequest{
		{RequestType: "Basic", Description: "Basic request 1"},
		{RequestType: "Advanced", Description: "Advanced request 1"},
		{RequestType: "Unknown", Description: "Unknown request 1"},
		{RequestType: "Basic", Description: "Basic request 2"},
		{RequestType: "Advanced", Description: "Advanced request 2"},
	}

	for _, req := range requests {
		firstLevel.HandleRequest(req)

		if !req.IsResolved {
			fmt.Printf("The request '%s' was not resolved.\n", req.Description)
		} else {
			fmt.Printf("The request '%s' was resolved.\n", req.Description)
		}
	}

}
