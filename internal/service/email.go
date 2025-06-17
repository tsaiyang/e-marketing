package service

import (
	"context"
	"e-marketing/internal/model"
	"e-marketing/internal/repository"
)

type EmailService interface {
	GetNotInstalledOffset(ctx context.Context, name string) (int64, error)
	IncrNotInstalledOffset(ctx context.Context, name string, num int) error
	GetSenderListByPurpose(ctx context.Context, purpose model.SenderPurpose) ([]model.Sender, error)
	GetScenarioByCode(ctx context.Context, code model.ScenarioCode) (model.Scenario, error)
	GetTriggerRuleByScenarioId(ctx context.Context, sid int64) (model.TriggerRule, error)
	GetFrequencyByScenarioId(ctx context.Context, sid int64) (model.SendingFrequency, error)
	GetRecipientList(ctx context.Context, offset int, limit int) ([]model.Recipient, error)
	GetEmailCountAndLimitTheDay(ctx context.Context, sid int64) (int, int, error)
	IncrCountStatTheDay(ctx context.Context, sid int64, count int) error
}

type emailService struct {
	cursorRepo    repository.CursorRepo
	senderRepo    repository.SenderRepo
	scenarioRepo  repository.ScenarioRepo
	recipientRepo repository.RecipientRepo
}

func (service *emailService) IncrNotInstalledOffset(ctx context.Context, name string, num int) error {
	return service.cursorRepo.Incr(ctx, name, num)
}

func (service *emailService) IncrCountStatTheDay(ctx context.Context, sid int64, count int) error {
	return service.senderRepo.IncrCountStatTheDay(ctx, sid, count)
}

func (service *emailService) GetEmailCountAndLimitTheDay(ctx context.Context, sid int64) (int, int, error) {
	return service.senderRepo.GetEmailCountAndLimitTheDay(ctx, sid)
}

func (service *emailService) GetRecipientList(ctx context.Context, offset, limit int) ([]model.Recipient, error) {
	return service.recipientRepo.GetRecipientList(ctx, offset, limit)
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

func (service *emailService) GetNotInstalledOffset(ctx context.Context, name string) (int64, error) {
	return service.cursorRepo.Get(ctx, name)
}

func NewEmailService(
	cursorRepo repository.CursorRepo,
	senderRepo repository.SenderRepo,
	scenarioRepo repository.ScenarioRepo,
	recipientRepo repository.RecipientRepo,
) EmailService {
	return &emailService{
		cursorRepo:    cursorRepo,
		senderRepo:    senderRepo,
		scenarioRepo:  scenarioRepo,
		recipientRepo: recipientRepo,
	}
}
