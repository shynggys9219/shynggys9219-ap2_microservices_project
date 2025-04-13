package usecase

import (
	"context"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/def"
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

func (c *Client) Create(ctx context.Context, request model.Client) (model.Client, error) {
	id, err := c.ai.Next(ctx, model.ClientAi)
	if err != nil {
		return model.Client{}, err
	}
	request.ID = id

	err = c.repo.Create(ctx, request)
	if err != nil {
		return model.Client{}, err
	}

	return model.Client{
		ID:    id,
		Email: request.Email,
	}, nil
}

func (c *Client) Update(ctx context.Context, request model.Client) (model.Client, error) {
	updateData := model.ClientUpdateData{
		ID:           def.Pointer(request.ID),
		Name:         def.Pointer(request.Name),
		Phone:        def.Pointer(request.Phone),
		Email:        def.Pointer(request.Email),
		PasswordHash: def.Pointer(request.NewPasswordHash),
		UpdatedAt:    def.Pointer(request.UpdatedAt),
	}
	err := c.repo.Update(ctx, model.ClientFilter{ID: &request.ID}, updateData)
	if err != nil {
		return model.Client{}, err
	}

	updatedClient, err := c.Get(ctx, request.ID)
	if err != nil {
		return model.Client{}, err
	}

	return updatedClient, nil
}

func (c *Client) Get(ctx context.Context, id uint64) (model.Client, error) {
	return c.repo.GetWithFilter(ctx, model.ClientFilter{ID: &id})
}

func (c *Client) Delete(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}
