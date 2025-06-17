package repository

import (
	"context"
	"e-marketing/internal/model"
	"e-marketing/internal/repository/dao"
	"time"
)

type RecipientRepo interface {
	GetRecipientList(ctx context.Context, offset int, limit int) ([]model.Recipient, error)
}

type recipientRepo struct {
	recipientDAO dao.RecipientDAO
}

func (repo *recipientRepo) GetRecipientList(ctx context.Context, offset int, limit int) ([]model.Recipient, error) {
	recipients, err := repo.recipientDAO.GetRecipientList(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	res := make([]model.Recipient, 0, len(recipients))
	for _, recipient := range recipients {
		res = append(res, repo.toModel(recipient))
	}

	return res, nil
}

func NewRecipientRepo(recipientDAO dao.RecipientDAO) RecipientRepo {
	return &recipientRepo{
		recipientDAO: recipientDAO,
	}
}

func (repo *recipientRepo) toModel(recipient dao.Recipient) model.Recipient {
	return model.Recipient{
		Id:       recipient.Id,
		Email:    recipient.Email,
		Name:     recipient.Name,
		Company:  recipient.Company,
		Position: recipient.Position,
		Attrs:    recipient.Attrs,
		Status:   model.RecipientStatus(recipient.Status),
		UpdateAt: time.UnixMilli(recipient.UpdateAt),
		CreateAt: time.UnixMilli(recipient.CreateAt),
	}
}
