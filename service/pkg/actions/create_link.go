package actions

import (
	"context"
	"regexp"
	"time"

	uuid "github.com/satori/go.uuid"

	"service/pkg/entities"
)

type CreateLinkArgs struct {
	ShortID string
	Target  string
	TTL     int64
}

func (a *Actions) CreateLink(
	ctx context.Context, args CreateLinkArgs, requester entities.Requester,
) (res entities.Link, err error) {
	if requester.TelegramID == nil {
		err = entities.ErrNotAuthenticated
		return
	}

	var match bool
	match, err = regexp.MatchString("\\w+", args.ShortID)
	if err != nil {
		return
	}
	if !match {
		err = entities.ErrInvalidShortID
		return
	}

	if !validateURL(args.Target) {
		err = entities.ErrInvalidTargetURL
		return
	}

	created := time.Now().UTC()
	var expires time.Time
	if args.TTL > 0 {
		expires = created.Add(time.Second * time.Duration(args.TTL))
	}

	creating := entities.Link{
		ID:      uuid.NewV4().String(),
		ShortID: args.ShortID,
		OwnerID: *requester.TelegramID,
		Target:  args.Target,
		Created: created,
		Expires: expires,
	}

	return a.storage.CreateLink(ctx, creating)
}
