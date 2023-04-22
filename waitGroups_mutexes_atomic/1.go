package main

/////////////////////////////////////
/// СИНХРОНИЗАЦИЯ ЧЕРЕЗ WaitGroup ///
/////////////////////////////////////

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32 = 0

	wg := sync.WaitGroup{} // позволит всем горутинам выполнится и только после этого завершить программу
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			// counter++
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	// time.Sleep(time.Second)
	fmt.Println(counter)
}
