package main

import "fmt"

func main() {
  var (
  	i int
  	f float64
  	b bool
  	s string 
  )
  fmt.Println(i)
  fmt.Println(f)
  fmt.Println(b)
  fmt.Println(s)
  const(
  	zero = iota*2
  	one
  	_
  	two
  )
  fmt.Println(zero)
  fmt.Println(one)
  fmt.Println(two)
}
