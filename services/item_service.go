package services

import (
	"gin-tutorial-udemy/models"
	"gin-tutorial-udemy/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

// repositoryの結果をそのまま返す
func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}