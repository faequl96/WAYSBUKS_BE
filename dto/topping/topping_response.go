package toppingdto

type ToppingResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
}

type ToppingDeleteResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title" form:"title" validate:"required"`
}
