package strategies

import (
	"net/http"
	"strings"

	strategiesUC "github.com/0mithun/go-clean-arch/internal/usecases/strategies"
	"github.com/0mithun/go-clean-arch/pkg/apierrors"
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create(ctx *gin.Context)
	GetById(ctx *gin.Context)
}

type handlers struct {
	svc strategiesUC.Service
}

func NewHandlers(svc strategiesUC.Service) Handlers {
	result := &handlers{
		svc: svc,
	}

	return result
}

func (h handlers) Create(ctx *gin.Context) {
	var request CreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		apiErr := apierrors.NewBadRequestError("invalid json body")

		ctx.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
	}

	svcReq := &strategiesUC.CreateStrategyRequest{
		Name:        request.Name,
		Description: request.Description,
	}

	res, err := h.svc.Create(ctx, svcReq)
	if err != nil {
		apiErr := apierrors.FromError(err)

		ctx.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
	}

	createResponse := &CreateResponse{
		StrategyId: res.StrategyId,
	}

	ctx.JSON(http.StatusCreated, createResponse)
}

func (h handlers) GetById(ctx *gin.Context) {
	svcReq := &strategiesUC.GetStrategyByIDRequest{
		StrategyID: strings.TrimSpace(ctx.Param("strategy_id")),
	}
	res, err := h.svc.GetById(ctx, svcReq)
	if err != nil {
		apiErr := apierrors.FromError(err)
		ctx.AbortWithStatusJSON(apiErr.StatusCode(), apiErr)
	}
	getResponse := &GetByIDResponse{
		Strategy: fromStrategyCoreToHTTP(res.Strategy),
	}
	ctx.JSON(http.StatusOK, getResponse)
}
