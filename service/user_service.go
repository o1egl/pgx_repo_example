package service

import (
	"context"
	"errors"
	"fmt"

	"repo_test/repo"
	"repo_test/sqli"
)

type UserService struct {
	readQuerier  sqli.Querier
	writeQuerier sqli.Querier
	userRepo     repo.UserRepo
}

func (s *UserService) Create(ctx context.Context, name string) error {
	return s.writeQuerier.WithTx(func(executor sqli.Executor) error {
		_, err := s.userRepo.FetchUserIDByName(ctx, executor, name)
		if !errors.Is(err, repo.ErrNotFount) {
			fmt.Errorf("user with name %s already exist", name)
		}

		return s.userRepo.Create(ctx, executor, name)
	})
}
