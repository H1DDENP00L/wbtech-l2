package patterns

/*
	6. Паттерн "Фабричный метод" (Factory Method)

	Назначение:
	Паттерн "Фабричный метод" определяет интерфейс для создания объекта,
	но оставляет создание конкретных объектов подклассам.
	Этот паттерн позволяет отложить решение о том,
	какой класс создавать до времени выполнения.

	Применимость:
	- Когда класс не знает, какие подклассы ему нужно создавать.
	- Когда класс хочет делегировать создание объектов своим подклассам.
	- Когда создание объектов включает в себя сложную логику или конфигурации.
	- Когда нужно избегать прямой зависимости от конкретных классов.

	Плюсы:
	- Гибкость:  Позволяет добавлять новые типы продуктов, не изменяя существующий код.
	- Разделение ответственности: Отделяет процесс создания объектов от их использования.
	- Расширяемость: Позволяет легко добавлять новые типы продуктов путем создания новых
	подклассов фабрики.
	- Соответствие Open/Closed Principle: Легко добавлять новые фабрики и продукты, не
	изменяя существующий код.

	Минусы:
	- Усложнение: Вводит дополнительные уровни абстракции, что может усложнить
	некоторые проекты.
	- Увеличение количества классов: Создание интерфейсов и подклассов для фабрик и
	продуктов может увеличить количество классов в проекте.

*/

import "fmt"


// Интерфейс документа
type Document interface {
    Open()
    Close()
	GetType() string
}


// Конкретные документы
type PDFDocument struct {
    FileName string
}
func NewPDFDocument(name string) *PDFDocument {
	return &PDFDocument{FileName: name}
}


func (p *PDFDocument) Open(){
    fmt.Printf("Opening PDF document: %s \n", p.FileName)
}
func (p *PDFDocument) Close(){
    fmt.Printf("Closing PDF document: %s\n", p.FileName)
}

func (p *PDFDocument) GetType() string{
	return "PDF"
}


type WordDocument struct {
    FileName string
}

func NewWordDocument(name string) *WordDocument {
	return &WordDocument{FileName: name}
}

func (w *WordDocument) Open(){
	fmt.Printf("Opening Word document: %s \n", w.FileName)
}

func (w *WordDocument) Close(){
    fmt.Printf("Closing Word document: %s\n", w.FileName)
}
func (w *WordDocument) GetType() string {
    return "Word"
}


type TextDocument struct {
    FileName string
}

func NewTextDocument(name string) *TextDocument{
	return &TextDocument{FileName: name}
}

func (t *TextDocument) Open() {
    fmt.Printf("Opening Text document: %s\n", t.FileName)
}
func (t *TextDocument) Close() {
    fmt.Printf("Closing Text document: %s\n", t.FileName)
}

func (t *TextDocument) GetType() string {
	return "Text"
}


// Интерфейс фабрики документов
type DocumentFactory interface {
    CreateDocument(name string) Document
}

// Конкретные фабрики
// PDF
type PDFDocumentFactory struct {}
func (pdfFactory *PDFDocumentFactory) CreateDocument(name string) Document{
    return NewPDFDocument(name)
}

// Word
type WordDocumentFactory struct{}

func (wordFactory *WordDocumentFactory) CreateDocument(name string) Document{
	return NewWordDocument(name)
}

// Text
type TextDocumentFactory struct {}

func (textFactory *TextDocumentFactory) CreateDocument(name string) Document{
	return NewTextDocument(name)
}

// Клиентский код
func CreateAndProcessDocument(factory DocumentFactory, name string) {

	doc := factory.CreateDocument(name)
	fmt.Println("Document type: ", doc.GetType())
	doc.Open()
    doc.Close()
	fmt.Println()


}



func ProcessDocuments() {
    pdfFactory := &PDFDocumentFactory{}
    wordFactory := &WordDocumentFactory{}
	textFactory := &TextDocumentFactory{}

    CreateAndProcessDocument(pdfFactory,"example.pdf")
	CreateAndProcessDocument(wordFactory,"example.docx")
	CreateAndProcessDocument(textFactory,"example.txt")
}