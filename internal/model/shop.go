package model

import "time"

type Shop struct {
	Id          int64
	App         string
	Name        string
	Email       string
	Info        string
	Domain      string
	AccessToken string
	IsActive    bool
	Scope       string
	UninstallAt time.Time
	ExpireAt    time.Time
	UpdateAt    time.Time
	CreateAt    time.Time
}
