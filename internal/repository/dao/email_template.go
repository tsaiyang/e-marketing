package dao

type EmailTemplate struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 模板名称，方便在系统中识别不同模板
	Name string `gorm:"type:varchar(100);not null"`
	// 邮件主题行，发送邮件时使用的标题
	Subject string `gorm:"type:varchar(200);not null"`
	// 纯文本格式的邮件内容
	Content string `gorm:"type:text;not null"`
	// HTML格式的邮件内容
	HtmlContent string `gorm:"type:mediumtext"`
	// 变量列表，JSON格式
	Variables map[string]string `gorm:"type:json;comment:'变量列表如 {\"name\":\"姓名\",\"company\":\"公司\"}'"`
	// 状态：1启用 0禁用
	Status uint8 `gorm:"type:tinyint;default:1;comment:'1:启用 0:禁用'"`
	// 时间戳
	CreateAt int64
	UpdateAt int64
}

type ScenarioTemplate struct {
	ID         int64 `gorm:"primaryKey;autoIncrement"`
	ScenarioID int64 `gorm:"index:idx_scenario_template"`
	TemplateID int64 `gorm:"index:idx_scenario_template"`
	SeqNo      int
	CreatedAt  int64
	UpdatedAt  int64
}
