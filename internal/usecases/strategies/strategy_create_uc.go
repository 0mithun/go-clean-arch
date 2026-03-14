package strategies

import (
	"context"
	"strings"

	"github.com/0mithun/go-clean-arch/internal/domain"
	"github.com/0mithun/go-clean-arch/internal/ports"
	"github.com/0mithun/go-clean-arch/pkg/apierrors"
)

type CreateUC interface {
	Handle(ctx context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error)
}

type strategyCreateUC struct {
	strategiesRepo ports.StrategiesRepository
}

func NewCreateUC(strategiesRepo ports.StrategiesRepository) CreateUC {
	result := &strategyCreateUC{
		strategiesRepo: strategiesRepo,
	}

	return result
}

type CreateStrategyRequest struct {
	Name        string
	Description string
}

type CreateStrategyResponse struct {
	StrategyId string
}

func (uc *strategyCreateUC) Handle(ctx context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error) {
	strategy, err := uc.fromCreateStrategyReqToStrategy(req)
	if err != nil {
		return nil, err
	}

	strategyID, err := uc.strategiesRepo.Insert(ctx, strategy)
	if err != nil {
		return nil, err
	}

	response := &CreateStrategyResponse{
		StrategyId: strategyID,
	}

	return response, nil
}

func (uc *strategyCreateUC) fromCreateStrategyReqToStrategy(req *CreateStrategyRequest) (*domain.Strategy, error) {
	if req == nil {
		return nil, apierrors.NewBadRequestError("invalid request")
	}

	if req.Name = strings.TrimSpace(req.Name); req.Name == "" {
		return nil, apierrors.NewBadRequestError("invalid name")
	}

	strategy := &domain.Strategy{
		Name:        req.Name,
		Description: req.Description,
	}

	return strategy, nil
}
