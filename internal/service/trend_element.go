package service

import (
	"context"

	"exusiai.dev/backend-next/internal/model"
	"exusiai.dev/backend-next/internal/repo"
)

type TrendElement struct {
	TrendElementRepo *repo.TrendElement
}

func NewTrendElement(trendElementRepo *repo.TrendElement) *TrendElement {
	return &TrendElement{
		TrendElementRepo: trendElementRepo,
	}
}

func (s *TrendElement) BatchSaveElements(ctx context.Context, elements []*model.TrendElement, server string) error {
	return s.TrendElementRepo.BatchSaveElements(ctx, elements, server)
}

func (s *TrendElement) DeleteByServer(ctx context.Context, server string) error {
	return s.TrendElementRepo.DeleteByServer(ctx, server)
}

func (s *TrendElement) GetElementsByServerAndSourceCategory(ctx context.Context, server string, sourceCategory string) ([]*model.TrendElement, error) {
	return s.TrendElementRepo.GetElementsByServerAndSourceCategory(ctx, server, sourceCategory)
}
