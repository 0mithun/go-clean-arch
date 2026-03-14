package mongo

import (
	"errors"
	"testing"

	"github.com/0mithun/go-clean-arch/internal/domain"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestFromStrategyCoreToDTO(t *testing.T) {
	t.Parallel()

	mockedMongoID := "69b446b1ba689b09d045a1d6"
	mockedObjectID, err := bson.ObjectIDFromHex(mockedMongoID)
	assert.Nil(t, err)
	testCases := []struct {
		Name        string
		Input       *domain.Strategy
		Output      *StrategiesDTO
		ExpectedErr error
	}{
		{
			Name:        "nil input returns error",
			Input:       nil,
			Output:      nil,
			ExpectedErr: errors.New("invalid input strategy"),
		},
		{
			Name: "invalid mongo id",
			Input: &domain.Strategy{
				ID:          "invalid mongo id",
				Name:        "name",
				Description: "description",
			},
			Output:      nil,
			ExpectedErr: errors.New("invalid strategy id: 'invalid mongo id'"),
		},
		{
			Name: "strategy successfully processed",
			Input: &domain.Strategy{
				ID:          mockedMongoID,
				Name:        "name",
				Description: "description",
			},
			Output: &StrategiesDTO{
				ID:          mockedObjectID,
				Name:        "name",
				Description: "description",
			},
			ExpectedErr: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(tt *testing.T) {
			tt.Parallel()
			output, err := fromStrategyCoreToDTO(testCase.Input)
			if testCase.ExpectedErr != nil {
				assert.Nil(tt, output)
				assert.NotNil(tt, err.Error())
				assert.EqualValues(tt, testCase.ExpectedErr.Error(), err.Error())
				return
			}

			assert.Nil(tt, err)
			assert.NotNil(tt, output)
			assert.EqualValues(tt, testCase.Input.ID, output.ID.Hex())
			assert.EqualValues(tt, testCase.Input.Name, output.Name)
			assert.EqualValues(tt, testCase.Input.Description, output.Description)
		})
	}
}

func TestFromStrategyDTOToCore(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name        string
		Input       *StrategiesDTO
		Output      *domain.Strategy
		ExpectedErr error
	}{
		{
			Name:        "nil input returns error",
			Input:       nil,
			Output:      nil,
			ExpectedErr: errors.New("invalid input dto"),
		},
		{
			Name:        "dto successfully processed",
			Input:       &StrategiesDTO{},
			Output:      &domain.Strategy{},
			ExpectedErr: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Name, func(tt *testing.T) {
			tt.Parallel()

			output, err := fromStrategyDTOToCore(testCase.Input)

			if testCase.ExpectedErr != nil {
				assert.Nil(tt, output)
				assert.NotNil(tt, err)
				assert.EqualValues(tt, testCase.ExpectedErr.Error(), err.Error())
				return
			}

			assert.Nil(tt, err)
			assert.NotNil(tt, output)
			assert.EqualValues(tt, testCase.Input.ID.Hex(), output.ID)
			assert.EqualValues(tt, testCase.Input.Name, output.Name)
			assert.EqualValues(tt, testCase.Input.Description, output.Description)
		})
	}
}
