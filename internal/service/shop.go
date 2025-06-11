package service

import (
	"context"
	"e-marketing/internal/model"
	"e-marketing/internal/repository"
)

type ShopService interface {
	FindByEmails(ctx context.Context, appName string, emails []string) ([]model.Shop, error)
}

type shopService struct {
	shopRepo repository.ShopRepo
}

func (service *shopService) FindByEmails(ctx context.Context, appName string, emails []string) ([]model.Shop, error) {
	return service.shopRepo.FindByEmails(ctx, appName, emails)
}

func NewShopService(shopRepo repository.ShopRepo) ShopService {
	return &shopService{
		shopRepo: shopRepo,
	}
}
