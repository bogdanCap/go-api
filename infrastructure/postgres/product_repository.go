package postgres

import (
	"context"
	"fmt"
	"strings"
	"github.com/bogdanCap/go-api/internal/product/domain"
	"github.com/bogdanCap/go-api/internal/product/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

// repo need to get domain object as parameter
type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(
	db *pgxpool.Pool,
) ProductRepository {
	return ProductRepository{
		db: db,
	}

}

func (r ProductRepository) Create(
	ctx context.Context,
	product *domain.Product,
) error {


	query := `
		INSERT INTO orders
		(
		 	name,
			created_at
		)
		VALUES
		($1,$2,$3,$4)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		product.Name,
		//product.Price,
		//product.Status,
		product.CreatedAt,
	)

	return err
}

func (r ProductRepository) List(
	ctx context.Context,
	filter dto.ProductListFilterDTO,
) ([]domain.Product, error) {

	query := `
		SELECT
			p.article_id,
			p.created_at,
			p.updated_at,
			ps.name,
			ps.desc
		FROM products p
		INNER JOIN product_specs ps ON ps.guid = p.spec_guid
	`

	var conditions []string
	var args []any
	argIndex := 1

	if filter.Name != nil {
		conditions = append(
			conditions,
			fmt.Sprintf("ps.name = $%d", argIndex),
		)
		args = append(args, *filter.Name)
		argIndex++
	}


	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY p.article_id DESC"

	query += fmt.Sprintf(" LIMIT $%d", argIndex)
	args = append(args, filter.Limit)
	argIndex++

	//this last argument and argIndex++ do not need
	query += fmt.Sprintf(" OFFSET $%d", argIndex)
	args = append(args, filter.Offset)

	rows, err := r.db.Query(ctx, query, args...)

	if err != nil {
		return nil, fmt.Errorf("query products error: %w", err)
	}

	defer rows.Close()

	products := make([]domain.Product, 0)

	for rows.Next() {
		var product domain.Product

		if err := rows.Scan(
			&product.ArticleID,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.Name,
			&product.Desc,
		); err != nil {
			return nil, fmt.Errorf("scan product: %w", err)
		}

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate products: %w", err)
	}

	return products, nil
}
