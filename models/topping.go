package models

type Topping struct {
	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
	Price int    `json:"price" form:"price" gorm:"type: int"`
	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
	// AdminID int          `json:"admin_id"`
	// Admin   UserResponse `json:"admin"`
}

// type ToppingSelected struct {
// 	ID    int    `json:"id" gorm:"primary_key:auto_increment"`
// 	Title string `json:"title" form:"title" gorm:"type: varchar(255)"`
// 	Price int    `json:"price" form:"price" gorm:"type: int"`
// 	Image string `json:"image" form:"image" gorm:"type: varchar(255)"`
// }

// func (ToppingSelected) TableName() string {
// 	return "toppings"
// }
