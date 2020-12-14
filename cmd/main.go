package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"

	"service/pkg/actions"
	"service/pkg/adapters/sql"
	"service/pkg/services"
)

func build() ([]services.Service, error) {

	cfg := buildConfig()

	storage, err := sql.NewPostgresStorage(cfg.PostgresDSN)
	if err != nil {
		return nil, err
	}

	acts := actions.NewActions(storage)

	return services.BuildServices(cfg, acts), nil
}

func run(ctx context.Context, svcs []services.Service) error {
	wg, wgCtx := errgroup.WithContext(ctx)
	for _, service := range svcs {
		wg.Go(func(svc services.Service) func() error {
			return func() error {
				log.Println("starting service", svc.Name())
				return svc.Run(wgCtx)
			}
		}(service))
	}

	return wg.Wait()
}

func runCtx() context.Context {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()
	return ctx
}

func main() {
	svcs, err := build()
	if err != nil {
		panic(err)
	}

	log.Fatal(run(runCtx(), svcs))
}
