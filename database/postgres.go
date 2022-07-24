package database

import (
	"context"
	"database/sql"
	"gows/models"
	"log"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (email, password, id) VALUES ($1, $2, $3)", user.Email, user.Password, user.Id)
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {

	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) InsertProduct(ctx context.Context, product *models.Product) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO products (id, product_name, created_by) VALUES ($1, $2, $3)", product.Id, product.ProductName, product.CreatedBy)
	return err
}

func (repo *PostgresRepository) GetProductById(ctx context.Context, id string) (*models.Product, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, product_name, created_by FROM products WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var product = models.Product{}
	for rows.Next() {
		if err = rows.Scan(&product.Id, &product.ProductName, &product.CreatedBy); err == nil {
			return &product, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &product, nil
}

func (repo *PostgresRepository) UpdateProduct(ctx context.Context, product *models.Product) error {
	_, err := repo.db.ExecContext(ctx, "UPDATE products SET product_name = $1 WHERE id = $2", product.ProductName, product.Id)
	return err
}

func (repo *PostgresRepository) DeleteProductById(ctx context.Context, id string) error {
	_, err := repo.db.ExecContext(ctx, "DELETE from products WHERE id = $1", id)
	return err
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}
