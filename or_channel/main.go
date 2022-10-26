package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Принимает на вход множество односторонних каналов(only read)
Возвращает один такой же канал
*/
func or(channels ...<-chan interface{}) <-chan interface{} {
	retChan := make(chan interface{}, len(channels)) // Сохдание канала с буфером размером с множество

	for _, channel := range channels {
		go func(channel <-chan interface{}) { // Выгружаем содержимое каждого канала в single канал
			retChan <- <-channel
		}(channel)
	}
	return retChan // При ретурне преодбразует двунаправленный канал в read-only channel
}

func main() {
	sig := func(tim time.Duration) <-chan interface{} { // функция создающая, заполняющая, и разгружающая канал
		retCahn := make(chan interface{}, 1) // создается буферизированный канал на одно значение
		go func() {                          // в отдельной горутине выполняется функция, которая ложит еденицу в retChan
			defer close(retCahn)
			time.Sleep(tim)
			retCahn <- rand.Intn(100)
		}()
		return retCahn
	}

	start := time.Now()
	ch := or(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
	)

	fmt.Println("1:", <-ch)
	fmt.Println("2:", <-ch)
	fmt.Println("3:", <-ch)
	//BIG:
	//	for {
	//		ok := false
	//		select {
	//		case channel, ok := <-ch:
	//			if !ok {
	//				break BIG
	//			}
	//			fmt.Println("1:", channel)
	//
	//		}
	//	}
	fmt.Println(time.Since(start))
}
