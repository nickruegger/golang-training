package main

import (
    "fmt"
    "time"
)

func main() {
    
    sayHi := func(n string) {
    	fmt.Printf("Hi %s\n", n)
	}

    for _, name := range []string{"Andrew", "Beavis", "Cory"} {
        go sayHi(name)
    }
    time.Sleep(1 * time.Second)
}
