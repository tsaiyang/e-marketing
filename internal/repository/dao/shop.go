package dao

import (
	"context"

	"gorm.io/gorm"
)

type ShopDAO interface {
	FindByEmails(ctx context.Context, appName string, emails []string) ([]Shop, error)
}

type gormShopDAO struct {
	db *gorm.DB
}

func (dao *gormShopDAO) FindByEmails(ctx context.Context,
	appName string, emails []string) (shops []Shop, err error) {
	err = dao.db.WithContext(ctx).
		Where("app = ? AND email in (?)", appName, emails).
		Find(&shops).
		Error

	return
}

func NewShopDAO(db *gorm.DB) ShopDAO {
	return &gormShopDAO{db: db}
}

type Shop struct {
	Id          int64  `gorm:"primaryKey;autoIncrement"`
	App         string `gorm:"type:varchar(255);index:idx_app_name,unique,priority:1;not null"`
	Name        string `gorm:"type:varchar(255);index:idx_app_name,unique,priority:2;not null"`
	Email       string `gorm:"type:varchar(255);not null"`
	Info        string `gorm:"type:text;not null"`
	Domain      string `gorm:"type:varchar(255)"`
	AccessToken string `gorm:"type:varchar(255)"`
	IsActive    bool   `gorm:"type:bool;default:true"`
	Scope       string `gorm:"type:varchar(1000)"`
	UninstallAt int64
	ExpireAt    int64
	UpdateAt    int64
	CreateAt    int64
}
