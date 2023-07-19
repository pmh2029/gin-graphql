package handler

import (
	"gin-graphql/app/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func (h *HTTPHandler) GetAllBrands(c *gin.Context) {
	result := graphql.Do(graphql.Params{
		Schema:  h.graphql,
		Context: c,
		RequestString: `
			query { 
				brands { 
					total 
					list {
						id 
                        brand_name
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
		Data: result.Data.(map[string]interface{})["brands"],
	})
}

func (h *HTTPHandler) GetBrandByID(c *gin.Context) {
	brandID, err := strconv.Atoi(c.Param("brand_id"))
	if err != nil {
		h.SetInternalErrorResponse(c, err, map[string]interface{}{
			"strconv.Atoi": "strconv.Atoi",
		})
		return
	}

	result := graphql.Do(graphql.Params{
		Schema: h.graphql,
		VariableValues: map[string]interface{}{
			"brand_id": brandID,
		},
		Context: c,
		RequestString: `
			query($brand_id: Int!) {
                get_brand_by_id(brand_id: $brand_id) {
                    id
					brand_name
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
		Data: result.Data.(map[string]interface{})["get_brand_by_id"],
	})
}
