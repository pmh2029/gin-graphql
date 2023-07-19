package query

import (
	"gin-graphql/app/repository"
	"gin-graphql/graphql/output"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

var GetAllBrandsQueryType = func(
	types map[string]*graphql.Object,
	db *gorm.DB,
	brandRepo repository.BrandRepositoryInterface,
) *graphql.Field {
	return &graphql.Field{
		Type: types["brand_list"],
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			brands, err := brandRepo.FindByConditionsWithScopes(p.Context, nil, func(db *gorm.DB) *gorm.DB {
				return db.Order("id ASC")
			})
			if err != nil {
				return nil, err
			}

			return output.ListBrand{
				Total:     len(brands),
				ListBrand: brands,
			}, nil
		},
	}
}

var GetBrandByIDQueryType = func(
	types map[string]*graphql.Object,
	brandRepo repository.BrandRepositoryInterface,
) *graphql.Field {
	return &graphql.Field{
		Type: types["brand"],
		Args: graphql.FieldConfigArgument{
			"brand_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			brand, err := brandRepo.TakeByConditionsWithScopes(p.Context, map[string]interface{}{
				"id": p.Args["brand_id"].(int),
			})
			if err != nil {
				return nil, err
			}

			return brand, nil
		},
	}
}
