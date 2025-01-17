package patterns

/*
	8. Паттерн "Состояние" (State)

	Назначение:
	Паттерн "Состояние" позволяет объекту изменять свое поведение в зависимости от
	внутреннего состояния. При этом внешне класс объекта остается одним и тем же.
	Вместо того, чтобы иметь явный код, управляющий изменениями состояния,
	данный код переносится в отдельные классы, представляющие различные состояния
	объекта, которые реализуют общий интерфейс.

	Применимость:
	- Когда объект изменяет свое поведение в зависимости от внутреннего состояния.
	- Когда необходимо инкапсулировать состояние и логику переходов между состояниями.
	- Когда нужно избежать громоздких условных операторов, проверяющих текущее состояние.
	- Когда нужно добавить новые состояния в систему, не изменяя существующий код.

	Плюсы:
	- Явность: Логика переходов и поведения при каждом состоянии является явной
	и распределена по классам состояний.
	- Расширяемость: Легко добавлять новые состояния, не меняя код контекста и
	существующих состояний.
	- Разделение ответственности: Каждое состояние отвечает только за свое
	поведение, что делает код более модульным.
	- Соответствие Open/Closed Principle: Легко добавлять новые состояния, не изменяя
	существующий код.
	- Упрощение кода: Замена условных операторов отдельными классами состояний.

	Минусы:
	- Усложнение: Увеличение количества классов может сделать проект более сложным,
	особенно если состояний немного.
	- Связанность: Состояния должны быть осведомлены о контексте, в котором они находятся.
*/

import "fmt"

// Интерфейс состояния
type State interface {
	Play(player *MediaPlayer)
	Pause(player *MediaPlayer)
	Stop(player *MediaPlayer)
}

// Конкретные состояния

// Состояние воспроизведения
type PlayingState struct{}

func (s *PlayingState) Play(player *MediaPlayer) {
	fmt.Println("Already playing")
}

func (s *PlayingState) Pause(player *MediaPlayer) {
	fmt.Println("Pausing")
	player.setState(player.pausedState)
}

func (s *PlayingState) Stop(player *MediaPlayer) {
	fmt.Println("Stopping playback")
	player.setState(player.stoppedState)
}

// Состояние паузы
type PausedState struct{}

func (s *PausedState) Play(player *MediaPlayer) {
	fmt.Println("Resuming playback")
	player.setState(player.playingState)

}

func (*PausedState) Pause(*MediaPlayer) {
	fmt.Println("Already paused")
}

func (s *PausedState) Stop(player *MediaPlayer) {
	fmt.Println("Stopping from paused")
	player.setState(player.stoppedState)

}

// Состояние стоп
type StoppedState struct{}

func (*StoppedState) Play(player *MediaPlayer) {
	fmt.Println("Starting playback from stopped")
	player.setState(player.playingState)
}

func (*StoppedState) Pause(player *MediaPlayer) {
	fmt.Println("Cant pause from stopped")
}

func (*StoppedState) Stop(player *MediaPlayer) {
	fmt.Println("Already Stopped")
}

// Контекст (MediaPlayer)
type MediaPlayer struct {
	state        State
	playingState State
	pausedState  State
	stoppedState State
}

func NewMediaPlayer() *MediaPlayer {
	mp := &MediaPlayer{
		playingState: &PlayingState{},
		pausedState:  &PausedState{},
		stoppedState: &StoppedState{},
	}

	mp.state = mp.stoppedState
	return mp
}

func (mp *MediaPlayer) setState(state State) {
	mp.state = state
}
func (mp *MediaPlayer) Play() {
	mp.state.Play(mp)
}
func (mp *MediaPlayer) Pause() {
	mp.state.Pause(mp)

}
func (mp *MediaPlayer) Stop() {
	mp.state.Stop(mp)

}
func (mp *MediaPlayer) GetCurrentState() State {
	return mp.state
}

// Клиентский код

func MediaActions() {
	player := NewMediaPlayer()
	fmt.Println("Current State: Stopped")

	player.Play()
	fmt.Println("Current State: Playing")

	player.Pause()
	fmt.Println("Current State: Paused")
	player.Pause()

	player.Play()
	fmt.Println("Current State: Playing")

	player.Stop()
	fmt.Println("Current State: Stopped")

	player.Pause()

	player.Play()
	fmt.Println("Current State: Playing")

	player.Stop()
	fmt.Println("Current State: Stopped")

}
