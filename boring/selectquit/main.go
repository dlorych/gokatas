// Use quit channel to stop the conversation. Also wait for them to tell us
// they're done talking.
//
// Level: intermediate
// Topics: goroutines, channels, select
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quit := make(chan string)
	c := boring("blah", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- "bye"
	fmt.Printf("They said: %q.\n", <-quit)
}

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf("%s, %d", msg, i):
			case <-quit:
				cleanup()
				quit <- "see you"
				return
			}
			n := rand.Intn(1e3)
			time.Sleep(time.Millisecond * time.Duration(n))
		}
	}()
	return c
}

func cleanup() {}
