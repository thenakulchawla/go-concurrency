package channels

import (
	"fmt"
	"time"
)

func Unbuffured() {

	uc := make(chan int)
	go func() {
		fmt.Println("starting goroutine with unbuffered channel")
		uc <- 1
		uc <- 2
		fmt.Println("message sent to unbuffered channel")
	}()

	fmt.Println("sleep for 1 second")
	time.Sleep(1 * time.Second)
	fmt.Println("receiving from unbuffered channel")
	fmt.Println(<-uc)
	fmt.Println(<-uc)
}
