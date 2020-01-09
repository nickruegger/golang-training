package main

import "fmt"

func main() {
	
	prntName := func() string {
		return "Jim"
	}

	testf(prntName)
	
}

func testf(f func() string) {
	name := f()
	fmt.Println(name)
	funcRet := returnFunc()
	name2 := funcRet()
	fmt.Println(name2)
}

func returnFunc() func() string{
	bob := func() string{
		return "Bob"
	}
	return bob
}
