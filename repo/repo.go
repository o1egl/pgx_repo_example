package repo

import (
	"context"
	"errors"

	"github.com/jackc/pgx"

	"repo_test/sqli"
)

var ErrNotFount = errors.New("not found")

type UserRepo interface {
	Create(ctx context.Context, executor sqli.Executor, name string) error
	FetchUserIDByName(ctx context.Context, executor sqli.Executor, name string) (int64, error)
}

type userRepo struct{}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (u *userRepo) Create(ctx context.Context, executor sqli.Executor, name string) error {
	_, err := executor.ExecEx(ctx, "INSERT INTO users(name) VALUES ($1)", nil, name)
	return err
}

func (u *userRepo) FetchUserIDByName(ctx context.Context, executor sqli.Executor, name string) (int64, error) {
	var id int64
	err := executor.QueryRowEx(ctx, "SELECT id FROM users WHERE name = $1", nil, name).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, ErrNotFount
		}
		return 0, err
	}
	return id, nil
}
