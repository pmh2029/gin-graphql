package mutation

import (
	"errors"
	"gin-graphql/app/repository"
	"gin-graphql/pkg/utils"
	"net/http"

	"github.com/graphql-go/graphql"
)

var UpdateUserMutationType = func(
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
			data, ok := p.Source.(map[string]interface{})
			if !ok {
				return nil, utils.NewCustomError(http.StatusBadRequest,
					map[string]interface{}{
						"p.Source.(map[string]interface{})": "p.Source.(map[string]interface{})",
					}, errors.New("bad request: p.Source.(map[string]interface{})"))
			}

			attributeToUpdate := map[string]interface{}{}
			for k, v := range data {
				if v != nil {
					attributeToUpdate[k] = v
				}
			}

			user, err := userRepo.UpdateByConditions(
				p.Context,
				map[string]interface{}{
					"id": p.Args["user_id"].(int),
				},
				attributeToUpdate,
			)
			if err != nil {
				return nil, err
			}

			return user, nil
		},
	}
}

var DeleteUserMutationType = func(
	userRepo repository.UserRepositoryInterface,
) *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"user_id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			err := userRepo.DeleteByConditions(p.Context, map[string]interface{}{
				"id": p.Args["user_id"].(int),
			})
			if err != nil {
				return nil, err
			}

			return nil, nil
		},
	}
}
