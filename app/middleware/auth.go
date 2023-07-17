package middleware

import (
	"gin-graphql/app/dto"
	auth "gin-graphql/pkg/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessToken string
		authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "Unauthorized",
				Details: `c.Request.Header.Get("Authorization")`,
			})
			c.Abort()
			return
		}

		token := strings.Fields(authorization)
		if len(token) > 0 && token[0] == "Bearer" {
			accessToken = token[1]
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "Unauthorized",
				Details: `strings.Fields(authorization)`,
			})
			c.Abort()
			return
		}

		if accessToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "Unauthorized",
				Details: `accessToken == ""`,
			})
			c.Abort()
			return
		}

		if !auth.VerifyJWT(accessToken) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "Unauthorized",
				Details: `!jwt.VerifyJWT(accessToken)`,
			})
			c.Abort()
			return
		}

		decodedToken, err := auth.Decode(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "Unauthorized",
				Details: `jwt.Decode(accessToken)`,
			})
			c.Abort()
			return
		}

		isAdmin, ok := decodedToken.Claims.(jwt.MapClaims)["is_admin"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "Unauthorized",
				Details: `ok := decodedToken.Claims.(jwt.MapClaims)["is_admin"]`,
			})
			c.Abort()
			return
		}

		c.Set("is_admin", isAdmin)
		c.Next()
	}
}

func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, ok := c.Get("is_admin")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Message: "Unauthorized",
				Details: `ok := c.Get("is_admin")`,
			})
			c.Abort()
			return
		}
		if !isAdmin.(bool) {
			c.AbortWithStatusJSON(http.StatusForbidden, dto.ErrorResponse{
				Message: "Forbidden",
				Details: `isAdmin.(bool)`,
			})
		}
	}
}
