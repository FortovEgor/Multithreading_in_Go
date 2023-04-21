package main

import (
	"fmt"
)

// Каналы (channels) - это инструменты коммуникации между горутинами. //

func main() {
	ch := make(chan int) // проинициализировали небуферизированный канал

	fmt.Println(1)

	fmt.Println("Значение из канала:", <-ch) // блокировка: пытаемся прочитать из пустого канала!

	fmt.Println(2)

	ch <- 5

	fmt.Println(3)
}
