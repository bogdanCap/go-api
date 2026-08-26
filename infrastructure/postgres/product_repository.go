package postgres

import (
	"context"
	"github.com/bogdanCap/go-api/internal/product/domain"

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
) ([]*domain.Product, error) {

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
-------------------------

	query := `
		SELECT
			id,
			name,
			email,
			status
		FROM users
	`

	var conditions []string
	var args []any
	argIndex := 1

	if filter.Name != nil {
		conditions = append(
			conditions,
			fmt.Sprintf("name ILIKE $%d", argIndex),
		)
		args = append(args, "%"+*filter.Name+"%")
		argIndex++
	}

	if filter.Email != nil {
		conditions = append(
			conditions,
			fmt.Sprintf("email ILIKE $%d", argIndex),
		)
		args = append(args, "%"+*filter.Email+"%")
		argIndex++
	}

	if filter.Status != nil {
		conditions = append(
			conditions,
			fmt.Sprintf("status = $%d", argIndex),
		)
		args = append(args, *filter.Status)
		argIndex++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += " ORDER BY id DESC"

	if filter.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, filter.Limit)
		argIndex++
	}

	if filter.Offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, filter.Offset)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query users: %w", err)
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var user User

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Status,
		); err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate users: %w", err)
	}

	return users, nil


	------------ or better to do this if need  
	query := "SELECT * FROM users WHERE name = $1"
	args := []any{*filter.Name}

	*/

	return []*domain.Product{}, nil
}
