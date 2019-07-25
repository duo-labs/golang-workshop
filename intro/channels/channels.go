package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	c <- "Hello!"
	
	fmt.Println(<-c)
}

// TODO 1. Try to run it. Try to fix it with a goroutine!
// TODO 2. Remove the goroutine and add a buffer to the channel.
// TODO 3. Use send and receive channels...

