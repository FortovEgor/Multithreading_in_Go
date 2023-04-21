package main

import "fmt"

func main() {
	intCh := make(chan int, 3) // инициализируем буферизированный канал с буфером в 3 int-овых элемента
	intCh <- 10
	intCh <- 3
	intCh <- 24
	// если добавить 4-ый элемент, возникнет deadlock
	// (по аналогии с небуферизированными каналами)
	// intCh <- 24 // UNCOMMENT THIS SH^T AND SEE WHAT WILL HAPPEN!
	fmt.Println(<-intCh) // 10
	fmt.Println(<-intCh) // 3
	fmt.Println(<-intCh) //24
}
