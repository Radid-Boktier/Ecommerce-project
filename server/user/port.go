package user

import (
	"ecommerce-server/domain"
	userHandler "ecommerce-server/rest/handlers/user"
)

type Service interface {
	userHandler.Service // embedding
}
type UserRepo interface {
	Create(user domain.User) (*domain.User,error)
	// Get(userID int) (*User,error)
	Find(email, pass string) (*domain.User, error)
	// List() ([]*User,error)
	// Delete(userID int) error
	// Update(user User) (*User,error)
}