package strategies

import (
	"testing"

	"github.com/0mithun/go-clean-arch/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestFromStrategyCoreToHTTP(t *testing.T) {
	t.Parallel()

	t.Run("nil input returns nil", func(tt *testing.T) {
		tt.Parallel()
		output := fromStrategyCoreToHTTP(nil)
		assert.Nil(tt, output)
	})

	t.Run("strategy successfully transformed", func(tt *testing.T) {
		input := &domain.Strategy{
			ID:          "id",
			Name:        "name",
			Description: "description",
		}
		output := fromStrategyCoreToHTTP(input)
		assert.NotNil(tt, output)
		assert.Equal(tt, input.ID, output.ID)
		assert.Equal(tt, input.Name, output.Name)
		assert.Equal(tt, input.Description, output.Description)
	})
}
