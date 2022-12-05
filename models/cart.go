package models

type Cart struct {
	ID         int          `json:"id" gorm:"primary_key:auto_increment"`
	CustomerID int          `json:"customer_id"`
	Customer   UserResponse `json:"customer" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	ProductID int             `json:"product_id"`
	Product   ProductResponse `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	ToppingsID []int     `json:"toppings_id" form:"toppings_id" gorm:"-"`
	Toppings   []Topping `json:"toppings" gorm:"many2many:selected_toppings;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price      int       `json:"price" gorm:"type: int"`
	TransDay   string    `json:"trans_day"`
	TransTime  string    `json:"trans_time"`
	IsPayed    bool      `json:"is_payed"`
}

type CartResponse struct {
	ID       int             `json:"id"`
	Product  ProductResponse `json:"product"`
	Toppings []Topping       `json:"toppings"`
	Price    int             `json:"price"`
}

func (CartResponse) TableName() string {
	return "carts"
}
