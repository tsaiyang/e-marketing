package repository

import (
	"context"
	"e-marketing/internal/repository/dao"
)

type CursorRepo interface {
	Incr(ctx context.Context, name string, num int) error
	Get(ctx context.Context, name string) (int64, error)
}

type cursorRepo struct {
	cursorDAO dao.CursorDAO
}

func (repo *cursorRepo) Get(ctx context.Context, name string) (int64, error) {
	return repo.cursorDAO.Get(ctx, name)
}

func (repo *cursorRepo) Incr(ctx context.Context, name string, num int) error {
	return repo.cursorDAO.Incr(ctx, name, num)
}

func NewCursorRepo(cursorDAO dao.CursorDAO) CursorRepo {
	return &cursorRepo{
		cursorDAO: cursorDAO,
	}
}
