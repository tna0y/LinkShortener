package services

import (
	"context"

	"service/pkg/actions"
	"service/pkg/config"
	"service/pkg/services/bot"
	"service/pkg/services/redirector"
)

type Service interface {
	Name() string
	Run(ctx context.Context) error
}

func BuildServices(cfg config.Config, actions *actions.Actions) []Service {
	return []Service{
		bot.NewBot(cfg, actions),
		redirector.NewRedirector(cfg, actions),
	}
}
