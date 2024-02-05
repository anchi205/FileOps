package main

import (
	"fmt"
	"os"

	"github.com/anchi205/FileOps/client"
	"github.com/anchi205/FileOps/server"
)

func main() {

	args := os.Args[1:]
	fmt.Println(args)
	if len(args) != 0 && args[0] == "server" {
		server.Server()

	} else {
		client.Client()
	}
}
