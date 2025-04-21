package usecase

import (
	"context"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/def"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Client struct {
	ai       AiRepo
	repo     ClientRepo
	producer ClientEventStorage
}

func NewUser(ai AiRepo, repo ClientRepo, producer ClientEventStorage) *Client {
	return &Client{
		ai:       ai,
		repo:     repo,
		producer: producer,
	}
}

func (c *Client) Create(ctx context.Context, request model.Client) (model.Client, error) {
	id, err := c.ai.Next(ctx, model.ClientAi)
	if err != nil {
		return model.Client{}, err
	}
	request.ID = id

	request.NewPasswordHash, err = c.hashNewPassword(request.NewPassword)
	if err != nil {
		return model.Client{}, fmt.Errorf("c.hashNewPassword")
	}

	err = c.repo.Create(ctx, request)
	if err != nil {
		return model.Client{}, err
	}

	newClient := model.Client{
		ID:    id,
		Email: request.Email,
	}

	err = c.producer.Push(ctx, newClient)
	if err != nil {
		log.Println("c.producer.Push: %w", err)
	}

	return newClient, nil
}

func (c *Client) Update(ctx context.Context, request model.Client) (model.Client, error) {
	dbClient, err := c.Get(ctx, request.ID)
	if err != nil {
		return model.Client{}, err
	}

	err = c.checkCurrentPassword(request.CurrentPassword, dbClient.PasswordHash)
	if err != nil {
		return model.Client{}, fmt.Errorf("passwords do not match")
	}

	request.NewPasswordHash, err = c.hashNewPassword(request.NewPassword)
	if err != nil {
		return model.Client{}, fmt.Errorf("c.hashNewPassword")
	}

	updateData := model.ClientUpdateData{
		ID:           def.Pointer(request.ID),
		Name:         def.Pointer(request.Name),
		Phone:        def.Pointer(request.Phone),
		Email:        def.Pointer(request.Email),
		PasswordHash: def.Pointer(request.NewPasswordHash),
		UpdatedAt:    def.Pointer(request.UpdatedAt),
	}

	err = c.repo.Update(ctx, model.ClientFilter{ID: &request.ID}, updateData)
	if err != nil {
		return model.Client{}, err
	}

	err = c.producer.Push(ctx, request)
	if err != nil {
		log.Println("c.producer.Push: %w", err)
	}

	return request, nil
}

func (c *Client) Get(ctx context.Context, id uint64) (model.Client, error) {
	return c.repo.GetWithFilter(ctx, model.ClientFilter{ID: &id})
}

func (c *Client) Delete(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (c *Client) hashNewPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (c *Client) checkCurrentPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
