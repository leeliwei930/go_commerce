package repository

import (
	"context"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/leeliwei930/go_commerce/ent"
	"github.com/leeliwei930/go_commerce/ent/product"
)

type ProductRepository struct {
	Client *ent.Client
}

func (repository *ProductRepository) PaginateProduct(c context.Context, option *PaginatorOption) (products []*ent.Product, paginator *PaginatorPayload, productsErr error) {
	productQuery := repository.Client.Product.Query()
	if !option.WithTrashed {
		productQuery.Where(product.DeletedAtNotNil())
	}
	resultsCount, _ := productQuery.Count(c)

	paginator = &PaginatorPayload{
		Total:       resultsCount,
		PerPage:     option.PerPage,
		CurrentPage: option.CurrentPage,
		LastPage:    int(math.Ceil(float64(resultsCount) / float64(option.PerPage))),
	}

	products, productsErr = productQuery.Limit(option.PerPage).Offset(option.CurrentPage).All(c)
	return
}

func (repository *ProductRepository) CreateProduct(c context.Context, product *ent.Product) (*ent.Product, error) {
	productCreate := repository.Client.Product.Create()
	return productCreate.SetName(product.Name).
		SetDescription(product.Description).
		SetPublishedAt(*product.PublishedAt).
		Save(c)

}

func (repository *ProductRepository) FindProduct(c context.Context, id uuid.UUID) (*ent.Product, error) {
	return repository.Client.Product.Query().Where(product.IDEQ(id)).First(c)
}

func (repository *ProductRepository) DeleteProduct(c context.Context, id uuid.UUID) (*ent.Product, error) {
	return repository.Client.Product.UpdateOneID(id).SetDeletedAt(time.Now()).Save(c)
}

func (repository *ProductRepository) RestoreProduct(c context.Context, id uuid.UUID) (*ent.Product, error) {
	updatedRow, updatedErr := repository.Client.
		Product.
		Update().
		Where(
			product.And(
				product.DeletedAtLTE(time.Now()),
				product.IDEQ(id),
			),
		).
		ClearDeletedAt().
		Save(c)

	if updatedRow == 1 {
		return repository.FindProduct(c, id)
	}

	return nil, updatedErr
}

func (repository *ProductRepository) DestroyProduct(c context.Context, id uuid.UUID) (*ent.Product, error) {
	deletedProduct, deletedErr := repository.Client.Product.Delete().
		Where(
			product.And(
				product.DeletedAtLTE(time.Now()),
				product.IDEQ(id),
			),
		).
		Exec(c)

	if deletedProduct == 1 {
		return &ent.Product{
			ID: id,
		}, nil
	}

	return nil, deletedErr
}
