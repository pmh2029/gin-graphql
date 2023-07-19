package router

import (
	"gin-graphql/app/dto"
	"gin-graphql/app/handler"
	"gin-graphql/app/middleware"
	"gin-graphql/graphql"
	"gin-graphql/pkg/validator"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Router is application struct
type Router struct {
	Engine *gin.Engine
	DBCon  *gorm.DB
	Logger *logrus.Logger
}

// InitializeRouter initializes Engine and middleware
func (r *Router) InitializeRouter(logger *logrus.Logger) {
	r.Engine.Use(gin.Logger())
	r.Engine.Use(gin.Recovery())
	r.Engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Set-Cookie"},
		AllowWebSockets:  true,
		AllowFiles:       true,
	}))
	r.Logger = logger
}

func (r *Router) SetupHandler() {
	graphql, _ := graphql.NewGraphqlSchema(r.DBCon)

	httpHandler := handler.NewHTTPHandler(graphql)
	_ = validator.New()

	r.Engine.GET("/", func(c *gin.Context) {
		data := dto.BaseSuccessResponse{
			Data: gin.H{"message": "ok"},
		}
		c.JSON(http.StatusOK, data)
	})

	router := r.Engine.Group("/api")
	{
		authRouter := router.Group("/auth")
		{
			authRouter.POST("/signup", httpHandler.SignUp)
			authRouter.POST("/signin", httpHandler.SignIn)
		}
	}

	routerPermission := r.Engine.Group("/api")
	routerPermission.Use(middleware.CheckAuthentication())
	{
		userRouter := routerPermission.Group("/users")
		{
			userRouter.GET("/", httpHandler.GetAllUsers)
			userRouter.GET("/:user_id", httpHandler.GetUserByID)
			userRouter.DELETE("/:user_id", httpHandler.DeleteUserByID)
			userRouter.PATCH("/:user_id", httpHandler.UpdateUserByID)
		}

		brandRouter := routerPermission.Group("/brands")
		{
			brandRouter.GET("/", httpHandler.GetAllBrands)
			brandRouter.GET("/:brand_id", httpHandler.GetBrandByID)
		}
	}
}
