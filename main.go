package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("Starting up")

	hourOfDay := time.Now().Hour()
	greeting, err := getGreeting(hourOfDay)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(greeting)
	logrus.Info("It's Slice time")

	langs := getLangs()

	for _, language := range langs {
		fmt.Println(language)
	}

	logrus.Infof("Normal exit")
}

func getLangs() []string {
	return []string{"Go", "Ruby", "C++", "APL"}
}

func getGreeting(hour int) (string, error) {
	var message string

	if hour < 7 {
		err := errors.New("Too early for Gophers")
		return message, err
	}

	if hour < 12 {
		message = "Good morning"
	} else if hour < 18 {
		message = "Good afternoon"
	} else {
		message = "Good evening"
	}

	return message, nil
}
