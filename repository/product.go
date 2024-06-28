package repository

import (
	"errors"
	"tpm_7_HendriHeryanto/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepo struct {
	DB *gorm.DB
}

func (p *ProductRepo) GetProduct() ([]*models.Product, error) {
	products := []*models.Product{}
	err := p.DB.Debug().Find(&products).Error
	return products, err
}

func (p *ProductRepo) GetProductById(id uint64) (*models.Product, error) {
	product := &models.Product{}
	result := p.DB.Debug().Where("id = ?", id).Find(&product)
	err := result.Error
	if result.RowsAffected < 1 {
		err = errors.New("product tidak ditemukan")
	}
	return product, err
}

func (p *ProductRepo) CreateProduct(product *models.Product) (*models.Product, error) {
	err := p.DB.Debug().Create(&product).Error
	return product, err
}

func (p *ProductRepo) UpdateProduct(id uint64, product *models.Product) (*models.Product, error) {
	result := p.DB.Debug().
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&models.Product{
			Name:  product.Name,
			Price: product.Price,
		}).Scan(&product)

	err := result.Error
	if result.RowsAffected < 1 {
		err = errors.New("product tidak ditemukan")
	}
	return product, err
}

func (p *ProductRepo) DeleteProduct(id uint64) error {
	result := p.DB.Table("products").
		Where("id = ?", id).
		Delete(&models.Product{})

	err := result.Error

	if result.RowsAffected < 1 {
		err = errors.New("product tidak ditemukan")
	}
	return err
}
