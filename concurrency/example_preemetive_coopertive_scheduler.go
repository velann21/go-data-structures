package main

import (
	"fmt"
	"runtime"
	"time"
)

// This function will never get terminated because the golang does the cooperative schduling, Meaning the go schduler will not
// allocate any task until the goroutine go into idle mode, Here we are creating/ occupying all the OS threads and at same
// all the threads are busy in infinte loop, So main goroutine will never get any chance to exit the program,
// Same will works fine in java, C# because it uses the preemtive schduling by OS.
func main() {
	var x int
	threads := runtime.GOMAXPROCS(0)
	fmt.Println(threads)
	for i := 0; i < threads; i++ {
		go func() {
			for {
				x++
				//time.Sleep(time.Second*10)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("x =", x)
}

