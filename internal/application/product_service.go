package application

import (
	"context"
	"errors"
	"test/internal/domain/product"
	"time"
	/*"myapp/internal/externalapi"
	"myapp/internal/kafka"
	"myapp/internal/redis"*/

	"go.uber.org/zap"
)

type ProductService struct {
	/*repo     *Repository
	redis    *redis.Client
	kafka    *kafka.Producer
	external *externalapi.Client*/
	logger *zap.Logger
}

func NewProductService(
	/*
	   repo *Repository,
	   redis *redis.Client,
	   kafka *kafka.Producer,
	   external *externalapi.Client,
	*/
	logger *zap.Logger,
) ProductService {

	return ProductService{
		/*repo: repo,
		redis: redis,
		kafka: kafka,
		external: external,*/
		logger: logger,
	}
}

func (s ProductService) List(ctx context.Context) (product.Product, error) {
	s.logger.Info(
		"get products request",
		zap.String("method", "some method"),
		zap.String("path", "products"),
	)

	return product.Product{"555", "test", 444, 3, time.Now()}, nil //s.repo.FindAll(ctx)
}

// TODO split logic into separate service
func (s ProductService) Create(
	ctx context.Context,
	product *product.Product,
) error {

	// business validation

	if product.Price <= 0 {
		return errors.New(
			"invalid price",
		)
	}

	/* TODO in future add repository
	return s.repository.Create(
		ctx,
		product,
	)*/

	return errors.New("TODO in future")
}
