package main

import (
	"fmt"
	"log"
	"log/slog"
	"share-basket/personal/core"
	"share-basket/personal/presentation/server"
	"share-basket/personal/registry"
)

func main() {
	conf := core.LoadConfig()
	logger := core.NewLogger(conf.Env)
	slog.SetDefault(logger.Logger)

	container, err := registry.Inject()
	if err != nil {
		log.Fatalf("failed to inject dependencies: %v", err)
	}

	services := registry.NewServices(container)

	srv := server.New(fmt.Sprintf(":%s", conf.Port))
	srv.MapServices(services)
	srv.Run()
}
