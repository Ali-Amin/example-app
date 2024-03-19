package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"shemsi.com/internal/interfaces"
	"shemsi.com/internal/models"
)

const (
	ConnectionSTR         = "postgres://postgres:password@localhost:5432/shemsi"
	GetProductsQuery      = "SELECT * FROM product"
	GetProductComposition = "SELECT * FROM composition WHERE product_id=$1"
)

type PostgresDB struct {
	conn *pgx.Conn
}

func NewPostgresDB() interfaces.DB {
	return &PostgresDB{}
}

func (d *PostgresDB) Connect() error {
	conn, err := pgx.Connect(context.Background(), ConnectionSTR)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return err
	}
	d.conn = conn
	return nil
}

func (d *PostgresDB) ListProducts() ([]models.Product, error) {
	var products []models.Product
	rows, err := d.conn.Query(context.Background(), GetProductsQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var name, id string
		rows.Scan(&name, &id)
		products = append(products, models.Product{Name: name, ID: id})
	}

	fmt.Println(products)
	return products, nil
}

func (d *PostgresDB) GetProductBOM(productID string) (models.BOM, error) {
	var bom []models.Material
	rows, err := d.conn.Query(context.Background(), GetProductComposition, productID)
	if err != nil {
		return models.BOM{}, err
	}

	for rows.Next() {
		var productID,
			materialName,
			materialPricePerUnit,
			materialUnit,
			materialQuantity,
			currency,
			manufactureDate string

		rows.Scan(&productID, &materialName, &materialPricePerUnit, &materialUnit, &materialQuantity, &currency, &manufactureDate)

		material := models.Material{
			Name:            materialName,
			PricePerUnit:    materialPricePerUnit,
			Unit:            materialUnit,
			Quantity:        materialQuantity,
			Currency:        currency,
			ManufactureDate: manufactureDate,
		}
		bom = append(bom, material)
	}

	return models.BOM{Materials: bom}, nil
}
