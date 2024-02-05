package main

import (
	"os"

	"github.com/anchi205/FileOps/client"
	"github.com/anchi205/FileOps/server"
)

func main() {

	args := os.Args[1:]
	if len(args) != 0 && args[0] == "server" {
		server.Server()

	} else {
		client.Client()
	}
}
