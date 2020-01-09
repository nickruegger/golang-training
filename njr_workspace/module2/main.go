package main

import "fmt"

func main() {
  fmt.Println("hello", true, 1, 3.4, 2+3i)
  fmt.Printf("%T %T %T %T %T\n", "hello", true, 1, 3.4, 2+3i)
}
