package main

import "fmt"

// HelloWorldPlugin implements the Plugin interface for a simple Helloworld plugin.
type HelloWorldPlugin struct{}

func (p *HelloWorldPlugin) GetName() string {
	return "helloworld"
}

func (p *HelloWorldPlugin) GetExpectedArgs() []string {
	// This simple plugin does not expect any arguments
	return []string{}
}

func (p *HelloWorldPlugin) GetHelp() string {
	return "Prints 'Hello, World!' message"
}

func (p *HelloWorldPlugin) Initialize() error {
	// Initialization logic here, if necessary
	return nil
}

func (p *HelloWorldPlugin) Execute(args map[string]string) error {
	fmt.Println("Hello, World!")
	return nil
}

// Export the plugin instance
var Plugin HelloWorldPlugin
