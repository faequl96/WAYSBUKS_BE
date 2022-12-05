package repositories

import (
	"waysbucks_api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	RepoGetUsers() ([]models.User, error)
	RepoGetUserByID(ID int) (models.User, error)
	RepoUpdateUser(user models.User) (models.User, error)
	RepoDeleteUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RepoGetUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error

	return users, err
}

func (r *repository) RepoGetUserByID(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}

func (r *repository) RepoUpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) RepoDeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
