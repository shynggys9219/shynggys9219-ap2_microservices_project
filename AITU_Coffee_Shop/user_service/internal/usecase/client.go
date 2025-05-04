package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/def"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/security"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/transactor"
)

type Customer struct {
	ai         AiRepo
	repo       CustomerRepo
	producer   CustomerEventStorage
	callTx     transactor.WithinTransactionFunc
	jwtManager *security.JWTManager
}

func NewCustomer(
	ai AiRepo,
	repo CustomerRepo,
	producer CustomerEventStorage,
	callTx transactor.WithinTransactionFunc,
	jwtManager *security.JWTManager,
) *Customer {
	return &Customer{
		ai:         ai,
		repo:       repo,
		producer:   producer,
		callTx:     callTx,
		jwtManager: jwtManager,
	}
}

func (uc *Customer) Register(ctx context.Context, request model.Customer) (uint64, error) {
	txFn := func(ctx context.Context) error {
		id, err := uc.ai.Next(ctx, model.CustomerAi)

		if err != nil {
			return err
		}
		request.ID = id

		request.NewPasswordHash, err = uc.hashNewPassword(request.NewPassword)
		if err != nil {
			return fmt.Errorf("uc.hashNewPassword")
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

func (uc *Customer) Update(ctx context.Context, request model.Customer) (model.Customer, error) {
	dbCustomer, err := uc.Get(ctx, request.ID)
	if err != nil {
		return model.Customer{}, err
	}

	err = uc.checkCurrentPassword(request.CurrentPassword, dbCustomer.PasswordHash)
	if err != nil {
		return model.Customer{}, fmt.Errorf("passwords do not match")
	}

	request.NewPasswordHash, err = uc.hashNewPassword(request.NewPassword)
	if err != nil {
		return model.Customer{}, fmt.Errorf("uc.hashNewPassword")
	}

	updateData := model.CustomerUpdateData{
		ID:           def.Pointer(request.ID),
		Name:         def.Pointer(request.Name),
		Phone:        def.Pointer(request.Phone),
		Email:        def.Pointer(request.Email),
		PasswordHash: def.Pointer(request.NewPasswordHash),
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

func (uc *Customer) Login(ctx context.Context, request model.Token) (model.Token, error) {

	return model.Token{}, nil
}
func (uc *Customer) RefreshToken(ctx context.Context, request model.Token) (model.Token, error) {
	return model.Token{}, nil
}

func (uc *Customer) Get(ctx context.Context, id uint64) (model.Customer, error) {
	return uc.repo.GetWithFilter(ctx, model.CustomerFilter{ID: &id})
}

func (uc *Customer) Delete(ctx context.Context, id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (uc *Customer) hashNewPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (uc *Customer) checkCurrentPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
