package repository

import "kawanishi/first_boiler/models"

type IUserRepositories interface {
	GetUser(userID int) (*models.MUser, error)
}
