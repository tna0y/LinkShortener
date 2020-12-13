package actions

import (
	"context"
	"service/pkg/entities"
)

func (a *Actions) GetLink(ctx context.Context, ID string, requester entities.Requester) (result entities.Link, err error) {
	if requester.TelegramID == nil {
		err = entities.ErrNotAuthenticated
		return
	}

	result, err = a.storage.GetLink(ctx, ID)
	if err != nil {
		return
	}

	if !requester.Owns(result) {
		result = entities.Link{}
		err = entities.ErrNotFound
	}
	return
}
