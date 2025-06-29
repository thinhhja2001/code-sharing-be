package code_block

import (
	"code-sharing-be/internal/entity"
	"code-sharing-be/internal/repo"
	"code-sharing-be/internal/usecase"
	"code-sharing-be/pkg/logger"
)

type CodeBlockUseCase struct {
	repo   repo.CodeBlockRepo
	logger *logger.Logger
}

func New(repo repo.CodeBlockRepo) usecase.CodeBlock {
	return &CodeBlockUseCase{repo: repo}
}

func (receiver *CodeBlockUseCase) Create(block entity.CodeBlock) (entity.CodeBlock, error) {
	return receiver.repo.Insert(block)
}

func (receiver *CodeBlockUseCase) GetCodeBlockById(id int) (entity.CodeBlock, error) {
	codeBlock := entity.CodeBlock{Id: id}
	return receiver.repo.Get(codeBlock)
}

func (receiver *CodeBlockUseCase) GetByIds(ids []int) ([]entity.CodeBlock, error) {
	var codeBlocks []entity.CodeBlock
	return receiver.repo.GetByIds(codeBlocks, ids)
}
