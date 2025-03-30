package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/http/service/handler/dto"
)

type Client struct {
	uc ClientUsecase
}

func NewClient(uc ClientUsecase) *Client {
	return &Client{
		uc: uc,
	}
}

func (c *Client) Create(ctx *gin.Context) {
	client, err := dto.FromClientCreateRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})

		return
	}

	newClient, err := c.uc.Create(ctx.Request.Context(), client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.ToClientCreateResponse(newClient))
}
