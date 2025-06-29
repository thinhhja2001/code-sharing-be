package v1

import (
	"code-sharing-be/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (v1 *V1) createCodeBlock(ctx *gin.Context) {
	var codeBlock entity.CodeBlock
	if err := ctx.BindJSON(&codeBlock); err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid code block data")
		return
	}
	codeBlock, err := v1.c.Create(codeBlock)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, "cannot create code block")
		return
	}
	ctx.JSON(http.StatusOK, codeBlock)
}

func (v1 *V1) getCodeBlockById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, "invalid id")
		return
	}
	codeBlock, err := v1.c.GetCodeBlockById(id)
	if err != nil {
		errorResponse(ctx, http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, codeBlock)
}

func (v1 *V1) getByIds(ctx *gin.Context) {
	var ids []int
	s := strings.Split(ctx.Query("ids"), ",")
	for i := range s {
		id, err := strconv.Atoi(s[i])
		if err != nil {
			errorResponse(ctx, http.StatusInternalServerError, "invalid id")
			return
		}
		ids = append(ids, id)
	}

	codeBlocks, err := v1.c.GetByIds(ids)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, codeBlocks)
}
