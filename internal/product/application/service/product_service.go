package service

import (
	"context"
	"errors"
	"github.com/bogdanCap/go-api/internal/product/domain"
	"github.com/bogdanCap/go-api/internal/product/application/port"
	"github.com/bogdanCap/go-api/internal/product/dto"
	//"time"
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
	productRepo port.IProductRepository
	logger *zap.Logger
}

func New(
	/*
	   repo *Repository,
	   redis *redis.Client,
	   kafka *kafka.Producer,
	   external *externalapi.Client,
	*/
	productRepo port.IProductRepository,
	logger *zap.Logger,
) ProductService {

	return ProductService{
		/*repo: repo,
		redis: redis,
		kafka: kafka,
		external: external,*/
		productRepo: productRepo,
		logger: logger,
	}
}

func (s ProductService) List(ctx context.Context, filter dto.ProductListFilterDTO) ([]domain.Product, error) {
	s.logger.Info(
		"get products request",
		zap.String("method", "some method"),
		zap.String("path", "products"),
	)


	return s.productRepo.List(ctx, filter)
	//return domain.Product{"888", "test", 444, 3, time.Now()}, nil //s.repo.FindAll(ctx)
}

// TODO split logic into separate service
func (s ProductService) Create(
	ctx context.Context,
	product *domain.Product,
) error {

	// business validation

	/*
	if product.Price <= 0 {
		return errors.New(
			"invalid price",
		)
	}*/

	/* TODO in future add repository
	return s.repository.Create(
		ctx,
		product,
	)*/

	return errors.New("TODO in future")
}
