package adapters

import (
	"context"
	"service/pkg/entities"
)

type StorageAdapter interface {
	CreateLink(ctx context.Context, link entities.Link) (entities.Link, error)
	GetLink(ctx context.Context, ID string) (entities.Link, error)
	DeleteLink(ctx context.Context, link entities.Link) error
	ListLinks(ctx context.Context, ownerID string) ([]entities.Link, error)
}
