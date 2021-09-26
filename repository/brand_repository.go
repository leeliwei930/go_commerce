package repository

import (
	"github.com/leeliwei930/go_commerce/ent"
)

type BrandRepository struct {
	Client *ent.Client
}

// func (repository *BrandRepository) PaginateBrand(c context.Context, option *PaginatorOption) (brands []*ent.Brand, paginator *PaginatorPayload, brandsErr error) {
// 	brandQuery := repository.Client.Brand.Query()

// }
