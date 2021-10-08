package repository

import (
	"context"

	"github.com/leeliwei930/go_commerce/ent"
	"github.com/leeliwei930/go_commerce/inputs"
)

type BrandRepository struct {
	Client *ent.Client
}

// func (repository *BrandRepository) PaginateBrand(c context.Context, option *PaginatorOption) (brands []*ent.Brand, paginator *PaginatorPayload, brandsErr error) {
// 	brandQuery := repository.Client.Brand.Query()

// }

func (repository *BrandRepository) CreateBrand(c context.Context, brandInput *inputs.BrandRequest) (brand *ent.Brand, createError error) {
	return repository.Client.Brand.Create().SetName(brand.Name).Save(c)
}
