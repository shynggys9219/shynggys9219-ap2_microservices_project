package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type ClientCreateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClientCreateResponse struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}

type ClientUpdateRequest struct {
	ID       uint64  `json:"id"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
	Name     *string `json:"fullName,omitempty"`
	Phone    *string `json:"phone,omitempty"`
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

	passwordHash, err := hashPassword(req.Password)
	if err != nil {
		return model.Client{}, err
	}

	return model.Client{
		Email:        req.Email,
		PasswordHash: passwordHash,
	}, nil
}

func FromClientToCreateResponse(client model.Client) ClientCreateResponse {
	return ClientCreateResponse{
		ID:    client.ID,
		Email: client.Email,
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
	var client model.Client
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return model.Client{}, err

	}

	err = validateClientUpdateRequest(req)
	if err != nil {
		return model.Client{}, err
	}
	if req.Password != nil {
		passwordHash, errHash := hashPassword(*req.Password)
		if errHash != nil {
			return model.Client{}, errHash
		}
		client.PasswordHash = passwordHash
	}

	client.ID = req.ID
	client.Name = *req.Name
	client.Phone = *req.Phone

	return client, nil
}

func FromClientToUpdateResponse(client model.Client) ClientUpdateResponse {
	return ClientUpdateResponse{
		Email: client.Email,
		Name:  client.Name,
		Phone: client.Phone,
	}
}
