package repository

import (
	"context"
	"e-marketing/internal/repository/dao"
)

type CursorRepo interface {
	Set(ctx context.Context, name string, offset int64) error
	Get(ctx context.Context, name string) (int64, error)
}

type cursorRepo struct {
	cursorDAO dao.CursorDAO
}

func (repo *cursorRepo) Get(ctx context.Context, name string) (int64, error) {
	return repo.cursorDAO.Get(ctx, name)
}

func (repo *cursorRepo) Set(ctx context.Context, name string, offset int64) error {
	return repo.cursorDAO.Set(ctx, name, offset)
}

func NewCursorRepo(cursorDAO dao.CursorDAO) CursorRepo {
	return &cursorRepo{
		cursorDAO: cursorDAO,
	}
}
