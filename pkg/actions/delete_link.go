package actions

import (
	"context"
	"service/pkg/entities"
)

func (a *Actions) DeleteLink(ctx context.Context, shortID string, requester entities.Requester) (err error) {
	if requester.TelegramID == nil {
		err = entities.ErrNotAuthenticated
		return
	}

	var link entities.Link
	link, err = a.storage.GetLink(ctx, shortID)
	if err != nil {
		return err
	}

	if !requester.Owns(link) {
		return entities.ErrPermissionDenied
	}

	return a.storage.DeleteLink(ctx, link)
}
