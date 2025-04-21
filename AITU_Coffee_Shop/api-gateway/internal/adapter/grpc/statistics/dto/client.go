package dto

import (
	statsbase "github.com/shynggys9219/ap2-apis-gen-statistics-service/base/frontend/v1"
	svc "github.com/shynggys9219/ap2-apis-gen-statistics-service/service/frontend/client_stats/v1"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
)

func FromGRPCClientGetResponse(resp *svc.GetResponse) model.ClientStatistic {
	return model.ClientStatistic{
		ID:          resp.Client.Id,
		OrdersCount: resp.Client.OrdersCount,
		Orders:      toClientOrders(resp.Client.Orders),
		Phone:       resp.Client.Phone,
		CreatedAt:   resp.Client.CreatedAt.AsTime(),
		UpdatedAt:   resp.Client.UpdatedAt.AsTime(),
		IsDeleted:   resp.Client.IsDeleted,
	}
}

func FROMGRPCClientListResponse(resp *svc.ListResponse) []model.ClientStatistic {
	clientsStatistic := make([]model.ClientStatistic, 0)
	for i := range resp.Clients {

		clientsStatistic = append(clientsStatistic, model.ClientStatistic{
			ID:          resp.Clients[i].Id,
			OrdersCount: resp.Clients[i].OrdersCount,
			Orders:      toClientOrders(resp.Clients[i].Orders),
			Phone:       resp.Clients[i].Phone,
			CreatedAt:   resp.Clients[i].CreatedAt.AsTime(),
			UpdatedAt:   resp.Clients[i].UpdatedAt.AsTime(),
			IsDeleted:   resp.Clients[i].IsDeleted,
		})
	}

	return clientsStatistic
}

func toClientOrders(pbOrders []*statsbase.Order) []*model.Order {
	orders := make([]*model.Order, 0)
	for i := range pbOrders {
		orders = append(orders, toClientOrder(pbOrders[i]))
	}

	return orders
}

func toClientOrder(orderResp *statsbase.Order) *model.Order {
	return &model.Order{
		ID:        orderResp.Id,
		ClientID:  orderResp.ClientId,
		OrderType: toClientOrderType(orderResp.Type),
		Sum:       orderResp.Sum,
		CreatedAt: orderResp.CreatedAt.AsTime(),
		UpdatedAt: orderResp.UpdatedAt.AsTime(),
		Status:    toClientOrderStatus(orderResp.Status),
	}
}

func toClientOrderType(orderRespType statsbase.OrderType) model.OrderType {
	switch orderRespType {
	case statsbase.OrderType_ORDER_TYPE_PRE_ORDER:
		return model.OrderTypePreOrder
	case statsbase.OrderType_ORDER_TYPE_INSTANT:
		return model.OrderTypeInstant
	case statsbase.OrderType_ORDER_TYPE_UNSPECIFIED:
		fallthrough
	default:
		return model.OrderTypeUnspecified
	}
}

func toClientOrderStatus(orderRespStatus statsbase.OrderStatus) model.OrderStatus {
	switch orderRespStatus {
	case statsbase.OrderStatus_ORDER_STATUS_CREATED:
		return model.OrderStatusCreated
	case statsbase.OrderStatus_ORDER_STATUS_PENDING:
		return model.OrderStatusPending
	case statsbase.OrderStatus_ORDER_STATUS_COMPLETED:
		return model.OrderStatusCompleted
	case statsbase.OrderStatus_ORDER_STATUS_UNSPECIFIED:
		fallthrough
	default:
		return model.OrderStatusUnspecified
	}
}
