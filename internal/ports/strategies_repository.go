package ports

import (
	"context"

	"github.com/0mithun/go-clean-arch/internal/domain"
)

type StrategiesRepository interface {
	Insert(ctx context.Context, strategy *domain.Strategy) (string, error)
	GetByID(ctx context.Context, id string) (*domain.Strategy, error)
}
