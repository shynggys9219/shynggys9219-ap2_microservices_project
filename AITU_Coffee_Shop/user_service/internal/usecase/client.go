package usecase

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
)

type Client struct {
	ai   AiRepo
	repo ClientRepo
}

func NewUser(ai AiRepo, repo ClientRepo) *Client {
	return &Client{
		ai:   ai,
		repo: repo,
	}
}

func (u *Client) Create(ctx context.Context, request model.Client) (model.Client, error) {
	id, err := u.ai.Next(ctx, model.ClientAi)
	if err != nil {
		return model.Client{}, err
	}

	err = u.repo.Create(ctx, request)
	if err != nil {
		return model.Client{}, err
	}

	return model.Client{
		ID:    id,
		Email: request.Email,
	}, nil
}

func (u *Client) Update(ctx context.Context, request model.Client) (model.Client, error) {
	//TODO implement me
	panic("implement me")
}

func (u *Client) Get(ctx context.Context, id uint64) (model.Client, error) {
	//TODO implement me
	panic("implement me")
}

func (u *Client) Delete(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}
