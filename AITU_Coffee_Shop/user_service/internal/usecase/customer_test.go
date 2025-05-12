package usecase_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/usecase"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/usecase/mocks"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/security"
)

func TestCustomer_Register(t *testing.T) {
	ctx := context.Background()
	//now := time.Now().UTC()
	id := uint64(1)

	tests := []struct {
		name      string
		ctx       context.Context
		aiRepo    func() usecase.AiRepo
		repo      func() usecase.CustomerRepo
		producer  func() usecase.CustomerEventStorage
		pwManager *security.PasswordManager
		callTx    func(context.Context, func(context.Context) error) error
		request   model.Customer
		wantID    uint64
		wantErr   error
	}{
		{
			name: "Success",
			ctx:  ctx,
			aiRepo: func() usecase.AiRepo {
				m := mocks.NewAiRepo(t)
				m.EXPECT().Next(ctx, model.CustomerAi).Return(id, nil).Once()
				return m
			},
			repo: func() usecase.CustomerRepo {
				m := mocks.NewCustomerRepo(t)
				m.EXPECT().Create(ctx, mock.Anything).Return(nil).Once()
				return m
			},
			producer: func() usecase.CustomerEventStorage {
				m := mocks.NewCustomerEventStorage(t)
				m.EXPECT().Push(ctx, mock.Anything).Return(nil).Once()
				return m
			},
			pwManager: security.NewPasswordManager(),
			callTx: func(_ context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
			request: model.Customer{
				NewPassword: "new-pass",
			},
			wantID:  id,
			wantErr: nil,
		},
		{
			name: "AI repo error",
			ctx:  ctx,
			aiRepo: func() usecase.AiRepo {
				m := mocks.NewAiRepo(t)
				m.EXPECT().Next(ctx, model.CustomerAi).Return(uint64(0), errors.New("ai error")).Once()
				return m
			},
			repo:      func() usecase.CustomerRepo { return mocks.NewCustomerRepo(t) },
			producer:  func() usecase.CustomerEventStorage { return mocks.NewCustomerEventStorage(t) },
			pwManager: security.NewPasswordManager(),
			callTx: func(_ context.Context, fn func(context.Context) error) error {
				return fn(ctx)
			},
			request: model.Customer{NewPassword: "x"},
			wantID:  0,
			wantErr: errors.New("ai error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := usecase.NewCustomer(
				tt.aiRepo(), tt.repo(), nil, tt.producer(), tt.callTx, nil, tt.pwManager,
			)

			gotID, gotErr := uc.Register(tt.ctx, tt.request)
			if gotErr != nil {
				require.ErrorContains(t, gotErr, tt.wantErr.Error())
			}
			require.Equal(t, tt.wantID, gotID)
		})
	}
}
