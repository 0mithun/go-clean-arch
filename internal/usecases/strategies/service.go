package strategies

import (
	"context"

	"github.com/0mithun/go-clean-arch/internal/ports"
)

type Service interface {
	Create(ctx context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error)
	GetById(ctx context.Context, req *GetStrategyByIDRequest) (*GetStrategyByIDResponse, error)
}

type service struct {
	createUseCase  CreateUC
	getByIDUseCase GetByIDUC
}

func NewService(strategiesRepo ports.StrategiesRepository) Service {
	createUseCase := NewCreateUC(strategiesRepo)
	getByIDUseCase := NewGetByIDUC(strategiesRepo)

	svc := &service{
		createUseCase:  createUseCase,
		getByIDUseCase: getByIDUseCase,
	}

	return svc
}

func (s *service) Create(ctx context.Context, req *CreateStrategyRequest) (*CreateStrategyResponse, error) {
	return s.createUseCase.Handle(ctx, req)
}

func (s *service) GetById(ctx context.Context, req *GetStrategyByIDRequest) (*GetStrategyByIDResponse, error) {
	return s.getByIDUseCase.Handle(ctx, req)
}
