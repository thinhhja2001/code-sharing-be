package repo

import (
	"code-sharing-be/internal/entity"
	"gorm.io/gorm"
)

type CodeBlockRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CodeBlockRepo {
	return &CodeBlockRepo{db: db}
}

func (repo *CodeBlockRepo) Insert(codeBlock entity.CodeBlock) (entity.CodeBlock, error) {
	result := repo.db.Create(&codeBlock)
	if result.Error != nil {
		return entity.CodeBlock{}, result.Error
	}
	return codeBlock, nil
}

func (repo *CodeBlockRepo) Get(codeBlock entity.CodeBlock) (entity.CodeBlock, error) {
	result := repo.db.First(&codeBlock)
	if result.Error != nil {
		return entity.CodeBlock{}, result.Error
	}
	return codeBlock, nil
}

func (repo *CodeBlockRepo) GetByIds(codeBlocks []entity.CodeBlock, ids []int) ([]entity.CodeBlock, error) {
	result := repo.db.Find(&codeBlocks, ids)
	if result.Error != nil {
		return []entity.CodeBlock{}, nil
	}
	return codeBlocks, nil
}
