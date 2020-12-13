package sqlite

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"service/pkg/entities"
	"time"
)

type Link struct {
	ID      uuid.UUID `gorm:"pk"`
	ShortID string    `gorm:"index"`
	OwnerID string    `gorm:"index"`
	Target  string
	Start   sql.NullTime
	End     sql.NullTime
}

func parseLink(link *Link) entities.Link {
	if link == nil {
		return entities.Link{}
	}
	expires := time.Time{}
	if link.End.Valid {
		expires = link.End.Time
	}

	return entities.Link{
		ID:      link.ID.String(),
		ShortID: link.ShortID,
		OwnerID: link.OwnerID,
		Target:  link.Target,
		Created: link.Start.Time,
		Expires: expires,
	}
}

func serializeLink(link entities.Link) *Link {
	var endTime sql.NullTime

	if link.Expires.IsZero() {
		endTime = sql.NullTime{
			Valid: false,
		}
	} else {
		endTime = sql.NullTime{
			Time:  link.Expires,
			Valid: true,
		}
	}

	return &Link{
		ID:      uuid.FromStringOrNil(link.ID),
		ShortID: link.ShortID,
		OwnerID: link.OwnerID,
		Target:  link.Target,
		Start: sql.NullTime{
			Time:  link.Created,
			Valid: true,
		},
		End: endTime,
	}
}
