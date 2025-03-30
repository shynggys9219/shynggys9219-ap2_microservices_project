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

func FromClientCreateRequest(ctx *gin.Context) (model.Client, error) {
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

func ToClientCreateResponse(client model.Client) ClientCreateResponse {
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
