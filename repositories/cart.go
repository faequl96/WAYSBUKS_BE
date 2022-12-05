package repositories

import (
	"waysbucks_api/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	RepoAddCart(cart models.Cart) error

	RepoGetCart(ID int) (models.Cart, error)
	RepoDelCart(cart models.Cart) (models.Cart, error)

	// RepoGetCartByIdUser(ID int) ([]models.Cart, error)

	RepoGetCartByIdUserIsPayedFalse(ID int) ([]models.Cart, error)
	RepoUpdateCart(ID []int, IsPayed bool, TransDay string, TransTime string) error

	// RepoUpdateCart(cart models.Cart) (models.Cart, error)
	RepoGetProductCart(ID int) (models.Product, error)
	RepoGetToppingCart(ID []int) ([]models.Topping, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RepoAddCart(cart models.Cart) error {
	err := r.db.Create(&cart).Error

	return err
}
func (r *repository) RepoDelCart(cart models.Cart) (models.Cart, error) {
	err := r.db.Delete(&cart).Error

	return cart, err
}

// func (r *repository) RepoGetCartByIdUser(ID int) ([]models.Cart, error) {
// 	var cart []models.Cart
// 	err := r.db.Preload("Product").Preload("Toppings").Preload("Customer").Where("customer_id = ?", ID).Find(&cart).Error

// 	return cart, err
// }

func (r *repository) RepoGetCart(ID int) (models.Cart, error) {
	var cart models.Cart
	err := r.db.Preload("Product").Preload("Toppings").Preload("Customer").First(&cart, ID).Error

	return cart, err
}

// func (r *repository) RepoUpdateCart(cart models.Cart) (models.Cart, error) {
// 	err := r.db.Save(&cart).Error

// 	return cart, err
// }

func (r *repository) RepoUpdateCart(ID []int, IsPayed bool, TransDay string, TransTime string) error {
	err := r.db.Table("carts").Where("id IN ?", ID).Updates(map[string]interface{}{"is_payed": IsPayed, "trans_day": TransDay, "trans_time": TransTime}).Error

	return err
}

func (r *repository) RepoGetCartByIdUserIsPayedFalse(ID int) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Product").Preload("Toppings").Preload("Customer").Where(map[string]interface{}{"customer_id": ID, "is_payed": false}).Find(&cart).Error

	return cart, err
}

func (r *repository) RepoGetProductCart(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error

	return product, err
}

func (r *repository) RepoGetToppingCart(ID []int) ([]models.Topping, error) {
	var topping []models.Topping
	err := r.db.Find(&topping, ID).Error

	return topping, err
}
