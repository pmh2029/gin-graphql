package mutation

import (
	"gin-graphql/app/repository"
	"gin-graphql/graphql/output"
	"gin-graphql/pkg/jwt"
	"time"

	"github.com/graphql-go/graphql"
)

var SignUpMutationType = func(
	types map[string]*graphql.Object,
	userRepo repository.UserRepositoryInterface,
) *graphql.Field {
	return &graphql.Field{
		Type: types["auth"],
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			data := p.Source.(map[string]interface{})

			user, err := userRepo.Create(p.Context, data)
			if err != nil {
				return nil, err
			}

			token, err := jwt.GenerateHS256JWT(map[string]interface{}{
				"sub":      user.ID,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
				"is_admin": user.IsAdmin,
			})
			if err != nil {
				return nil, err
			}

			return output.AuthOutput{
				AccessToken: token,
				UserInfo:    user,
			}, nil
		},
	}
}

var SignInMutationType = func(
	types map[string]*graphql.Object,
	userRepo repository.UserRepositoryInterface,
) *graphql.Field {
	return &graphql.Field{
		Type: types["auth"],
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			conditions := p.Source.(map[string]interface{})

			user, err := userRepo.TakeByConditionsWithScopes(p.Context, conditions)
			if err != nil {
				return nil, err
			}

			token, err := jwt.GenerateHS256JWT(map[string]interface{}{
				"sub":      user.ID,
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
				"is_admin": user.IsAdmin,
			})
			if err != nil {
				return nil, err
			}

			return output.AuthOutput{
				AccessToken: token,
				UserInfo:    user,
			}, nil
		},
	}
}
