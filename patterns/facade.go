package patterns

import "fmt"

/*
	1. Паттерн «Фасад» (Facade).
	Паттерн, позволяющий взаимодействовать со всеми подсистемами через "фасадный объект" и скрыть сложность их реализации.
	Фасадный объект - точка взаимодействия с элементами системы.

	В качестве примера был выбран вариант с умным домом. Используя функции умного дома (фасадного объекта), такие, как:
	"Смотреть фильм" или "Приготовиться к встрече гостей", пользователь использует функции сразу нескольких подсистем, не задумываясь
	об их реализации, и не тревожится использованием каждой из них по отдельности.


		Плюсы:
	- Упрощение: Предоставляет простой интерфейс для сложной подсистемы.
	- Уменьшение связанности: Клиент не зависит от конкретных деталей
	реализации подсистем.
	- Сокрытие сложности: Скрывает внутреннюю структуру и сложность подсистем.
	- Централизация: Централизует управление подсистемами в одном объекте.
	- Изменяемость: Легче изменять и развивать подсистемы без влияния на клиентов.
	- Улучшение читаемости кода: Фасад делает код более читаемым и легким в использовании.

	Минусы:
	- Ограничение: Фасад может не предоставлять доступа ко всем возможностям
	подсистем, предоставляя только наиболее востребованные функции.
	- Множественность: Создание множества фасадов для разных подсистем может привести к
	усложнению проекта.
	- Жесткость: Изменение API фасада может потребовать изменение всех клиентов.
	- Не оптимальность: При сложных операциях, может быть более целесообразным
  прямое использование внутренних компонентов, вместо работы через фасад.
*/

// MultimediaSystem - система для управления мультимедией в доме (звук, экран домашнего кинотеатра)
type MultimediaSystem struct {
}

// TurnOnCinemaScreen - функция для включения экрана домашнего кинотеатра
func (m *MultimediaSystem) TurnOnCinemaScreen() {
	fmt.Println("Экран домашнего кинотеатра включен")
}

// TurnOffCinemaScreen - функция для выключения экрана домашнего кинотеатра
func (m *MultimediaSystem) TurnOffCinemaScreen() {
	fmt.Println("Экран домашнего кинотеатра выключен")
}

// TurnOnAudio - функция для включения аудиосистемы домашнего кинотеатра
func (m *MultimediaSystem) TurnOnAudio() {
	fmt.Println("Звук включен")
}

// TurnOffAudio - функция для выключения аудиосистемы домашнего кинотеатра
func (m *MultimediaSystem) TurnOffAudio() {
	fmt.Println("Звук выключен")
}

// TemperatureSystem - система для управления температурой в доме
type TemperatureSystem struct {
	defaultTemperature int
}

// HeatRoomUp - функция включения обогрева
func (t *TemperatureSystem) HeatRoomUp() {
	fmt.Println("Обогрев включен")
	t.defaultTemperature += 3
}

// CoolRoom - функция для охлаждения комнаты
func (t *TemperatureSystem) CoolRoom() {
	fmt.Println("Комната охлаждается")
	t.defaultTemperature -= 3
}

type LightningSystem struct {
	defaultBrightness int
}

func (l *LightningSystem) LightBrighter() {
	fmt.Println("Яркость освещения повышена")
}

func (l *LightningSystem) LightDimmer() {
	fmt.Println("Яркость освещение понижена")
}

type NotificationSystem struct {
}

func (n *NotificationSystem) NotificationON() {
	fmt.Println("Оповещения от системы включены")
}
func (n *NotificationSystem) NotificationOFF() {
	fmt.Println("Оповещения от системы отключены")
}

type SmartHouse struct {
	multimediaSystem   *MultimediaSystem
	temperatureSystem  *TemperatureSystem
	lightningSystem    *LightningSystem
	notificationSystem *NotificationSystem
}

func NewSmartHouse(
	multimedia *MultimediaSystem,
	temperature *TemperatureSystem,
	lightning *LightningSystem,
	notification *NotificationSystem,
) *SmartHouse {
	return &SmartHouse{
		multimediaSystem:   multimedia,
		temperatureSystem:  temperature,
		lightningSystem:    lightning,
		notificationSystem: notification,
	}
}

func (sh *SmartHouse) WatchMovie() {
	sh.lightningSystem.LightDimmer()
	sh.multimediaSystem.TurnOnCinemaScreen()
	sh.multimediaSystem.TurnOnAudio()
	sh.notificationSystem.NotificationOFF()
	sh.temperatureSystem.CoolRoom()
	fmt.Println("Приятного просмотра...")
}

func (sh *SmartHouse) EndMovie() {
	sh.lightningSystem.LightBrighter()
	sh.multimediaSystem.TurnOffCinemaScreen()
	sh.multimediaSystem.TurnOffAudio()
	sh.temperatureSystem.HeatRoomUp()
	sh.notificationSystem.NotificationON()
	fmt.Println("Просмотр окончен...")
}

func (sh *SmartHouse) PrepareForGuestsVisit() {
	sh.notificationSystem.NotificationON()
	sh.temperatureSystem.HeatRoomUp()
	sh.multimediaSystem.TurnOnAudio()
	sh.lightningSystem.LightBrighter()
	fmt.Println("Вы готовы встречать гостей!")
}
