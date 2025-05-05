package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/def"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/security"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/transactor"
)

type Customer struct {
	ai              AiRepo
	repo            CustomerRepo
	tokenRepo       RefreshTokenRepo
	producer        CustomerEventStorage
	callTx          transactor.WithinTransactionFunc
	jwtManager      *security.JWTManager
	passwordManager *security.PasswordManager
}

func NewCustomer(
	ai AiRepo,
	repo CustomerRepo,
	tokenRepo RefreshTokenRepo,
	producer CustomerEventStorage,
	callTx transactor.WithinTransactionFunc,
	jwtManager *security.JWTManager,
	passwordManager *security.PasswordManager,
) *Customer {
	return &Customer{
		ai:              ai,
		repo:            repo,
		tokenRepo:       tokenRepo,
		producer:        producer,
		callTx:          callTx,
		jwtManager:      jwtManager,
		passwordManager: passwordManager,
	}
}

func (uc *Customer) Register(ctx context.Context, request model.Customer) (uint64, error) {
	txFn := func(ctx context.Context) error {
		id, err := uc.ai.Next(ctx, model.CustomerAi)

		if err != nil {
			return err
		}
		request.ID = id

		request.PasswordHash, err = uc.passwordManager.HashPassword(request.NewPassword)
		if err != nil {
			return fmt.Errorf("uc.passwordManager.HashPassword")
		}

		request.CreatedAt = time.Now().UTC()
		request.UpdatedAt = time.Now().UTC()
		err = uc.repo.Create(ctx, request)
		if err != nil {
			return err
		}

		err = uc.producer.Push(ctx, request)
		if err != nil {
			log.Println("uc.producer.Push: %w", err)
		}

		return nil
	}

	err := uc.callTx(ctx, txFn)
	if err != nil {
		return 0, fmt.Errorf("uc.callTx: %w", err)
	}

	return request.ID, nil
}

func (uc *Customer) Update(ctx context.Context, token string, request model.Customer) (model.Customer, error) {
	claims, err := uc.jwtManager.Verify(token)
	if err != nil {
		return model.Customer{}, model.ErrInvalidID
	}
	customerID, ok := claims["user_id"].(float64)
	if !ok {
		return model.Customer{}, model.ErrInvalidID
	}
	if uint64(customerID) != request.ID {
		return model.Customer{}, model.ErrInvalidID
	}

	dbCustomer, err := uc.Get(ctx, token, request.ID)
	if err != nil {
		return model.Customer{}, err
	}

	err = uc.passwordManager.CheckPassword(dbCustomer.PasswordHash, request.CurrentPassword)
	if err != nil {
		return model.Customer{}, fmt.Errorf("uc.passwordManager.CheckPassword: %w", err)
	}

	request.PasswordHash, err = uc.passwordManager.HashPassword(request.NewPassword)
	if err != nil {
		return model.Customer{}, fmt.Errorf("uc.passwordManager.HashPassword: %w", err)
	}

	updateData := model.CustomerUpdateData{
		ID:           def.Pointer(request.ID),
		Name:         def.Pointer(request.Name),
		Phone:        def.Pointer(request.Phone),
		Email:        def.Pointer(request.Email),
		PasswordHash: def.Pointer(request.PasswordHash),
		UpdatedAt:    def.Pointer(request.UpdatedAt),
	}

	err = uc.repo.Update(ctx, model.CustomerFilter{ID: &request.ID}, updateData)
	if err != nil {
		return model.Customer{}, err
	}

	err = uc.producer.Push(ctx, request)
	if err != nil {
		log.Println("uc.producer.Push: %w", err)
	}

	return request, nil
}

func (uc *Customer) Get(ctx context.Context, token string, id uint64) (model.Customer, error) {
	claims, err := uc.jwtManager.Verify(token)
	if err != nil {
		return model.Customer{}, model.ErrInvalidID
	}

	customerID, ok := claims["user_id"].(float64)
	if !ok {
		return model.Customer{}, model.ErrInvalidID
	}

	if uint64(customerID) != id {
		return model.Customer{}, model.ErrInvalidID
	}

	return uc.repo.GetWithFilter(ctx, model.CustomerFilter{ID: &id})
}

func (uc *Customer) Delete(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (uc *Customer) Login(ctx context.Context, email, password string) (model.Token, error) {
	customer, err := uc.repo.GetWithFilter(ctx, model.CustomerFilter{Email: def.Pointer(email)})
	if err != nil {
		return model.Token{}, err
	}

	err = uc.passwordManager.CheckPassword(customer.PasswordHash, password)
	if err != nil {
		return model.Token{}, err
	}

	accessToken, err := uc.jwtManager.GenerateAccessToken(customer.ID, model.CustomerRole)
	if err != nil {
		return model.Token{}, err
	}
	refreshToken, err := uc.jwtManager.GenerateRefreshToken(customer.ID)
	if err != nil {
		return model.Token{}, err
	}

	session := model.Session{
		UserID:       customer.ID,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	err = uc.tokenRepo.Create(ctx, session)
	if err != nil {
		return model.Token{}, err
	}

	return model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *Customer) RefreshToken(ctx context.Context, refreshToken string) (model.Token, error) {
	session, err := uc.tokenRepo.GetByToken(ctx, refreshToken)
	if err != nil {
		return model.Token{}, err
	}
	if session.ExpiresAt.Before(time.Now().UTC()) {
		return model.Token{}, model.ErrRefreshTokenExpired
	}

	customer, err := uc.repo.GetWithFilter(ctx, model.CustomerFilter{ID: def.Pointer(session.UserID)})
	if err != nil {
		return model.Token{}, err
	}

	accessToken, err := uc.jwtManager.GenerateAccessToken(customer.ID, model.CustomerRole)
	if err != nil {
		return model.Token{}, err
	}

	newRefreshToken, err := uc.jwtManager.GenerateRefreshToken(customer.ID)
	if err != nil {
		return model.Token{}, err
	}

	// delete old refresh and insert new one (rotation)
	err = uc.tokenRepo.DeleteByToken(ctx, refreshToken)
	if err != nil {
		return model.Token{}, err
	}

	newSession := model.Session{
		UserID:       customer.ID,
		RefreshToken: newRefreshToken,
		ExpiresAt:    time.Now().Add(7 * 24 * time.Hour),
		CreatedAt:    time.Now(),
	}

	err = uc.tokenRepo.Create(ctx, newSession)
	if err != nil {
		return model.Token{}, err
	}

	return model.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
