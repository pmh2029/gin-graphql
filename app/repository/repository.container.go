package repository

import "gorm.io/gorm"

type RepositoryContainer struct {
	UserRepository UserRepositoryInterface
}

func NewRepositoryContainer(db *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		UserRepository: NewUserRepository(db),
	}
}
