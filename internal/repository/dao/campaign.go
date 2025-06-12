package dao

type Campaign struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`

	// 基本信息
	Name        string `gorm:"type:varchar(100);not null"`   // 活动名称
	Description string `gorm:"type:text"`                    // 活动描述
	Code        string `gorm:"type:varchar(50);uniqueIndex"` // 活动代码，唯一

	// 时间控制
	StartTime int64 // 活动开始时间
	EndTime   int64 // 活动结束时间

	// 发送配置
	SenderId   int64 `gorm:"index"` // 关联的发件人ID
	ScenarioId int64 `gorm:"index"` // 关联的场景ID

	// 目标受众
	TargetType  string `gorm:"type:varchar(30)"`  // 目标受众类型：all, segment, tag等
	TargetValue string `gorm:"type:varchar(255)"` // 目标受众值：segment_id或tag名称

	// 状态控制
	Status      uint8 `gorm:"type:tinyint;default:0;comment:'0:草稿 1:待发送 2:发送中 3:已完成 4:已暂停 5:已取消'"` // 活动状态
	IsAutomatic bool  // 是否为自动化活动

	// 时间戳
	CreateAt int64
	UpdateAt int64
}
