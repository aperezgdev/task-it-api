package main

import "github.com/aperezgdev/task-it-api/cmd/bootstrap"

func main() {
	if err := bootstrap.Run(); err != nil {
		panic(err)
	}
}