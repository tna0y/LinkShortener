package sql

import (
	"context"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"service/pkg/adapters"
	"service/pkg/entities"
)

type sqliteStorage struct {
	db *gorm.DB
}

func NewSQLiteStorage(path string) (adapters.StorageAdapter, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Link{})
	if err != nil {
		return nil, err
	}

	return &sqliteStorage{
		db: db,
	}, nil
}

func NewPostgresStorage(path string) (adapters.StorageAdapter, error) {
	db, err := gorm.Open(postgres.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Link{})
	if err != nil {
		return nil, err
	}

	return &sqliteStorage{
		db: db,
	}, nil
}

func (p *sqliteStorage) CreateLink(ctx context.Context, link entities.Link) (res entities.Link, err error) {

	out := serializeLink(link)

	err = p.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		var existing []Link
		result := tx.Where("short_id = ? AND start <= ? AND (\"end\" > ? OR \"end\" is NULL)", out.ShortID, out.Start, out.Start).Find(&existing)
		if result.Error != nil {
			return result.Error
		}
		if len(existing) > 0 {
			return entities.ErrExists
		}

		return tx.Create(&out).Error
	})
	if err != nil {
		return
	}

	res = parseLink(out)
	return
}

func (p *sqliteStorage) GetLink(ctx context.Context, ID string) (res entities.Link, err error) {
	var out []Link
	result := p.db.WithContext(ctx).Where("short_id = ? AND (\"end\" > ? OR \"end\" is NULL)", ID, time.Now()).Find(&out)
	if result.Error != nil {
		err = result.Error
		return
	}
	if len(out) == 0 {
		err = entities.ErrNotFound
		return
	}
	res = parseLink(&out[0])
	return
}

func (p *sqliteStorage) DeleteLink(ctx context.Context, link entities.Link) (err error) {
	r := serializeLink(link)

	result := p.db.WithContext(ctx).Delete(&r)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return entities.ErrNotFound
	}
	return nil
}

func (p *sqliteStorage) ListLinks(ctx context.Context, ownerID string) (res []entities.Link, err error) {
	var out []Link
	result := p.db.WithContext(ctx).Where("owner_id = ? AND (\"end\" > ? OR \"end\" is NULL)", ownerID, time.Now()).Find(&out)
	if result.Error != nil {
		return nil, result.Error
	}
	res = make([]entities.Link, 0, len(out))
	for _, item := range out {
		res = append(res, parseLink(&item))
	}
	return
}
