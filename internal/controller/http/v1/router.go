package v1

import (
	"code-sharing-be/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewCodeBlockRoute(apiV1Group *gin.RouterGroup, c usecase.CodeBlock) {
	r := &V1{c: c}
	codeBlockGroup := apiV1Group.Group("/code-blocks")
	{
		codeBlockGroup.PUT("", func(context *gin.Context) {
			for i := 0; i < 5; i++ {
				fmt.Println(i)
			}
		})
		codeBlockGroup.POST("", func(context *gin.Context) {
			_ = r.createCodeBlock(context)
		})
		codeBlockGroup.GET("/:id", func(context *gin.Context) {
			_ = r.getCodeBlockById(context)
		})
		codeBlockGroup.GET("", func(context *gin.Context) {
			_ = r.getByIds(context)
		})
	}
}
