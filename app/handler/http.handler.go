package handler

import (
	"encoding/json"
	"gin-graphql/app/dto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"gorm.io/gorm"
)

type HTTPHandler struct {
	graphql graphql.Schema
}

func NewHTTPHandler(
	graphql graphql.Schema,
) *HTTPHandler {
	return &HTTPHandler{
		graphql: graphql,
	}
}

func (h *HTTPHandler) GetInputAsMap(c *gin.Context, input interface{}) (map[string]interface{}, error) {
	jsonBytes, err := json.Marshal(&input)
	if err != nil {
		return nil, err
	}

	output := make(map[string]interface{})
	err = json.Unmarshal(jsonBytes, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func (h *HTTPHandler) SetInternalErrorResponse(c *gin.Context, err error, msg map[string]interface{}) {
	data := &dto.BaseErrorResponse{
		Error: &dto.ErrorResponse{
			Details: msg,
			Message: err.Error(),
		},
	}

	c.JSON(http.StatusInternalServerError, data)
}

func (h *HTTPHandler) SetNotFoundErrorResponse(c *gin.Context, err error, msg map[string]interface{}) {
	data := &dto.BaseErrorResponse{
		Error: &dto.ErrorResponse{
			Details: msg,
			Message: err.Error(),
		},
	}

	c.JSON(http.StatusNotFound, data)
}

func (h *HTTPHandler) SetBadRequestErrorResponse(c *gin.Context, err error, msg map[string]interface{}) {
	data := &dto.BaseErrorResponse{
		Error: &dto.ErrorResponse{
			Details: msg,
			Message: err.Error(),
		},
	}

	c.JSON(http.StatusBadRequest, data)
}

func (h *HTTPHandler) SetGraphqlErrorResponse(c *gin.Context, err gqlerrors.FormattedError, msg map[string]interface{}) {
	if err.OriginalError().Error() == gorm.ErrRecordNotFound.Error() {
		h.SetNotFoundErrorResponse(c, err, msg)
	} else if strings.Contains(err.OriginalError().Error(), "bad request") {
		h.SetBadRequestErrorResponse(c, err, msg)
	} else {
		h.SetInternalErrorResponse(c, err, msg)
	}
}
