package repository

import (
	"context"
	"gows/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertProduct(ctx context.Context, product *models.Product) error
	GetProductById(ctx context.Context, id string) (*models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProductById(ctx context.Context, id string) error
	ListProduct(ctx context.Context, page uint64) ([]*models.Product, error)
	Close() error
}

// Injeccion de dependencias para poder ser compatible
// con diferentes bases de datos

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func InsertProduct(ctx context.Context, product *models.Product) error {
	return implementation.InsertProduct(ctx, product)
}

func GetProductById(ctx context.Context, id string) (*models.Product, error) {
	return implementation.GetProductById(ctx, id)
}

func UpdateProduct(ctx context.Context, product *models.Product) error {
	return implementation.UpdateProduct(ctx, product)
}

func DeleteProductById(ctx context.Context, id string) error {
	return implementation.DeleteProductById(ctx, id)
}

func ListProduct(ctx context.Context, page uint64) ([]*models.Product, error) {
	return implementation.ListProduct(ctx, page)
}

func Close() error {
	return implementation.Close()
}
