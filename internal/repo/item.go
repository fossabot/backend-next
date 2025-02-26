package repo

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"

	"exusiai.dev/backend-next/internal/model"
	modelv2 "exusiai.dev/backend-next/internal/model/v2"
	"exusiai.dev/backend-next/internal/pkg/pgerr"
	"exusiai.dev/gommon/constant"
)

type Item struct {
	DB *bun.DB
}

func NewItem(db *bun.DB) *Item {
	return &Item{DB: db}
}

func (c *Item) GetItems(ctx context.Context) ([]*model.Item, error) {
	var items []*model.Item
	err := c.DB.NewSelect().
		Model(&items).
		Order("item_id ASC").
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Item) GetItemById(ctx context.Context, itemId int) (*model.Item, error) {
	var item model.Item
	err := c.DB.NewSelect().
		Model(&item).
		Where("item_id = ?", itemId).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *Item) GetItemByArkId(ctx context.Context, arkItemId string) (*model.Item, error) {
	var item model.Item
	err := c.DB.NewSelect().
		Model(&item).
		Where("ark_item_id = ?", arkItemId).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *Item) GetShimItems(ctx context.Context) ([]*modelv2.Item, error) {
	var items []*modelv2.Item

	err := c.DB.NewSelect().
		Model(&items).
		Order("item_id ASC").
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Item) GetShimItemByArkId(ctx context.Context, itemId string) (*modelv2.Item, error) {
	var item modelv2.Item
	err := c.DB.NewSelect().
		Model(&item).
		Where("ark_item_id = ?", itemId).
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *Item) SearchItemByName(ctx context.Context, name string) (*model.Item, error) {
	var item model.Item
	err := c.DB.NewSelect().
		Model(&item).
		Where("\"name\"::TEXT ILIKE ?", "%"+name+"%").
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &item, nil
}

func (c *Item) GetRecruitTagItems(ctx context.Context) ([]*model.Item, error) {
	var items []*model.Item
	err := c.DB.NewSelect().
		Model(&items).
		Where("type = ?", constant.RecruitItemType).
		Order("item_id ASC").
		Scan(ctx)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pgerr.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return items, nil
}
