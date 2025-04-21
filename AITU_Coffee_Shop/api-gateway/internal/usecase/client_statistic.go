package usecase

import (
	"context"

	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/model"
)

type ClientStatistic struct {
	clientStatisticsPresenter ClientStatisticPresenter
}

func NewClientStatistic(clientStatisticsPresenter ClientStatisticPresenter) *ClientStatistic {
	return &ClientStatistic{
		clientStatisticsPresenter: clientStatisticsPresenter,
	}
}

func (c *ClientStatistic) Get(ctx context.Context, id uint64) (model.ClientStatistic, error) {
	return c.clientStatisticsPresenter.Get(ctx, id)
}

func (c *ClientStatistic) List(ctx context.Context) ([]model.ClientStatistic, error) {
	return c.clientStatisticsPresenter.List(ctx)
}
