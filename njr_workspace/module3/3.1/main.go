package main

import (
 "fmt"
 "strconv"
)

func main() {

	var (
		i1, i2 int
		i3, i4 int64
		e error
	)
	fmt.Println(fmt.Sprintf("%d", &i1))
	fmt.Println(fmt.Sprintf("%d", &i2))
	addr1, e := strconv.Atoi(fmt.Sprintf("%d", &i1))
	addr2, e := strconv.Atoi(fmt.Sprintf("%d", &i2))
	addr3, e := strconv.Atoi(fmt.Sprintf("%d", &i3))
	addr4, e := strconv.Atoi(fmt.Sprintf("%d", &i4))
	if e == nil {
		intSize := addr2 - addr1
		fmt.Println("int size ", intSize)
		intSize64 := addr4 - addr3
		fmt.Println("int64 size ", intSize64)
	}

}
