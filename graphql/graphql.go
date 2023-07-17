package graphql

import (
	"gin-graphql/app/repository"
	"gin-graphql/graphql/mutation"
	"gin-graphql/graphql/output"
	"gin-graphql/graphql/query"

	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

var NewGraphqlSchema = func(
	db *gorm.DB,
) (graphql.Schema, error) {
	repositories := repository.NewRepositoryContainer(db)
	outputTypes := make(map[string]*graphql.Object)
	for _, graphqlType := range []*graphql.Object{
		output.AuthOutputType(outputTypes),
		output.UserOutputType(),
		output.ListUserOutputType(outputTypes),
	} {
		outputTypes[graphqlType.Name()] = graphqlType
	}

	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"users": query.GetAllUsersQueryType(
					outputTypes,
					repositories.UserRepository,
				),
				"get_user_by_id": query.GetUserByIDQueryType(
					outputTypes,
					repositories.UserRepository,
				),
			},
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				"signup": mutation.SignUpMutationType(
					outputTypes,
					repositories.UserRepository,
				),
				"signin": mutation.SignInMutationType(
					outputTypes,
					repositories.UserRepository,
				),
				"update_user": mutation.UpdateUserMutationType(
					outputTypes,
					repositories.UserRepository,
				),
				"delete_user": mutation.DeleteUserMutationType(
					repositories.UserRepository,
				),
			},
		}),
	})
}
