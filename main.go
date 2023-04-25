package main

import (
	"fmt"
	"os/user"

	"github.com/mikepianka/gostarter/example"
)

func main() {
	var msg string

	u, err := user.Current()

	if err != nil {
		msg = example.Greeting("buddy")
	} else {
		msg = example.Greeting(u.Username)
	}

	fmt.Println(msg)
}
