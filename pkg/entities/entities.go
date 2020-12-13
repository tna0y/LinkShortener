package entities

import (
	"time"
)

type Link struct {
	ID      string
	ShortID string
	OwnerID string
	Target  string
	Created time.Time
	Expires time.Time
}

type Requester struct {
	TelegramID *string
}

var EmptyRequester = Requester{}

func NewRequester(telegramID string) Requester {
	return Requester{TelegramID: &telegramID}
}

func (r *Requester) Owns(link Link) bool {
	return r.TelegramID != nil && link.OwnerID == *r.TelegramID
}
