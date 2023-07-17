package handler

import (
	"fmt"
	"gin-graphql/app/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func (h *HTTPHandler) GetAllUsers(c *gin.Context) {
	result := graphql.Do(graphql.Params{
		Schema:  h.graphql,
		Context: c,
		RequestString: `
			query {
                users {
					total
					list {
						id
						username
						email
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
		Data: result.Data.(map[string]interface{})["users"],
	})
}

func (h *HTTPHandler) GetUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		h.SetInternalErrorResponse(c, err, map[string]interface{}{
			"userID": "strconv.Atoi",
		})
		return
	}

	result := graphql.Do(graphql.Params{
		Schema: h.graphql,
		VariableValues: map[string]interface{}{
			"user_id": userID,
		},
		Context: c,
		RequestString: `
			query($user_id: Int!) {
                get_user_by_id(user_id: $user_id) {
                    id
					email
                    username
					is_admin
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
		Data: result.Data.(map[string]interface{})["get_user_by_id"],
	})
}

func (h *HTTPHandler) DeleteUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		h.SetInternalErrorResponse(c, err, map[string]interface{}{
			"userID": "strconv.Atoi",
		})
		return
	}

	result := graphql.Do(graphql.Params{
		Schema: h.graphql,
		VariableValues: map[string]interface{}{
			"user_id": userID,
		},
		Context: c,
		RequestString: `
			mutation($user_id: Int!) {
                delete_user(user_id: $user_id) 
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
		Data: result.Data,
	})
}

func (h *HTTPHandler) UpdateUserByID(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		h.SetInternalErrorResponse(c, err, map[string]interface{}{
			"userID": "strconv.Atoi",
		})
		return
	}

	input := dto.UpdateUserDto{}
	err = c.ShouldBindJSON(&input)
	if err != nil {
		h.SetBadRequestErrorResponse(c, err, map[string]interface{}{
			"ShouldBindJSON": "ShouldBindJSON",
		})
		return
	}
	fmt.Println(input)

	data, err := h.GetInputAsMap(c, input)
	fmt.Println(data)
	if err != nil {
		h.SetInternalErrorResponse(c, err, map[string]interface{}{
			"GetInputAsMap": "GetInputAsMap",
		})
		return
	}

	result := graphql.Do(graphql.Params{
		Schema:     h.graphql,
		RootObject: data,
		VariableValues: map[string]interface{}{
			"user_id": userID,
		},
		Context: c,
		RequestString: `
			mutation($user_id: Int!) {
                update_user(user_id: $user_id) {
					id
					username
					email
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
		Data: result.Data.(map[string]interface{})["update_user"],
	})
}
