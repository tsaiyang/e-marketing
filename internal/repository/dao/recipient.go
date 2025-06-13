package dao

import (
	"context"

	"gorm.io/gorm"
)

type RecipientDAO interface {
	GetRecipientList(ctx context.Context, offset int, limit int) ([]Recipient, error)
}

type gormRecipientDAO struct {
	db *gorm.DB
}

func (dao *gormRecipientDAO) GetRecipientList(ctx context.Context,
	offset int, limit int) (recipients []Recipient, err error) {
	err = dao.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&recipients).Error
	return
}

func NewRecipientDAO(db *gorm.DB) RecipientDAO {
	return &gormRecipientDAO{
		db: db,
	}
}

type Recipient struct {
	Id       int64          `gorm:"primaryKey,autoIncrement"`
	Email    string         `gorm:"type:varchar(255);uniqueIndex"`
	Name     string         `gorm:"type:varchar(50)"`
	Company  string         `gorm:"varchar(100)"`
	Position string         `gorm:"varchar(100)"`
	Attrs    map[string]any `gorm:"type:json"`
	// 1:正常 0:退订 -1:黑名单
	Status   uint8
	UpdateAt int64
	CreateAt int64
}
