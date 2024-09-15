package main

import (
	"Learning-Go/monkey-v1/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to start typing commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
