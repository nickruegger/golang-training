package main

import (
	"fmt"                           // stdlib
	"math/rand"                     // stdlib
	"github.com/sirupsen/logrus"    // remote package
)

func main() {
	fmt.Println("hi", rand.Int())
	logrus.Info("Complete")
}
