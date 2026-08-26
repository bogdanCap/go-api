package app

import (
	"fmt"
	"net/http"
	"github.com/bogdanCap/go-api/config"
	"github.com/bogdanCap/go-api/infrastructure/logger"
	//"github.com/bogdanCap/go-api/internal/application"
	productService "github.com/bogdanCap/go-api/internal/product/application/service"
	productController "github.com/bogdanCap/go-api/internal/product/controller"
	"context"
	"github.com/bogdanCap/go-api/infrastructure/postgres"
	"time"
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

	ctx := context.Background()
	cfg := config.Load()
	

	log, logErr := logger.New(cfg.App.Environment)

	if logErr != nil {
		return nil, logErr
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
	
	db, dbErr := postgres.NewPool(ctx, postgres.Config{
		Host: cfg.Database.Host,
		Port: cfg.Database.Port,
		User: cfg.Database.User,
		Password: cfg.Database.Password,
		Database: cfg.Database.Name,
		MaxConn:        20,
		MinConn:        5,
		MaxConnLifetime: time.Hour,
		MaxConnIdleTime: 30 * time.Minute,
	})

	if dbErr != nil {
		dbErrMsg := fmt.Sprintf("database connection failed: %v", dbErr)
		log.Fatal(
			dbErrMsg,
		)
	}

	defer db.Close()

	productRepo := postgres.NewProductRepository(db)



	productService := productService.New(
		productRepo,
		log,
	)

	productController := productController.NewProductController(productService)

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
