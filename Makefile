PHONY: config

config:
	pkl-gen-go pkl/PluginConfig.pkl

build:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 go build -buildmode=plugin -o /Users/renato/.modular/plugins/helloworld.so plugin.go
