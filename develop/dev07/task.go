package main

import (
	"fmt"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {
	// Определение
	var or func(channels ...<-chan interface{}) <-chan interface{}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		// Базовые случаи - если нет аргументов выдаем nil, если один - возвращаем его
		switch len(channels) {
		case 0:
			return nil
		case 1:
			return channels[0]
		}

		// Канал, по которому будем сообщать другим рекурсивно созданным горутинам о завершении работы
		orDone := make(chan interface{})

		go func() {
			defer close(orDone)

			switch len(channels) {
			// При n=2 прописываем case вручную
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			// В случае n>2, рекурсивно создаем новую горутину, куда передаем необработанные аргументы
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}

			}

		}()
		// По завершении передаем закрытый orDone
		return orDone
	}

	// Горутина для тестирования
	sig := func(after time.Duration) <- chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
	}()
	return c
	}
	
	start := time.Now()
	// Тест
	<-or (
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	
	fmt.Printf("fone after %v\n", time.Since(start))
}
