package dao

import "gorm.io/gorm"

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&EmailLog{},
		&EmailTemplate{},
		&Recipient{},
		&Sender{},
		&Strategy{},
		&SenderStrategy{},
		&SenderDailyStat{},
		&Shop{},
	)
}
