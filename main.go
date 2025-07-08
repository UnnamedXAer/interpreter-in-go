package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/unnamedxaer/interpreter-in-go/repl"
)

func main() {
	print("interpreter in go\n")

	user, err := user.Current()

	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Fell free to type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}

// let func = fn(a,b,c) { if (1== a) {return b;} else {c + 2*7;}};
