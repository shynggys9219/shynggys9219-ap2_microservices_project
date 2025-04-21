package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Client struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`

	IsDeleted bool `json:"isDeleted"`
}

type ClientCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClientCreateResponse struct {
	ID uint64 `json:"id"`
}

type ClientUpdateRequest struct {
	ID          uint64  `json:"id"`
	Email       *string `json:"email,omitempty"`
	Password    *string `json:"currentPassword,omitempty"`
	NewPassword *string `json:"newPassword,omitempty"`
	Name        *string `json:"fullName,omitempty"`
	Phone       *string `json:"phone,omitempty"`
}

type ClientUpdateResponse struct {
	Email string `json:"email"`
	Name  string `json:"fullName"`
	Phone string `json:"phone"`
}

func ToClientFromCreateRequest(ctx *gin.Context) (model.Client, error) {
	var req ClientCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return model.Client{}, err

	}

	err = validateClientCreateRequest(req)
	if err != nil {
		return model.Client{}, err
	}

	return model.Client{
		Email:           req.Email,
		CurrentPassword: req.Password,
	}, nil
}

func FromClientToCreateResponse(client model.Client) ClientCreateResponse {
	return ClientCreateResponse{
		ID: client.ID,
	}
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ToClientFromUpdateRequest(ctx *gin.Context) (model.Client, error) {
	var req ClientUpdateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return model.Client{}, err

	}

	err = validateClientUpdateRequest(req)
	if err != nil {
		return model.Client{}, err
	}

	return model.Client{
		ID:              req.ID,
		Name:            *req.Name,
		Phone:           *req.Phone,
		Email:           *req.Email,
		CurrentPassword: *req.Password,
		NewPassword:     *req.NewPassword,
	}, nil
}

func FromClientToUpdateResponse(client model.Client) ClientUpdateResponse {
	return ClientUpdateResponse{
		Email: client.Email,
		Name:  client.Name,
		Phone: client.Phone,
	}
}

func FromClient(client model.Client) Client {
	return Client{
		ID:        client.ID,
		Name:      client.Name,
		Phone:     client.Phone,
		Email:     client.Email,
		IsDeleted: client.IsDeleted,
	}
}
