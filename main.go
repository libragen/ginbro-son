package main

import (
	"mojotv/zerg/handlers"
)

func main() {
	defer handlers.Close()
	handlers.ServerRun()
}
