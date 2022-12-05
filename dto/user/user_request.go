package userdto

type UserUpdateRequest struct {
	Name     string `json:"name" gorm:"type: varchar(255)" validate:"required"`
	Email    string `json:"email" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"password" gorm:"type: varchar(255)" validate:"required"`
	Image    string `json:"Image" gorm:"type: varchar(255)" validate:"required"`
}
