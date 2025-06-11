package dao

// 邮件模板表
// 发件人表
// 收件人表
// 邮件发送记录表
// 邮件点击跟踪表
// 发送策略表
// 发送统计表

type EmailLog struct {
	Id            int64  `gorm:"primaryKey,autoIncrement"`
	MessageId     string `gorm:"type:varchar(100);uniqueIndex"`
	SenderId      int64  `gorm:"index"`
	RecipientId   int64  `gorm:"index"`
	TemplateId    int64  `gorm:"index"`
	ScenarioId    int64  `gorm:"index"`
	Subject       string `gorm:"type:varchar(255)"`
	Content       string `gorm:"type:text"`
	Status        uint8
	FailReason    string `gorm:"type:varchar(255)"`
	RetryCount    int
	SmtpResponse  string `gorm:"type:varchar(255)"`
	IpAddress     string `gorm:"type:varchar(50)"`
	SendTime      int64
	ScheduledTime int64
	CreateAt      int64
	UpdateAt      int64
}

// 阈值通知发送记录表
type ThresholdNotification struct {
	Id         int64 `gorm:"primaryKey,autoIncrement"`
	EmailLogId int64
	App        string `gorm:"type:varchar(50);index:app_threshold"`
	Threshold  int    `gorm:"index:app_threshold"`
	CreateAt   int64
}
