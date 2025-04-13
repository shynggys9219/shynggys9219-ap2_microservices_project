package backoffice

import (
	svc "github.com/shynggys9219/ap2-apis-user-service/generated/github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1"
)

type User struct {
	svc.UnimplementedClientServiceServer

	uc ClientUsecase
}

func NewUser() *User {

}
