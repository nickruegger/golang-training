package main

import "fmt"

func main() {
	fmt.Println("Doing a thing")
	defer func(){
		fmt.Println("First deferal")
	}()
	fmt.Println("Counting to 5")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 3 {
			defer func(){
				fmt.Println("Remind me that 3 is halfway through counting")
			}()
		}
	}
	fmt.Println("Done counting")
}
