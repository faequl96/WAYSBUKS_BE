package repositories

import (
	"waysbucks_api/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	RepoRegister(user models.User) (models.User, error)
	RepoLogin(email string) (models.User, error)
	RepoGetuser(ID int) (models.User, error)
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RepoRegister(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) RepoLogin(email string) (models.User, error) {
	var user models.User
	err := r.db.First(&user, "email=?", email).Error

	return user, err
}

func (r *repository) RepoGetuser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
