package v1

import (
	"code-sharing-be/internal/controller/http/v1/response"
	"errors"
	"github.com/gin-gonic/gin"
)

func errorResponse(ctx *gin.Context, code int, msg string) error {
	ctx.JSON(code, response.Error{Error: msg})
	return gin.Error{Err: errors.New(msg)}
}
