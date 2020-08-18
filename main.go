package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting up")
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	} else {
		fmt.Println("Hello, I am Gopher!")
	}

	logrus.Infof("Normal exit")
}
