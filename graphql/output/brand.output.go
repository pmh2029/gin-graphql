package output

import (
	"gin-graphql/app/model"

	"github.com/graphql-go/graphql"
)

var BrandOutputType = func() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "brand",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.Int,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Brand).ID, nil
					},
				},
				"brand_name": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(model.Brand).BrandName, nil
					},
				},
			}
		}),
	})
}

type ListBrand struct {
	Total     int
	ListBrand []model.Brand
}
var ListBrandOutputType = func(
	types map[string]*graphql.Object,
) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "brand_list",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"total": &graphql.Field{
					Type: graphql.Int,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(ListBrand).Total, nil
					},
				},
				"list": &graphql.Field{
					Type: graphql.NewList(types["brand"]),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return p.Source.(ListBrand).ListBrand, nil
					},
				},
			}
		}),
	})
}
