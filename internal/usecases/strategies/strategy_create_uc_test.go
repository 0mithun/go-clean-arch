package strategies

import (
	"context"
	"errors"
	"testing"

	"github.com/0mithun/go-clean-arch/internal/domain"
	"github.com/0mithun/go-clean-arch/internal/mocks"
	"github.com/0mithun/go-clean-arch/pkg/apierrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFromCreateStrategyReqToStrategy(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		Input       *CreateStrategyRequest
		Output      *domain.Strategy
		ExpectedErr error
	}{
		{
			Name:        "nil request",
			Input:       nil,
			Output:      nil,
			ExpectedErr: apierrors.NewBadRequestError("invalid request"),
		},
		{
			Name: "invalid strategy name",
			Input: &CreateStrategyRequest{
				Name: "",
			},
			Output:      nil,
			ExpectedErr: apierrors.NewBadRequestError("invalid name"),
		},
		{
			Name: "strategy successfully obtained",
			Input: &CreateStrategyRequest{
				Name:        "test",
				Description: "description",
			},
			Output: &domain.Strategy{
				Name:        "test",
				Description: "description",
			},
			ExpectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			uc := &strategyCreateUC{}
			output, err := uc.fromCreateStrategyReqToStrategy(testCase.Input)

			if testCase.ExpectedErr != nil {
				assert.Nil(t, output)
				assert.NotNil(t, err.Error())
				assert.Equal(t, testCase.ExpectedErr.Error(), err.Error())
				return
			}

			assert.Nil(t, err)
			assert.NotNil(t, output)
			assert.EqualValues(t, testCase.Output.Name, output.Name)
			assert.EqualValues(t, testCase.Output.Description, output.Description)
		})
	}
}

func TestCreateStrategyUCHandler(t *testing.T) {
	t.Parallel()

	t.Run("invalid request returns error", func(tt *testing.T) {
		tt.Parallel()

		uc := &strategyCreateUC{}
		ctx := context.Background()
		req := &CreateStrategyRequest{}
		res, err := uc.Handle(ctx, req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.EqualValues(t, "invalid name", err.Error())
	})

	t.Run("error creating strategy returns error", func(tt *testing.T) {
		tt.Parallel()
		ctx := context.Background()
		strategiesRepo := mocks.NewStrategiesRepositoryMock(tt)
		strategiesRepo.On("Insert", ctx, mock.Anything).Return("", errors.New("some db error when inserting strategy")).Once()

		uc := &strategyCreateUC{
			strategiesRepo: strategiesRepo,
		}

		req := &CreateStrategyRequest{
			Name: "test",
		}
		res, err := uc.Handle(ctx, req)

		assert.Nil(t, res)
		assert.NotNil(t, err)
		assert.EqualValues(t, "some db error when inserting strategy", err.Error())
	})

	t.Run("strategy successfully created", func(tt *testing.T) {
		tt.Parallel()

		ctx := context.Background()
		strategiesRepo := mocks.NewStrategiesRepositoryMock(tt)
		strategiesRepo.On("Insert", ctx, mock.Anything).Return("strategy-id", nil).Once()

		uc := &strategyCreateUC{
			strategiesRepo: strategiesRepo,
		}
		req := &CreateStrategyRequest{
			Name:        "test",
			Description: "description",
		}
		res, err := uc.Handle(ctx, req)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.EqualValues(tt, "strategy-id", res.StrategyId)
	})
}
