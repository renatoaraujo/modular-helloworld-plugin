package main

import "fmt"

var _ Plugin = (*HelloWorld)(nil)

type HelloWorld struct{}

func (h *HelloWorld) GetExpectedArgs() []string {
	return []string{}
}

func (h *HelloWorld) GetHelp() string {
	return "Prints 'Hello, World!' to the console."
}

func (h *HelloWorld) Initialize() error {
	fmt.Println("HelloWorld plugin initialized")
	return nil
}

func (h *HelloWorld) Execute(args map[string]string) error {
	fmt.Println("Hello, World from the plugin!")
	return nil
}

var Plugin HelloWorld
