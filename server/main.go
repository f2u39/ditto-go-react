package main

import (
	"ditto/core"
)

func main() {
	// Create and initialize core components
	core.Init()

	// Run server
	core.Ditto.Engine.Run()
}
