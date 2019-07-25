package main

import "fmt"
import "math/rand"

func main() {
	// TODO: write a for loop that prints all the even numbers between 1 and 100
	
	

	fmt.Println("This is how you might implement a while loop")
	for {
		num := rand.Intn(100)
		fmt.Println(num)
		if num == 42 {
			fmt.Println("You found the magic number!")
			break
		}
	}
}

