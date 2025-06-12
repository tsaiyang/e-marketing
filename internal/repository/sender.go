package repository

import (
	"context"
	"e-marketing/internal/model"
	"e-marketing/internal/repository/dao"
	"time"
)

type SenderRepo interface {
	GetSenderListByPurpose(ctx context.Context, purpose model.SenderPurpose) ([]model.Sender, error)
	GetEmailCountAndLimitTheDay(ctx context.Context, sid int64) (int, int, error)
}

type senderRepo struct {
	senderDAO dao.SenderDAO
}

func (repo *senderRepo) GetEmailCountAndLimitTheDay(ctx context.Context, sid int64) (int, int, error) {
	return repo.senderDAO.GetEmailCountAndLimitTheDay(ctx, sid)
}

func (repo *senderRepo) GetSenderListByPurpose(ctx context.Context,
	purpose model.SenderPurpose) ([]model.Sender, error) {
	senders, err := repo.senderDAO.GetSenderListByPurpose(ctx, string(purpose))
	if err != nil {
		return nil, err
	}

	res := make([]model.Sender, 0, len(senders))
	for _, sender := range senders {
		res = append(res, repo.toModel(sender))
	}

	return res, nil
}

func NewSenderRepo(senderDAO dao.SenderDAO) SenderRepo {
	return &senderRepo{
		senderDAO: senderDAO,
	}
}

func (repo *senderRepo) toModel(sender dao.Sender) model.Sender {
	return model.Sender{
		Id:       sender.Id,
		Name:     sender.Name,
		Email:    sender.Email,
		Purpose:  model.SenderPurpose(sender.Purpose),
		Host:     sender.Host,
		Port:     sender.Port,
		Username: sender.Username,
		Password: sender.Password,
		Status:   model.SenderStatus(sender.Status),
		UpdateAt: time.UnixMilli(sender.UpdateAt),
		CreateAt: time.UnixMilli(sender.CreateAt),
	}
}
