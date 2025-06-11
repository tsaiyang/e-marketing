package repository

import (
	"context"
	"e-marketing/internal/model"
	"e-marketing/internal/repository/dao"
	"time"
)

type ShopRepo interface {
	FindByEmails(ctx context.Context, appName string, emails []string) ([]model.Shop, error)
}

type shopRepo struct {
	shopDAO dao.ShopDAO
}

func (repo *shopRepo) FindByEmails(ctx context.Context, appName string, emails []string) ([]model.Shop, error) {
	shops, err := repo.shopDAO.FindByEmails(ctx, appName, emails)
	if err != nil {
		return nil, err
	}

	res := make([]model.Shop, 0, len(shops))
	for _, shop := range shops {
		res = append(res, repo.toModel(shop))
	}

	return res, nil
}

func NewShopRepo(shopDAO dao.ShopDAO) ShopRepo {
	return &shopRepo{
		shopDAO: shopDAO,
	}
}

func (repo *shopRepo) toModel(shop dao.Shop) model.Shop {
	return model.Shop{
		Id:          shop.Id,
		App:         shop.App,
		Name:        shop.Name,
		Email:       shop.Email,
		Info:        shop.Info,
		Domain:      shop.Domain,
		AccessToken: shop.AccessToken,
		IsActive:    shop.IsActive,
		Scope:       shop.Scope,
		UninstallAt: time.UnixMilli(shop.UninstallAt),
		ExpireAt:    time.UnixMilli(shop.ExpireAt),
		UpdateAt:    time.UnixMilli(shop.UpdateAt),
		CreateAt:    time.UnixMilli(shop.CreateAt),
	}
}

// func (repo *shopRepo) toEntity(shop model.Shop) dao.Shop {
// 	return dao.Shop{
// 		Id:          shop.Id,
// 		App:         shop.App,
// 		Name:        shop.Name,
// 		Email:       shop.Email,
// 		Info:        shop.Info,
// 		Domain:      shop.Domain,
// 		AccessToken: shop.AccessToken,
// 		IsActive:    shop.IsActive,
// 		Scope:       shop.Scope,
// 		UninstallAt: shop.UninstallAt.UnixMilli(),
// 		ExpireAt:    shop.ExpireAt.UnixMilli(),
// 		UpdateAt:    shop.UpdateAt.UnixMilli(),
// 		CreateAt:    shop.CreateAt.UnixMilli(),
// 	}
// }
