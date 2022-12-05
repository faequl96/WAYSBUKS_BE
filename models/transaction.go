package models

type Transaction struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	PosCode    string `json:"pos_code"`
	Address    string `json:"address"`
	TotalPrice int    `json:"total_price"`
	Status     string `json:"status"`

	CustomerID int          `json:"customer_id"`
	Customer   UserResponse `json:"customer"`
	CartsID    []int        `json:"carts_id" gorm:"-"`
	Carts      []Cart       `json:"carts" gorm:"many2many:selected_transaction;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// type TransactionResponse struct {
// 	ID     int            `json:"id"`
// 	UserID int            `json:"userOrder_id"`
// 	User   UserResponse   `json:"userOrder"`
// 	Total  int            `json:"total"`
// 	Status string         `json:"status"`
// 	CartID int            `json:"cart_id"`
// 	Cart   []CartResponse `json:"cart"`
// }

// func (TransactionResponse) TableName() string {
// 	return "transactions"
// }
