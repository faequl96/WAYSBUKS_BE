package repositories

import (
	"waysbucks_api/models"

	"gorm.io/gorm"
)

type ToppingRepository interface {
	RepoGetToppings() ([]models.Topping, error)
	RepoGetToppingByID(ID int) (models.Topping, error)
	RepoCreateTopping(topping models.Topping) (models.Topping, error)
	RepoUpdateTopping(topping models.Topping) (models.Topping, error)
	RepoDeleteTopping(topping models.Topping) (models.Topping, error)
}

func RepositoryTopping(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RepoGetToppings() ([]models.Topping, error) {
	var toppings []models.Topping
	err := r.db.Find(&toppings).Error

	return toppings, err
}

func (r *repository) RepoGetToppingByID(ID int) (models.Topping, error) {
	var topping models.Topping
	err := r.db.First(&topping, ID).Error

	return topping, err
}

func (r *repository) RepoCreateTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Create(&topping).Error

	return topping, err
}

func (r *repository) RepoUpdateTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Save(&topping).Error

	return topping, err
}

func (r *repository) RepoDeleteTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Delete(&topping).Error

	return topping, err
}
