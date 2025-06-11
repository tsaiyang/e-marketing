package dao

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
