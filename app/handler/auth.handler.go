package handler

import (
	"gin-graphql/app/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func (h *HTTPHandler) SignUp(c *gin.Context) {
	input := dto.SignUpDto{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		h.SetBadRequestErrorResponse(c, err, map[string]interface{}{
			"ShouldBindJSON": "ShouldBindJSON",
		})
		return
	}

	data, err := h.GetInputAsMap(c, input)
	if err != nil {
		h.SetInternalErrorResponse(c, err, map[string]interface{}{
			"GetInputAsMap": "GetInputAsMap",
		})
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:     h.graphql,
		RootObject: data,
		Context:    c,
		RequestString: `
			mutation {
				signup {
					access_token
					user_info {
						id
						email
						username
						is_admin
					}
				}
			}
		`,
	})

	if result.HasErrors() {
		h.SetGraphqlErrorResponse(c, result.Errors[0], map[string]interface{}{
			"result.HasErrors": "result.HasErrors",
		})
		return
	}

	c.JSON(http.StatusOK, dto.BaseSuccessResponse{
		Data: result.Data.(map[string]interface{})["signup"],
	})
}

func (h *HTTPHandler) SignIn(c *gin.Context) {
	input := dto.SignInDto{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		h.SetBadRequestErrorResponse(c, err, map[string]interface{}{
			"ShouldBindJSON": "ShouldBindJSON",
		})
		return
	}

	data, err := h.GetInputAsMap(c, input)
	if err != nil {
		h.SetInternalErrorResponse(c, err, map[string]interface{}{
			"GetInputAsMap": "GetInputAsMap",
		})
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:     h.graphql,
		RootObject: data,
		Context:    c,
		RequestString: `
			mutation {
				signin {
					access_token
					user_info {
						id
						email
						username
						is_admin
					}
				}
			}
		`,
	})

	if result.HasErrors() {
		h.SetGraphqlErrorResponse(c, result.Errors[0], map[string]interface{}{
			"result.HasErrors": "result.HasErrors",
		})
		return
	}

	c.JSON(http.StatusOK, dto.BaseSuccessResponse{
		Data: result.Data.(map[string]interface{})["signin"],
	})
}
