package interfaces

import "shemsi.com/internal/models"

type DB interface {
	Connect() error
	ListProducts() ([]models.Product, error)
	GetProductBOM(productID string) (models.BOM, error)
}
