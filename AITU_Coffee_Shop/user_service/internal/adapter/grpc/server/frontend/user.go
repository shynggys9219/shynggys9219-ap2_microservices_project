package frontend

import (
	"context"

	svc "github.com/shynggys9219/ap2-apis-gen-user-service/service/github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1"
)

type User struct {
	svc.UnimplementedClientServiceServer

	uc ClientUsecase
}

func NewUser(uc ClientUsecase) *User {
	return &User{
		uc: uc,
	}
}

func (u *User) Create(ctx context.Context, req *svc.CreateRequest) (*svc.CreateResponse, error) {

	return &svc.CreateResponse{Id: 0}, nil
}
