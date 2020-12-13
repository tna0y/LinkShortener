package actions

import (
	"context"
	"service/pkg/entities"
)

func (a *Actions) ListLinks(ctx context.Context, requester entities.Requester) (result []entities.Link, err error) {
	if requester.TelegramID == nil {
		return nil, entities.ErrNotAuthenticated
	}

	return a.storage.ListLinks(ctx, *requester.TelegramID)
}
