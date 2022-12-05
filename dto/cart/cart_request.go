package cartdto

type AddCartRequest struct {
	// CustomerID int   `json:"customer_id" form:"customer_id"`
	ProductID  int   `json:"product_id" form:"product_id"`
	ToppingsID []int `json:"toppings_id" form:"toppings_id"`
	Price      int   `json:"price" form:"price"`
}

type UpdateCartRequest struct {
	ID        []int  `json:"id" form:"id"`
	TransDay  string `json:"trans_day" form:"trans_day"`
	TransTime string `json:"trans_time" form:"trans_time"`
	IsPayed   bool   `json:"is_payed" form:"is_payed"`
}
