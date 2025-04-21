package dto

import (
	"time"

	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
)

type Order struct {
	ID        uint64            `json:"id"`
	ClientID  uint64            `json:"clientID"`
	OrderType model.OrderType   `json:"orderType"`
	Sum       float32           `json:"sum"`
	CreatedAt time.Time         `json:"createdAt"`
	UpdatedAt time.Time         `json:"UpdatedAt"`
	Status    model.OrderStatus `json:"status"`
}
