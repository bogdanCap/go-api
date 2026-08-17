package posgress

import (
	"context"
	"test/internal/domain/product"

	"github.com/jackc/pgx/v5/pgxpool"
)

// repo need to get domain object as parameter
type ProductRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(
	db *pgxpool.Pool,
) ProductRepository {

	return ProductRepository{
		db: db,
	}

}

func (r ProductRepository) Create(
	ctx context.Context,
	product *product.Product,
) error {

	query := `
		INSERT INTO orders
		(
		 	name,
			price,
			status,
			created_at
		)
		VALUES
		($1,$2,$3,$4)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		product.Name,
		product.Price,
		product.Status,
		product.CreatedAt,
	)

	return err
}

func (r ProductRepository) List(
	ctx context.Context,
	limit int,
	offset int,
) ([]*product.Product, error) {

	/* LIST return domain object
	rows, err := r.queries.ListOrders(
		ctx,
		ListOrdersParams{
			Limit: limit,
			Offset: offset,
		},
	)


	if err != nil {
		return nil, err
	}


	result := make(
		[]*product.Product,
		0,
		len(rows),
	)


	for _, row := range rows {

		result = append(
			result,
			&product.Product{
				Price: row.Price,
				Name: row.Name,
				CustomerID: row.CustomerID,
				Status: order.Status(row.Status),
				CreatedAt: row.CreatedAt,
			},
		)
	}
	*/

	return []*product.Product{}, nil
}
