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

type Scenario struct {
	Id        int64
	App       string
	Code      ScenarioCode
	Name      string
	Objective string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TriggerRuleType string

const (
	TriggerRuleTypeNotInstalled           TriggerRuleType = "rule_type_not_installed"
	TriggerRuleTypeInstalledSucceed       TriggerRuleType = "installed_succeed"
	TriggerRuleTypeNotConfigured          TriggerRuleType = "not_configured"
	TriggerRuleTypeFreeUserFrequently     TriggerRuleType = "free_user_frequently"
	TriggerRuleTypeNoActionAfterInstalled TriggerRuleType = "no_action_after_installed"
	TriggerRuleTypeBeyondActionCount      TriggerRuleType = "beyond_action_count"
	TriggerRuleTypeDealDone               TriggerRuleType = "deal_done"
	TriggerRuleTypeUninstalled            TriggerRuleType = "uninstalled"
	TriggerRuleTypeNoReinstalledAfterDays TriggerRuleType = "no_reinstalled_after_days"
	TriggerRuleTypeUpdateCallback         TriggerRuleType = "update_callback"
)

type TriggerRule struct {
	Id         int64
	ScenarioId int64
	Type       TriggerRuleType
	Params     map[string]any
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type SendingFrequencyType string

const (
	SendingFrequencyTypeOneTime        SendingFrequencyType = "one_time"
	SendingFrequencyTypeInterval       SendingFrequencyType = "interval"
	SendingFrequencyTypeDaysAfterEvent SendingFrequencyType = "days_after_event"
	SendingFrequencyTypeRepeat         SendingFrequencyType = "repeat"
	SendingFrequencyTypeOnClick        SendingFrequencyType = "on_click"
	SendingFrequencyTypeOnSettlement   SendingFrequencyType = "on_settlement"
)

type SendingFrequency struct {
	Id         int64
	ScenarioId int64
	Type       SendingFrequencyType
	Params     map[string]any
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
