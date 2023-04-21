package main

import "fmt"

func main() {
	intCh := make(chan int, 3)
	ch := make(chan bool)
	intCh <- 10

	fmt.Println(cap(intCh)) // 3 - емкость БУФЕРА
	fmt.Println(len(intCh)) // 1 - кол-во элементов, положенных в БУФЕР И КАНАЛ
	fmt.Println(<-intCh)

	fmt.Println(cap(ch)) // 0 - ЕМКОСТЬ БУФЕРА (always 0, otherwise it's buffered channel)
	fmt.Println(len(ch)) // 0 - кол-во элементов, положенных в КАНАЛ (always 0, otherwise deadlock)
}
