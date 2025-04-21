package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/http/server/handler/dto"
)

type ClientStatistic struct {
	uc ClientStatisticUsecase
}

func NewClientStatistic(uc ClientStatisticUsecase) *ClientStatistic {
	return &ClientStatistic{
		uc: uc,
	}
}

func (c *ClientStatistic) Get(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})

		return
	}

	clientStat, err := c.uc.Get(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.FromClientStatistic(clientStat))
}

func (c *ClientStatistic) List(ctx *gin.Context) {
	clientStat, err := c.uc.List(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.FromClientStatisticToList(clientStat))
}
