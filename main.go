package main

import (
	"asteroid/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	// Start the REPL
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Asteroid programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
