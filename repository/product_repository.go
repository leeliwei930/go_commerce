package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/leeliwei930/go_commerce/ent"
	"github.com/leeliwei930/go_commerce/ent/product"
)

type ProductRepository struct {
	Client *ent.Client
}

func (repository *ProductRepository) PaginateProduct(c context.Context, perPage, page int) ([]*ent.Product, error) {
	return repository.Client.Product.Query().Limit(perPage).Offset(page).All(c)
}

func (repository *ProductRepository) CreateProduct(c context.Context, product *ent.Product) (*ent.Product, error) {
	productCreate := repository.Client.Product.Create()
	productCreate.SetName(product.Name).
		SetDescription(product.Description).
		SetPublishedAt(*product.PublishedAt)

	return productCreate.Save(c)
}

func (repository *ProductRepository) FindProduct(c context.Context, id uuid.UUID) (*ent.Product, error) {

	return repository.Client.Product.Query().Where(product.ID(id)).First(c)

}
