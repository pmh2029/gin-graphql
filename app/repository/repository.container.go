package repository

import "gorm.io/gorm"

type RepositoryContainer struct {
	UserRepository        UserRepositoryInterface
	BrandRepository       BrandRepositoryInterface
	ProductRepository     ProductRepositoryInterface
	OderRepository        OrderRepositoryInterface
	OderProductRepository OrderProductRepositoryInterface
}

func NewRepositoryContainer(db *gorm.DB) RepositoryContainer {
	return RepositoryContainer{
		UserRepository:        NewUserRepository(db),
		BrandRepository:       NewBrandRepository(db),
		ProductRepository:     NewProductRepository(db),
		OderRepository:        NewOrderRepository(db),
		OderProductRepository: NewOrderProductRepository(db),
	}
}
