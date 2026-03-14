package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/0mithun/go-clean-arch/internal/domain"
	"github.com/0mithun/go-clean-arch/internal/ports"
	"github.com/0mithun/go-clean-arch/pkg/apierrors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type strategiesRepo struct {
	strategiesColl *mongo.Collection
}

func NewStrategiesRepository(strategiesColl *mongo.Collection) ports.StrategiesRepository {
	repo := &strategiesRepo{
		strategiesColl: strategiesColl,
	}

	return repo
}

func (repo strategiesRepo) Insert(ctx context.Context, strategy *domain.Strategy) (string, error) {
	strategyDTO, err := fromStrategyCoreToDTO(strategy)
	if err != nil {
		return "", apierrors.NewBadRequestError(err.Error())
	}
	result, err := repo.strategiesColl.InsertOne(ctx, strategyDTO)
	if err != nil {
		return "", apierrors.NewInternalServerError("error creating strategy")
	}

	strategyID, ok := result.InsertedID.(bson.ObjectID)
	if !ok {

		return "", apierrors.NewInternalServerError("error getting inserted id")
	}

	return strategyID.Hex(), nil
}

func (repo strategiesRepo) GetByID(ctx context.Context, id string) (*domain.Strategy, error) {
	mongoID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, apierrors.NewBadRequestError(fmt.Sprintf("invalid strategy id: '%s'", err))
	}

	filter := bson.M{"_id": mongoID}

	result := repo.strategiesColl.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, apierrors.NewNotFoundError(fmt.Sprintf("strategy '%s' not found", id))
		}
		return nil, apierrors.NewInternalServerError("error getting strategy by id")
	}

	var strategyDTO StrategiesDTO
	if err := result.Decode(&strategyDTO); err != nil {
		return nil, apierrors.NewInternalServerError(fmt.Sprintf("error decoding strategy '%s", id))
	}

	strategy, err := fromStrategyDTOToCore(&strategyDTO)
	if err != nil {
		return nil, err
	}

	return strategy, nil
}
