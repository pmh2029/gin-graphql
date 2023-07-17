package query

import (
	"gin-graphql/app/repository"
	"gin-graphql/graphql/output"

	"github.com/graphql-go/graphql"
)

var GetAllUsersQueryType = func(
	types map[string]*graphql.Object,
	userRepo repository.UserRepositoryInterface,
) *graphql.Field {
	return &graphql.Field{
		Type: types["user_list"],
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			users, err := userRepo.FindByConditionsWithScopes(p.Context, map[string]interface{}{
				"is_admin": false,
			})
			if err != nil {
				return nil, err
			}

			return output.ListUser{
				Total:    len(users),
				ListUser: users,
			}, nil
		},
	}
}

var GetUserByIDQueryType = func(
	types map[string]*graphql.Object,
	userRepo repository.UserRepositoryInterface,
) *graphql.Field {
	return &graphql.Field{
		Type: types["user"],
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return userRepo.TakeByConditionsWithScopes(p.Context, map[string]interface{}{
				"id":       p.Args["user_id"].(int),
				"is_admin": false,
			})
		},
	}
}
