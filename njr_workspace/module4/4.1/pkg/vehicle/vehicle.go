package vehicle

import "fmt"

type vehicle interface {
	Make() string
	Model() string
	Wheels() int
}

func Output(v vehicle) {
	fmt.Println(v.Make())
	fmt.Println(v.Model())
	fmt.Println(v.Wheels())
}
