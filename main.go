package main

import (
	"fmt"
	"goodb/repl"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is goodb - a toy database system!\n", currentUser.Username)
	repl.Start(os.Stdin, os.Stdout)
}
