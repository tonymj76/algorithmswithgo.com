package main

import "fmt"

func ThingToDo() {
	fmt.Println("Doing the thing...")
}

// first we create a plugin using the following commands
// go build -buildmode=plugin -o=plugin.so src/plugin/plugin.go
