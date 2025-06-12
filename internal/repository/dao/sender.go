package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type SenderDAO interface {
	GetSenderListByPurpose(ctx context.Context, purpose string) ([]Sender, error)
	GetEmailCountAndLimitTheDay(ctx context.Context, sid int64) (int, int, error)
}

type gormSenderDAO struct {
	db *gorm.DB
}

func (dao *gormSenderDAO) GetEmailCountAndLimitTheDay(ctx context.Context, sid int64) (int, int, error) {
	// 将今天转成 YYYY-MM-DD 格式
	now := time.Now()
	date := now.Format("2006-01-02")

	// 查询当天发送数量
	var count int64
	if err := dao.db.WithContext(ctx).
		Model(&SenderDailyStat{}).
		Where("sender_id = ? AND date = ?", sid, date).
		Count(&count).
		Error; err != nil {
		return 0, 0, err
	}

	// 查询发件人信息以获取创建时间
	var sender Sender
	if err := dao.db.WithContext(ctx).
		Select("create_at").
		Where("id = ?", sid).
		First(&sender).Error; err != nil {
		return 0, 0, err
	}

	// 计算创建日期到今天的天数差
	createTime := time.UnixMilli(sender.CreateAt)
	daysSinceCreation := int(now.Sub(createTime).Hours() / 24)

	// 计算当前是第几周（从 1 开始）
	weekNumber := min(daysSinceCreation/7+1, 10)

	// strategy 和 senderStrategy 联合查询，根据 week 和 senderId 查询DailyLimited
	var strategy Strategy
	if err := dao.db.WithContext(ctx).
		Model(&Strategy{}).
		Joins("JOIN sender_strategies ON strategies.id = sender_strategies.strategy_id").
		Where("sender_strategies.sender_id = ? AND strategies.week = ?", sid, weekNumber).
		First(&strategy).Error; err != nil {
		return 0, 0, err
	}

	// 返回当天已发送数量和当日的限制数量
	return int(count), strategy.DailyLimited, nil
}

func (dao *gormSenderDAO) GetSenderListByPurpose(ctx context.Context,
	purpose string) (senders []Sender, err error) {
	err = dao.db.WithContext(ctx).Where("purpose = ?", purpose).Find(&senders).Error

	return
}

func NewSenderDAO(db *gorm.DB) SenderDAO {
	return &gormSenderDAO{
		db: db,
	}
}

type Sender struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Name     string `gorm:"type:varchar(50)"`
	Email    string `gorm:"type:varchar(100);unique"`
	Purpose  string `gorm:"type:varchar(20)"`
	Host     string `gorm:"type:varchar(100)"`
	Port     int
	Username string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(255)"`
	Status   uint8
	UpdateAt int64
	CreateAt int64
}

type Strategy struct {
	Id           int64  `gorm:"primaryKey,autoIncrement"`
	Name         string `gorm:"type:varchar(50)"`
	Week         int
	DailyLimited int
	Status       uint8
	UpdateAt     int64
	CreateAt     int64
}

type SenderStrategy struct {
	Id         int64 `gorm:"primaryKey,autoIncrement"`
	SenderId   int64
	StrategyId int64
	UpdateAt   int64
	CreateAt   int64
}

type SenderDailyStat struct {
	Id       int64 `gorm:"primaryKey,autoIncrement"`
	SenderId int64
	Date     string `gorm:"type:varchar(10)"` // 格式：YYYY-MM-DD
	Count    int    // 当天已发送数量
	Week     int    // 所属周数（可选，便于统计）
	UpdateAt int64
	CreateAt int64
}
