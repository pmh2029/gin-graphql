package output

import (
	"gin-graphql/app/model"

	"github.com/graphql-go/graphql"
)

type AuthOutput struct {
	AccessToken string
	UserInfo    model.User
}

var AuthOutputType = func(
	types map[string]*graphql.Object,
) *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "auth",
			Fields: graphql.FieldsThunk(func() graphql.Fields {
				return graphql.Fields{
					"access_token": &graphql.Field{
						Type: graphql.String,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(AuthOutput).AccessToken, nil
						},
					},
					"user_info": &graphql.Field{
						Type: types["user"],
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(AuthOutput).UserInfo, nil
						},
					},
				}
			}),
		},
	)
}
