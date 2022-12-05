package repositories

import (
	"waysbucks_api/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	RepoGetProducts() ([]models.Product, error)
	RepoGetProductByID(ID int) (models.Product, error)
	RepoCreateProduct(product models.Product) (models.Product, error)
	RepoUpdateProduct(product models.Product) (models.Product, error)
	RepoDeleteProduct(product models.Product) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RepoGetProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error

	return products, err
}

func (r *repository) RepoGetProductByID(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error

	return product, err
}

func (r *repository) RepoCreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

func (r *repository) RepoUpdateProduct(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error

	return product, err
}

func (r *repository) RepoDeleteProduct(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}
