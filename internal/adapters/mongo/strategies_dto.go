package mongo

import (
	"fmt"

	"github.com/0mithun/go-clean-arch/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type StrategiesDTO struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
}

func fromStrategyCoreToDTO(input *domain.Strategy) (*StrategiesDTO, error) {
	if input == nil {
		return nil, fmt.Errorf("invalid input strategy")
	}

	result := &StrategiesDTO{
		Name:        input.Name,
		Description: input.Description,
	}

	if input.ID != "" {
		id, err := bson.ObjectIDFromHex(input.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid strategy id: '%s'", input.ID)
		}
		result.ID = id
	}

	return result, nil
}

func fromStrategyDTOToCore(input *StrategiesDTO) (*domain.Strategy, error) {
	if input == nil {
		return nil, fmt.Errorf("invalid input dto")
	}

	result := &domain.Strategy{
		ID:          input.ID.Hex(),
		Name:        input.Name,
		Description: input.Description,
	}

	return result, nil
}
