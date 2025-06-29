package v1

import (
	"code-sharing-be/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func (v1 *V1) createCodeBlock(ctx *gin.Context) error {
	var codeBlock entity.CodeBlock
	if err := ctx.BindJSON(&codeBlock); err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid code block data")
	}
	codeBlock, err := v1.c.Create(codeBlock)
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "cannot create code block")

	}
	ctx.JSON(http.StatusOK, codeBlock)
	return nil
}

func (v1 *V1) getCodeBlockById(ctx *gin.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, "invalid id")
	}
	codeBlock, err := v1.c.GetCodeBlockById(id)
	if err != nil {
		return errorResponse(ctx, http.StatusNotFound, err.Error())
	}
	ctx.JSON(http.StatusOK, codeBlock)
	return nil
}

func (v1 *V1) getByIds(ctx *gin.Context) error {
	var ids []int
	s := strings.Split(ctx.Query("ids"), ",")
	for i := range s {
		id, err := strconv.Atoi(s[i])
		if err != nil {
			return errorResponse(ctx, http.StatusInternalServerError, "invalid id")
		}
		ids = append(ids, id)
	}

	codeBlocks, err := v1.c.GetByIds(ids)
	if err != nil {
		return errorResponse(ctx, http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, codeBlocks)
	return nil
}
