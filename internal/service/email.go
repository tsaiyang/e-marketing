package service

import (
	"context"
	"e-marketing/internal/model"
	"e-marketing/internal/repository"
)

type EmailService interface {
	GetNotInstalledOffet(ctx context.Context, name string) (int64, error)
	GetSenderListByPurpose(ctx context.Context, purpose model.SenderPurpose) ([]model.Sender, error)
	GetScenarioByCode(ctx context.Context, code model.ScenarioCode) (model.Scenario, error)
	GetTriggerRuleByScenarioId(ctx context.Context, sid int64) (model.TriggerRule, error)
	GetFrequencyByScenarioId(ctx context.Context, sid int64) (model.SendingFrequency, error)
}

type emailService struct {
	cursorRepo   repository.CursorRepo
	senderRepo   repository.SenderRepo
	scenarioRepo repository.ScenarioRepo
}

func (service *emailService) GetFrequencyByScenarioId(ctx context.Context, sid int64) (model.SendingFrequency, error) {
	return service.scenarioRepo.GetSendingFrequencyByScenarioId(ctx, sid)
}

func (service *emailService) GetTriggerRuleByScenarioId(ctx context.Context, sid int64) (model.TriggerRule, error) {
	return service.scenarioRepo.GetTriggerRuleByScenarioId(ctx, sid)
}

func (service *emailService) GetScenarioByCode(ctx context.Context, code model.ScenarioCode) (model.Scenario, error) {
	return service.scenarioRepo.GetScenarioByCode(ctx, code)
}

func (service *emailService) GetSenderListByPurpose(ctx context.Context, purpose model.SenderPurpose) ([]model.Sender, error) {
	return service.senderRepo.GetSenderListByPurpose(ctx, purpose)
}

func (service *emailService) GetNotInstalledOffet(ctx context.Context, name string) (int64, error) {
	return service.cursorRepo.Get(ctx, name)
}

func NewEmailService(
	cursorRepo repository.CursorRepo,
	senderRepo repository.SenderRepo,
	scenarioRepo repository.ScenarioRepo,
) EmailService {
	return &emailService{
		cursorRepo:   cursorRepo,
		senderRepo:   senderRepo,
		scenarioRepo: scenarioRepo,
	}
}
