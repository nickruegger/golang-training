package test

import (
	"sync"
	"testing"
)

func TestInterfaces(t *testing.T) {
	var (
		s   = make(chan string)
		ret []string
		_   sync.WaitGroup
		_   sync.WaitGroup
	)

	go func(retPtr *[]string) {
		*retPtr = append(*retPtr, <-s)
	}(&ret)

	// Add WaitGroups to ensure "hello" executes first.

	go func() {
		s <- "goodbye"
	}()

	go func() {
		s <- "hello"
	}()

	if len(ret) == 0 {
		t.Error("nothing in ret")
		return
	}

	if ret[0] != "hello" {
		t.Error("hello is not", ret[0])
	}
}
