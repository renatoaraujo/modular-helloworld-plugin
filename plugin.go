package main

import (
	"context"

	"github.com/renatoaraujo/modular-helloworld-plugin/gen/pluginconfig"
	"github.com/renatoaraujo/modular-helloworld-plugin/internal"
)

// HelloWorldPlugin implements the HelloWorldPlugin interface for a simple Hello world plugin.
type HelloWorldPlugin struct {
	name string
	help string
}

// LoadConfiguration Loads the configuration in the pkl files
func (p *HelloWorldPlugin) LoadConfiguration() {
	cfg, err := pluginconfig.LoadFromPath(context.Background(), "pkl/PluginDefinition.pkl")
	if err != nil {
		panic(err)
	}
	p.name = cfg.Name
	p.help = cfg.Help
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
	// Initialization logic here, if necessary
	return nil
}

func (p *HelloWorldPlugin) Execute(args map[string]string) error {
	service := internal.NewService()
	service.SayHello(args["name"])
	return nil
}

// Export the plugin instance
var Plugin HelloWorldPlugin
