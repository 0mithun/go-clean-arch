package app

import (
	"context"
	"fmt"

	strategiesHTTP "github.com/0mithun/go-clean-arch/internal/adapters/http/strategies"
	mongoAdapters "github.com/0mithun/go-clean-arch/internal/adapters/mongo"
	"github.com/0mithun/go-clean-arch/internal/config"
	strategiesUC "github.com/0mithun/go-clean-arch/internal/usecases/strategies"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Application interface {
	Run(ctx context.Context) error
}

type application struct {
	httpRouter *gin.Engine
	config     *config.Config
}

func NewApplication(ctx context.Context) (Application, error) {
	conf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	app := &application{
		httpRouter: gin.Default(),
		config:     conf,
	}

	fmt.Println(conf.Mongo.GetConnectionString())
	mongoOpts := options.Client().
		SetAppName(conf.Mongo.AppName).
		ApplyURI(conf.Mongo.GetConnectionString()).
		SetMinPoolSize(uint64(conf.Mongo.MinPoolSize)).
		SetMaxPoolSize(uint64(conf.Mongo.MaxPoolSize))

	mongoClient, err := mongo.Connect(mongoOpts)
	if err != nil {
		return nil, fmt.Errorf("mongo connect err: %w", err)
	}
	if err = mongoClient.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("mongo ping err: %w", err)
	}

	mongoDB := mongoClient.Database(conf.Mongo.Database)
	strategiesColl := mongoDB.Collection("strategies")

	strategiesRepo := mongoAdapters.NewStrategiesRepository(strategiesColl)
	strategiesService := strategiesUC.NewService(strategiesRepo)
	strategiesHandler := strategiesHTTP.NewHandlers(strategiesService)

	strategiesHTTP.RegisterRoutes(app.httpRouter, strategiesHandler)

	return app, nil
}

func (app *application) Run(ctx context.Context) error {

	return app.httpRouter.Run(fmt.Sprintf(":%s", app.config.Http.Port))
}
