package dao

import (
	"context"
	"e-marketing/pkg/sqlx"

	"gorm.io/gorm"
)

type ScenarioDAO interface {
	GetScenarioByCode(ctx context.Context, code string) (Scenario, error)
	GetTriggerRuleByScenarioId(ctx context.Context, sid int64) (TriggerRule, error)
	GetFrequencyByScenarioId(ctx context.Context, sid int64) (SendingFrequency, error)
}

type gormScenarioDAO struct {
	db *gorm.DB
}

func (dao *gormScenarioDAO) GetScenarioByCode(ctx context.Context,
	code string) (scenario Scenario, err error) {
	err = dao.db.WithContext(ctx).Where("code = ?", code).First(&scenario).Error
	return
}

func (dao *gormScenarioDAO) GetTriggerRuleByScenarioId(ctx context.Context,
	sid int64) (triggerRule TriggerRule, err error) {
	err = dao.db.WithContext(ctx).Where("scenario_id = ?", sid).First(&triggerRule).Error
	return
}

func (dao *gormScenarioDAO) GetFrequencyByScenarioId(ctx context.Context,
	sid int64) (frequency SendingFrequency, err error) {
	err = dao.db.WithContext(ctx).Where("scenario_id = ?", sid).First(&frequency).Error
	return
}

func NewScenarioDAO(db *gorm.DB) ScenarioDAO {
	return &gormScenarioDAO{
		db: db,
	}
}

// 场景
type Scenario struct {
	Id        int64  `gorm:"primaryKey;autoIncrement"`
	App       string `gorm:"type:varchar(20)"`
	Code      string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Name      string `gorm:"type:varchar(100);not null"`
	Objective string `gorm:"type:text"`
	IsActive  bool   `gorm:"default:true"`
	CreatedAt int64
	UpdatedAt int64
}

// 触发规则
type TriggerRule struct {
	Id         int64     `gorm:"primaryKey;autoIncrement"`
	ScenarioId int64     `gorm:"index:idx_trigger_scenario"`
	Type       string    `gorm:"type:varchar(100);not null"`
	Params     sqlx.JSON `gorm:"type:JSON"`
	CreatedAt  int64
	UpdatedAt  int64
}

// SendingFrequency 发送频率
type SendingFrequency struct {
	Id         int64     `gorm:"primaryKey;autoIncrement"`
	ScenarioId int64     `gorm:"uniqueIndex:idx_frequency_scenario"`
	Type       string    `gorm:"type:varchar(50);not null"`
	Params     sqlx.JSON `gorm:"type:JSON"`
	CreatedAt  int64
	UpdatedAt  int64
}
