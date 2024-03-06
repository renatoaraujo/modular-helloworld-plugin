package main

import (
	"embed"
	"github.com/renatoaraujo/modular-helloworld-plugin/internal"
)

//go:embed pkl/PluginConfig.pkl
var cfg embed.FS

// HelloWorldPlugin implements the HelloWorldPlugin interface for a simple Hello world plugin.
type HelloWorldPlugin struct {
	initialized bool
	name        string
	help        string
}

func (p *HelloWorldPlugin) GetName() string {
	return p.name
}

func (p *HelloWorldPlugin) GetExpectedArgs() []string {
	return []string{"name"}
}

func (p *HelloWorldPlugin) GetHelp() string {
	return p.help
}

func (p *HelloWorldPlugin) Initialize() error {
	//sm, _ := cfg.ReadDir("pkl/PluginConfig.pkl")
	//config.LoadFromPath(context.Background(), sm)
	return nil
}

func (p *HelloWorldPlugin) Execute(args map[string]string) error {
	service := internal.NewService()
	service.SayHello(args["name"])
	return nil
}

// Export the plugin instance
var Plugin HelloWorldPlugin
