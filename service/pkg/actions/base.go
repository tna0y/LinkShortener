package actions

import "service/pkg/adapters"

type Actions struct {
	storage adapters.StorageAdapter
}

func NewActions(storage adapters.StorageAdapter) *Actions {
	return &Actions{storage: storage}
}
