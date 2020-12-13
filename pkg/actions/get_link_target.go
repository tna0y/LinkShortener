package actions

import (
	"context"
)

func (a *Actions) GetLinkTarget(ctx context.Context, ShortID string) (result string, err error) {
	link, err := a.storage.GetLink(ctx, ShortID)
	if err != nil {
		return "", err
	}
	return link.Target, nil
}
