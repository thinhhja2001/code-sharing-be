package controller

import (
	"code-sharing-be/cmd/config"
	v1 "code-sharing-be/internal/controller/http/v1"
	"code-sharing-be/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(engine *gin.Engine, cfg *config.Config, c usecase.CodeBlock) {
	apiV1Group := engine.Group("/v1")
	{
		v1.NewCodeBlockRoute(apiV1Group, c)
	}
}
