package strategies

import "github.com/0mithun/go-clean-arch/internal/domain"

func fromStrategyCoreToHTTP(input *domain.Strategy) *Strategy {
	if input == nil {
		return nil
	}

	strategy := &Strategy{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
	}

	return strategy
}
