package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting up")

	args := os.Args
	if len(args) > 1 {
		fmt.Println(args)
	} else {
		fmt.Println("Hello, I am Gopher!")
	}

	logrus.Infof("Normal exit")
}
