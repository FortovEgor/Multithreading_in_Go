package main

import (
	"fmt"
	"time"
)

// import "fmt"

func main() {

	intCh := make(chan int, 2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("ahah:", <-intCh)
		fmt.Println("bebeb:", <-intCh)
		fmt.Println("bebeb:", <-intCh)
		fmt.Println("bebeb:", <-intCh)
	}()
	fmt.Println("1")
	intCh <- 10
	intCh <- 11
	intCh <- 12
	// intCh <- 12  // TRY UNCOMMENT THIS LINE - nothing's gonna happen!
	fmt.Println("2")
	// intCh <- 24
	// fmt.Println(<-intCh) // 10
	// fmt.Println(<-intCh) // 3
	// fmt.Println(<-intCh) //24
	time.Sleep(time.Second)
}
