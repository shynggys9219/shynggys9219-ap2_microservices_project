package dto

import (
	"fmt"
	base "github.com/shynggys9219/ap2-apis-gen-statistics-service/base/frontend/v1"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func FromClientToStatistic(client model.Client) *base.Statistic {
	pbOrders := make([]*base.Order, 0)
	for i := range client.Orders {
		pbOrder, err := fromOrder(client.Orders[i])
		if err != nil {
			// TODO: choose what is better - to give error or skip iteration
			log.Println("wrong order", err)

			continue
		}
		pbOrders = append(pbOrders, pbOrder)
	}

	return &base.Statistic{
		Id:          client.ID,
		OrdersCount: client.OrdersCount,
		Orders:      pbOrders,
		Phone:       client.Phone,
		CreatedAt:   timestamppb.New(client.CreatedAt),
		UpdatedAt:   timestamppb.New(client.UpdatedAt),
		IsDeleted:   client.IsDeleted,
	}
}

func fromOrder(order *model.Order) (*base.Order, error) {
	orderType, err := fromOrderType(order.OrderType)
	if err != nil {
		return nil, err
	}

	orderStatus, err := fromOrderStatus(order.Status)
	if err != nil {
		return nil, err
	}

	return &base.Order{
		Id:        order.ID,
		ClientId:  order.ClientID,
		Type:      orderType,
		Sum:       order.Sum,
		CreatedAt: timestamppb.New(order.CreatedAt),
		UpdatedAt: timestamppb.New(order.UpdatedAt),
		Status:    orderStatus,
	}, nil
}

func fromOrderType(orderType model.OrderType) (base.OrderType, error) {
	switch orderType {
	case model.OrderTypePreOrder:
		return base.OrderType_ORDER_TYPE_PRE_ORDER, nil
	case model.OrderTypeInstant:
		return base.OrderType_ORDER_TYPE_INSTANT, nil
	case model.OrderTypeUnspecified:
		fallthrough
	default:
		return base.OrderType_ORDER_TYPE_UNSPECIFIED, fmt.Errorf("invalid order type")
	}
}

func fromOrderStatus(orderStatus model.OrderStatus) (base.OrderStatus, error) {
	switch orderStatus {
	case model.OrderStatusCreated:
		return base.OrderStatus_ORDER_STATUS_CREATED, nil
	case model.OrderStatusPending:
		return base.OrderStatus_ORDER_STATUS_PENDING, nil
	case model.OrderStatusCompleted:
		return base.OrderStatus_ORDER_STATUS_COMPLETED, nil
	case model.OrderStatusUnspecified:
		fallthrough
	default:
		return base.OrderStatus_ORDER_STATUS_UNSPECIFIED, fmt.Errorf("invalid order status")
	}
}
