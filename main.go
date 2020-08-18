package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting up")

	hourOfDay := time.Now().Hour()
	greeting := getGreeting(hourOfDay)
	fmt.Println(greeting)
	logrus.Infof("Normal exit")
}

func getGreeting(hour int) string {
	if hour < 12 {
		return "Good morning"
	} else if hour < 18 {
		return "Good afternoon"
	}

	return "Good evening"
}
