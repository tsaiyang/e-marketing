package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CursorDAO interface {
	Set(ctx context.Context, name string, offset int64) error
	Get(ctx context.Context, name string) (int64, error)
}

type gormCursorDAO struct {
	db *gorm.DB
}

func (dao *gormCursorDAO) Get(ctx context.Context, name string) (int64, error) {
	var cursor Cursor
	err := dao.db.WithContext(ctx).Where("name = ?", name).First(&cursor).Error

	return cursor.Offset, err
}

func (dao *gormCursorDAO) Set(ctx context.Context, name string, offset int64) error {
	return dao.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "name"}},
		DoUpdates: clause.Assignments(map[string]any{
			"offset": offset,
		}),
	}).Create(&Cursor{
		Name:   name,
		Offset: offset,
	}).Error
}

func NewCursorDAO(db *gorm.DB) CursorDAO {
	return &gormCursorDAO{
		db: db,
	}
}

type Cursor struct {
	Id     int64  `gorm:"primaryKey,autoIncrement"`
	Name   string `gorm:"type:varchar(10);uniqueIndex"`
	Offset int64
}
