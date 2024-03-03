// Code generated from Pkl module `modular.org.PluginConfig`. DO NOT EDIT.
package pluginconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type PluginConfig struct {
	Name string `pkl:"name"`

	Help string `pkl:"help"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a PluginConfig
func LoadFromPath(ctx context.Context, path string) (ret *PluginConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a PluginConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*PluginConfig, error) {
	var ret PluginConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
