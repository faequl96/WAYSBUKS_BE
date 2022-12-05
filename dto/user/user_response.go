package userdto

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name" form:"name" validate:"required"`
	Email string `json:"email" form:"email" validate:"required"`
	Image string `json:"image" form:"image" validate:"required"`
}

type UserUpdateResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
}

type UserDeleteResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name" form:"name" validate:"required"`
}
