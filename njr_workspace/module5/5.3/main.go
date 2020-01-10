package main

import (
    "fmt"
    "sync"
)

var wg1 sync.WaitGroup
var lock sync.Mutex

func main() {
	list := make([]int, 100)
	for i := range list {
		list[i] = i+1
	}
	
	for i := 0; i < len(list); i +=10 {
		processTen(list[i:i+10])
	}
	
}

func processTen(a []int) {
	lock.Lock()
	for _, v := range a {
		wg1.Add(1)
		go printNum(v)
	}
	wg1.Wait()
	lock.Unlock()
}

func printNum(i int) {
	fmt.Println(i)
	wg1.Done()
}
