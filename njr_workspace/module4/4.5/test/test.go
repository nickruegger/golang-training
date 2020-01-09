package test

import (
	"testing"
	"fmt"
)

func TestSqrt(t *testing.T) {
	val, err := sqrt(9)
	if err != nil {
		t.Error(err.Error())
        return
	}
	val2, err2 := sqrt(-9)
	if err2 != nil {
		t.Error(err2.Error())
        return
	}
}
