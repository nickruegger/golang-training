package main

import (
	"./pkg/vehicle"
)


type car struct {}

func (c car) Make() string {
	return "Ford"
}

func (c car) Model() string {
	return "Pinto"
}

func (c car) Wheels() int {
	return 4
}

type motorcycle struct {}

func (m motorcycle) Make() string {
	return "Kawasaki"
}

func (m motorcycle) Model() string {
	return "Ninja"
}

func (m motorcycle) Wheels() int {
	return 2
}


func main() {
	ford := car{}
	ninja := motorcycle{}
	vehicle.Output(ford)
	vehicle.Output(ninja)
}
