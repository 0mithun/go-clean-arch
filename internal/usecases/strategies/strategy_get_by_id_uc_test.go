package strategies

import (
	"context"
	"errors"
	"testing"

	"github.com/0mithun/go-clean-arch/internal/domain"
	"github.com/0mithun/go-clean-arch/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestStrategyGetByIdUCHandle(t *testing.T) {
	t.Parallel()

	t.Run("error getting strategy by id returns error", func(tt *testing.T) {
		tt.Parallel()

		ctx := context.Background()

		strategiesRepo := mocks.NewStrategiesRepositoryMock(tt)
		strategiesRepo.On("GetByID", ctx, "strategy-id").
			Return(nil, errors.New("some error from mongodb")).
			Once()

		uc := NewGetByIDUC(strategiesRepo)
		req := &GetStrategyByIDRequest{
			StrategyID: "strategy-id",
		}

		response, err := uc.Handle(ctx, req)
		assert.Nil(t, response)
		assert.NotNil(tt, err)
		assert.EqualValues(tt, "some error from mongodb", err.Error())
	})

	t.Run("success getting strategy by id", func(tt *testing.T) {
		tt.Parallel()

		ctx := context.Background()

		mockedStrategy := &domain.Strategy{
			ID:          "strategy-id",
			Name:        "name",
			Description: "description",
		}

		strategiesRepo := mocks.NewStrategiesRepositoryMock(tt)
		strategiesRepo.On("GetByID", ctx, "strategy-id").
			Return(mockedStrategy, nil).
			Once()

		uc := NewGetByIDUC(strategiesRepo)

		req := &GetStrategyByIDRequest{
			StrategyID: "strategy-id",
		}

		response, err := uc.Handle(ctx, req)

		assert.Nil(tt, err)
		assert.NotNil(tt, response)
		assert.NotNil(tt, response.Strategy)
		assert.EqualValues(tt, mockedStrategy.ID, response.Strategy.ID)
		assert.EqualValues(tt, mockedStrategy.Name, response.Strategy.Name)
		assert.EqualValues(tt, mockedStrategy.Description, response.Strategy.Description)
	})
}
