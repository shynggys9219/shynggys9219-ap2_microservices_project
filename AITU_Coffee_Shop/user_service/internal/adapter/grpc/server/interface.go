package server

import "github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/grpc/server/frontend"

type ClientUsecase interface {
	frontend.ClientUsecase
}
