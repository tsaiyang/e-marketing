package service

import (
	"context"
	"e-marketing/internal/model"
	"e-marketing/internal/repository"
)

type EmailService interface {
	GetNotInstalledOffet(ctx context.Context, name string) (int64, error)
	GetSenderListByPurpose(ctx context.Context, purpose model.SenderPurpose) ([]model.Sender, error)
}

type emailService struct {
	cursorRepo repository.CursorRepo
	senderRepo repository.SenderRepo
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
) EmailService {
	return &emailService{
		cursorRepo: cursorRepo,
		senderRepo: senderRepo,
	}
}
