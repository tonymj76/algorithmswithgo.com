package main

import (
	"flag"
	"log"
	"plugin"
)

func main() {
	path := flag.String("plugin", "", "Plugin to execute")
	flag.Parse()

	if *path == "" {
		log.Fatal("Path to plugin must be provided")
	}

	p, err := plugin.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	f, err := p.Lookup("ThingToDo")
	if err != nil {
		log.Fatal(err)
	}
	thingToDo, ok := f.(func())
	if !ok {
		log.Fatal("the result of thing to do is not a type of func()")
	}
	thingToDo()

	// to run this file we make use of buildmode plugin
	// go run action.go -plugin=plugin.so
}
