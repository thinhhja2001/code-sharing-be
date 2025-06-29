package usecase

import "code-sharing-be/internal/entity"

type (
	CodeBlock interface {
		GetCodeBlockById(id int) (entity.CodeBlock, error)
		Create(block entity.CodeBlock) (entity.CodeBlock, error)
		GetByIds(ids []int) ([]entity.CodeBlock, error)
	}
)
