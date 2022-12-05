package repositories

import (
	"waysbucks_api/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	RepoCheckout(transaction models.Transaction) (models.Transaction, error)
	RepoGetCartById(ID []int) ([]models.Cart, error)
	RepoGetTransactionByIdUser(ID int) ([]models.Transaction, error)
	RepoGetTransactions() ([]models.Transaction, error)
	RepoGetTransactionByIdTrans(ID int) (models.Transaction, error)
	UpdateTransactionUser(status string, ID int) error
}

func RepoTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RepoCheckout(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

func (r *repository) RepoGetTransactionByIdTrans(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Carts").Preload("Customer").Where("id = ?", ID).Find(&transaction).Error

	return transaction, err
}

func (r *repository) RepoGetTransactionByIdUser(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Preload("Carts").Preload("Carts.Product").Preload("Carts.Toppings").Where("customer_id = ?", ID).Find(&transaction).Error

	return transaction, err
}

func (r *repository) RepoGetTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Find(&transactions).Error

	return transactions, err
}

// clause.Associations

func (r *repository) RepoGetCartById(ID []int) ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("Customer").Preload("Product").Preload("Toppings").Find(&cart, ID).Error

	return cart, err
}

// func (r *repository) RepoGetCartByIdUserIsPayedFalse(ID int) ([]models.Cart, error) {
// 	var cart []models.Cart
// 	err := r.db.Preload("Product").Preload("Toppings").Preload("Customer").Where(map[string]interface{}{"customer_id": ID, "is_payed": false}).Find(&cart).Error

// 	return cart, err
// }

func (r *repository) UpdateTransactionUser(status string, ID int) error {
	var transaction models.Transaction
	r.db.Preload("Carts.Product").First(&transaction, ID)
	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return err

	// If is different & Status is "success" decrement product quantity
	// if status != transaction.Status && status == "success" {
	// 	var product models.Product
	// 	r.db.First(&product, transaction.Product.ID)
	// 	product.Qty = product.Qty - 1
	// 	r.db.Save(&product)
	// }

}
