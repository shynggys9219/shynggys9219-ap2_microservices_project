package dto

import (
	base "github.com/shynggys9219/ap2-apis-gen-user-service/base/frontend/v1"
	svc "github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToClientFromCreateRequest(req *svc.CreateRequest) (model.Client, error) {
	passwordHash, err := hashPassword(req.Password)
	if err != nil {
		return model.Client{}, err
	}

	return model.Client{
		Email:        req.Email,
		PasswordHash: passwordHash,
	}, nil
}

func ToClientFromUpdateRequest(req *svc.UpdateRequest) (model.Client, error) {
	oldPasswordHash, err := hashPassword(req.OldPassword)
	if err != nil {
		return model.Client{}, err
	}
	err = checkCurrentPassword(req.OldPassword, oldPasswordHash)
	if err != nil {
		return model.Client{}, err
	}

	newpPasswordHash, err := hashPassword(req.Password)
	if err != nil {
		return model.Client{}, err
	}

	return model.Client{
		ID:              req.Id,
		Name:            req.Name,
		Phone:           req.Phone,
		Email:           req.Email,
		PasswordHash:    oldPasswordHash,
		NewPasswordHash: newpPasswordHash,
	}, nil
}

func FromClient(client model.Client) *base.Client {
	return &base.Client{
		Id:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		CreatedAt: timestamppb.New(client.CreatedAt),
		IsDeleted: client.IsDeleted,
	}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func checkCurrentPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
