package dao

type Recipient struct {
	Id    int64  `gorm:"primaryKey,autoIncrement"`
	Email string `gorm:"type:varchar(255);unique"`
	Name  string `gorm:"type:varchar(255)"`
}
