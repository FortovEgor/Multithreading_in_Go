package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool, 3)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println(<-ch) // вот тут блокировка
	}()
	fmt.Println(1)
	ch <- true
	fmt.Println(2)
	ch <- true
	fmt.Println(3)
	ch <- true
	fmt.Println(4)
	ch <- true // вот тут блокировка наступает
	fmt.Println(5)
}

/*
ПРИМЕЧАНИЕ: по сути при full заполнении буфера и канала блокировка наступает всегда.
	Вопрос в том, есть ли у нас кто-то со стороны, кто снимет эту блокировку
	(т е to resolve deadlock нам нужен кто-то в отдельной горутине, кто начнет
		читать из канала и тем самым будет опустошать его)
*/
