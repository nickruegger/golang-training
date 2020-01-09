package main

import "fmt"

func main() {
	s := make([][4]map[string]int, 4)
	m := map[string]int{
		"value1": 1,
		"value2": 2,
	}
	n := [4]map[string]int{m, m}
	s[0] = n
	s = append(s, n)
	fmt.Println(s)
}
