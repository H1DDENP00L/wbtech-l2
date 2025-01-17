package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализовать функцию, которая будет объединять один или более done-каналов в single-канал,
если один из его составляющих каналов закроется.

Очевидным вариантом решения могло бы стать выражение при использовании select,
которое бы реализовывало эту связь, однако иногда неизвестно общее число done-каналов,
с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов единственной функции,
которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.
*/
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		// Обработка отсутствия каналов
		return nil
	case 1:
		// Если канал один возвращаем его
		return channels[0]
	}
	// Создаем новый канал, который будет использоваться для объединения входных каналов
	out := make(chan interface{})
	var once sync.Once

	launch := func(ch <-chan interface{}) {
		go func() {
			<-ch
			once.Do(func() {
				close(out)
			})
		}()
	}

	for _, ch := range channels {
		launch(ch)
	}

	// Возвращаем объединенный канал
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()

		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(3*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("Done after %v", time.Since(start))
}
