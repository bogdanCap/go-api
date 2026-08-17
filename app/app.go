package app

import (
	"fmt"
	"net/http"
	"test/config"
	"test/infrastructure/logger"
	"test/internal/application"
	"test/internal/controller"
	/*
		"myapp/internal/config"
		"myapp/internal/database"
		"myapp/internal/externalapi"
		"myapp/internal/kafka"
		"myapp/internal/product"
		"myapp/internal/redis"
		"myapp/internal/router"*/

	"go.uber.org/zap"
)

type App struct {
	router http.Handler
	config config.Config
	logger *zap.Logger
}

func New() (*App, error) {

	cfg := config.Load()

	log, err := logger.New(cfg.App.Environment)
	if err != nil {
		return nil, err
	}

	/*
		db, err := database.New(cfg.Database)
		if err != nil {
			return nil, err
		}

		redisClient := redis.New(cfg.Redis)

		kafkaProducer := kafka.NewProducer(cfg.Kafka)

		externalClient := externalapi.New(cfg.ExternalAPI)

		productRepo := product.NewRepository(db)
	*/
	productService := application.NewProductService(
		/*productRepo,
		redisClient,
		kafkaProducer,
		externalClient,*/
		log,
	)

	productController := controller.NewProductController(productService)

	r := NewRouter(productController)

	return &App{
		router: r,
		config: cfg,
		logger: log,
	}, nil
}

func (a *App) Run() {
	serverPort := a.config.Server.Port
	fmt.Println("Server started on:", serverPort)

	http.ListenAndServe(":"+serverPort, a.router)
}
