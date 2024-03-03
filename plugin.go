package helloworld

import (
	"context"

	"github.com/renatoaraujo/modular-helloworld-plugin/gen/pluginconfig"
	"github.com/renatoaraujo/modular-helloworld-plugin/internal"
)

// Plugin implements the Plugin interface for a simple Hello world plugin.
type Plugin struct {
	name string
	help string
}

// LoadConfiguration Loads the configuration in the pkl files
func (p *Plugin) LoadConfiguration() {
	cfg, err := pluginconfig.LoadFromPath(context.Background(), "pkl/PluginDefinition.pkl")
	if err != nil {
		panic(err)
	}
	p.name = cfg.Name
	p.help = cfg.Help
}

func (p *Plugin) GetName() string {
	return p.name
}

func (p *Plugin) GetExpectedArgs() []string {
	// This simple plugin does not expect any arguments
	return []string{}
}

func (p *Plugin) GetHelp() string {
	return p.help
}

func (p *Plugin) Initialize() error {
	// Initialization logic here, if necessary
	return nil
}

func (p *Plugin) Execute(args map[string]string) error {
	service := internal.NewService()
	service.SayHello(args["name"])
	return nil
}
