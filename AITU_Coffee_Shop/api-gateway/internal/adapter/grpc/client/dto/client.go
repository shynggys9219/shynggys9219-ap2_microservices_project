package dto

import (
	svc "github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func FromGRPCClientCreateResponse(resp *svc.CreateResponse) model.Client {
	return model.Client{ID: resp.Id}
}

func FromGRPCClientGetResponse(resp *svc.GetResponse) model.Client {

	return model.Client{
		ID:        resp.Client.Id,
		Name:      resp.Client.Name,
		Phone:     resp.Client.Phone,
		Email:     resp.Client.Email,
		CreatedAt: FromPbTime(resp.Client.CreatedAt),
		IsDeleted: resp.Client.IsDeleted,
	}
}

func FromPbTime(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{} // or handle nil as needed
	}

	return ts.AsTime()
}
