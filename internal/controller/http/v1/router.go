package v1

import (
	"code-sharing-be/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewCodeBlockRoute(apiV1Group *gin.RouterGroup, c usecase.CodeBlock) {
	r := &V1{c: c}
	codeBlockGroup := apiV1Group.Group("/code-blocks")
	{
		codeBlockGroup.POST("", r.createCodeBlock)
		codeBlockGroup.GET("/:id", r.getCodeBlockById)
		codeBlockGroup.GET("", r.getByIds)
	}
}
