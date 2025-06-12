package dao

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"

	"gorm.io/gorm"
)

type ScenarioDAO interface {
	GetScenarioByCode(ctx context.Context, code string) (*Scenario, error)
	GetTriggerRuleByScenarioId(ctx context.Context, sid int64) (TriggerRule, error)
	GetFrequencyByScenarioId(ctx context.Context, sid int64) (SendingFrequency, error)
}

type gormScenarioDAO struct {
	db *gorm.DB
}

func NewScenarioDAO(db *gorm.DB) ScenarioDAO {
	return &gormScenarioDAO{
		db: db,
	}
}

// 场景
type Scenario struct {
	Id        int64  `gorm:"primaryKey;autoIncrement"`
	Code      string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Name      string `gorm:"type:varchar(100);not null"`
	Objective string `gorm:"type:text"`
	IsActive  bool   `gorm:"default:true"`
	CreatedAt int64
	UpdatedAt int64
}

// 触发规则
type TriggerRule struct {
	Id          int64  `gorm:"primaryKey;autoIncrement"`
	ScenarioId  int64  `gorm:"index:idx_trigger_scenario"`
	TriggerType string `gorm:"type:varchar(100);not null"`
	Params      JSON   `gorm:"type:json"`
	CreatedAt   int64
	UpdatedAt   int64
}

// SendingFrequency 发送频率
type SendingFrequency struct {
	ID            int64  `gorm:"primaryKey;autoIncrement"`
	ScenarioID    int64  `gorm:"uniqueIndex:idx_frequency_scenario"`
	FrequencyType string `gorm:"type:varchar(50);not null"`
	Parameters    JSON   `gorm:"type:json"`
	CreatedAt     int64
	UpdatedAt     int64
}

// JSON类型处理
type JSON map[string]any

func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JSON) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, &j)
}
