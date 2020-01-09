package main

import (
	"fmt"
	"errors"
	"math"
	"testing"
)

func main() {
	val, err := Sqrt(9)
	if err != nil {
		fmt.Println(err.Error())
        return
	}
	fmt.Println("Squar root of 9 = ", val)
	val2, err2 := Sqrt(-9)
	if err2 != nil {
		fmt.Println(err2.Error())
        return
	}
	fmt.Println("Squar root of -9 = ", val2)
}

func Sqrt(val float64) (float64, error) {
	if val <= 0 {
		return 0, errors.New("Value must be greater than 0")
	}
	return math.Sqrt(val), nil
}
