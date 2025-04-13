package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/http/server/handler/dto"
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
	client, err := dto.ToClientFromCreateRequest(ctx)
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

	ctx.JSON(http.StatusOK, dto.FromClientToCreateResponse(newClient))
}

func (c *Client) Update(ctx *gin.Context) {
	client, err := dto.ToClientFromUpdateRequest(ctx)
	if err != nil {
		errCtx := dto.FromError(err)
		ctx.JSON(errCtx.Code, gin.H{"error": errCtx.Message})

		return
	}

	updatedClient, err := c.uc.Update(ctx.Request.Context(), client)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.FromClientToUpdateResponse(updatedClient))
}

func (c *Client) Get(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})

		return
	}

	client, err := c.uc.Get(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, dto.FromClient(client))
}
