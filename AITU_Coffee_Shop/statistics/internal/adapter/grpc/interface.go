package server

import "github.com/shynggys9219/ap2_microservices_project/statistics/internal/adapter/grpc/server/frontend"

type ClientUsecase interface {
	frontend.ClientUsecase
}
