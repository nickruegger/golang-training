package main

import (
  "github.com/user/project/pkg/hello"
  "github.com/sirupsen/logrus"
)

func main() {
  hello.Greet()
  logrus.Info("hello")
}
