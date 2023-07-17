package output

import (
	"gin-graphql/app/model"

	"github.com/graphql-go/graphql"
)

var UserOutputType = func() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "user",
			Fields: graphql.FieldsThunk(func() graphql.Fields {
				return graphql.Fields{
					"id": &graphql.Field{
						Type: graphql.Int,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(model.User).ID, nil
						},
					},
					"username": &graphql.Field{
						Type: graphql.String,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(model.User).Username, nil
						},
					},
					"email": &graphql.Field{
						Type: graphql.String,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(model.User).Email, nil
						},
					},
					"is_admin": &graphql.Field{
						Type: graphql.Boolean,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(model.User).IsAdmin, nil
						},
					},
				}
			}),
		},
	)
}


type ListUser struct {
	Total    int
	ListUser []model.User
}
var ListUserOutputType = func(
	types map[string]*graphql.Object,
) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "user_list",
			Fields: graphql.FieldsThunk(func() graphql.Fields {
				return graphql.Fields{
					"total": &graphql.Field{
						Type: graphql.Int,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(ListUser).Total, nil
						},
					},
					"list": &graphql.Field{
						Type: graphql.NewList(types["user"]),
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(ListUser).ListUser, nil
						},
					},
				}
			}),
		},
	)
}
