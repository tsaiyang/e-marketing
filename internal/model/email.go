package model

import "time"

type Sender struct {
	Id       int64
	Name     string
	Email    string
	Purpose  SenderPurpose
	Host     string
	Port     int
	Username string
	Password string
	Status   SenderStatus
	UpdateAt time.Time
	CreateAt time.Time
}

type SenderStatus uint8

const (
	SenderStatusActive SenderStatus = iota + 1
	SenderStatusInactive
)

func (s SenderStatus) ToUint8() uint8 {
	return uint8(s)
}

type SenderPurpose string

const (
	SenderPurposeUniverse  SenderPurpose = "universe"
	SenderPurposeInstalled SenderPurpose = "installed"
)

const EmailNumPerTime = 50

type ScenarioCode string

const (
	ScenarioCodeNotInstalled         ScenarioCode = "not_installed"
	ScenarioCodeWelcomeInstalled     ScenarioCode = "welcome_installed"
	ScenarioCodeNotConfigured        ScenarioCode = "not_configured"
	ScenarioCodeAdvancedFunction     ScenarioCode = "advanced_function"
	ScenarioCodeDeclinedUsage        ScenarioCode = "declined_usage"
	ScenarioCodeInteactive           ScenarioCode = "interactive"
	ScenarioCodeDealDone             ScenarioCode = "deal_done"
	ScenarioCodeUninstalledFeedback  ScenarioCode = "uninstalled_feedback"
	ScenarioCodeUninstalledAfterDays ScenarioCode = "uninstalled_after_days"
	ScenarioCodeUpdateCallback       ScenarioCode = "update_callback"
	ScenarioCodeReconnected          ScenarioCode = "reconnected"
)
