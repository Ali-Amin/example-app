package services

import (
	"shemsi.com/internal/interfaces"
	"shemsi.com/internal/models"
)

type ProductsService struct {
	db interfaces.DB
}

func NewProductService(db interfaces.DB) *ProductsService {
	return &ProductsService{
		db: db,
	}
}

func (s *ProductsService) ListProducts() ([]models.Product, error) {
	products, err := s.db.ListProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductsService) GetProductBOM(productID string) (models.BOM, error) {
	bom, err := s.db.GetProductBOM(productID)
	if err != nil {
		return models.BOM{}, err
	}

	return bom, nil
}
