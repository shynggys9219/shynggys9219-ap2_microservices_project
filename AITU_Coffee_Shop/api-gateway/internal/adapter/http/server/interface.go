package service

import (
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/http/server/handler"
)

type ClientUsecase interface {
	handler.ClientUsecase
}
