package repository

import (
	"context"
	"gin-graphql/app/model"
	"gin-graphql/pkg/utils"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	Create(
		ctx context.Context,
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Product, error)
	CreateWithTx(
		tx *gorm.DB,
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Product, error)
	TakeByConditionsWithScopes(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Product, error)
	FindByConditionsWithScopes(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) ([]model.Product, error)
	UpdateByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Product, error)
	UpdateByConditionsWithTx(
		tx *gorm.DB,
		conditions map[string]interface{},
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Product, error)
	DeleteByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) error
	DeleteByConditionsWithTx(
		tx *gorm.DB,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) error
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{db}
}

func (r *ProductRepository) Create(
	ctx context.Context,
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Product, error) {
	var product model.Product
	err := utils.MapToStruct(data, &product)
	if err != nil {
		return product, err
	}

	cdb := r.DB.WithContext(ctx)
	err = cdb.Scopes(scopes...).Create(&product).Error

	return product, err
}

func (r *ProductRepository) CreateWithTx(
	tx *gorm.DB,
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Product, error) {
	var product model.Product
	err := utils.MapToStruct(data, &product)
	if err != nil {
		return product, err
	}

	err = tx.Scopes(scopes...).Create(&product).Error
	return product, err
}

func (r *ProductRepository) TakeByConditionsWithScopes(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Product, error) {
	var product model.Product
	cdb := r.DB.WithContext(ctx)
	err := cdb.Scopes(scopes...).Where(conditions).Take(&product).Error

	return product, err
}

func (r *ProductRepository) FindByConditionsWithScopes(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) ([]model.Product, error) {
	var products []model.Product
	cdb := r.DB.WithContext(ctx)
	err := cdb.Scopes(scopes...).Where(conditions).Find(&products).Error

	return products, err
}

func (r *ProductRepository) UpdateByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Product, error) {
	var product model.Product
	cdb := r.DB.WithContext(ctx)
	err := cdb.Model(&model.Product{}).Scopes(scopes...).Where(conditions).Updates(data).First(&product).Error

	return product, err
}

func (r *ProductRepository) UpdateByConditionsWithTx(
	tx *gorm.DB,
	conditions map[string]interface{},
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Product, error) {
	var product model.Product
	cdb := tx.Model(&model.Product{})
	err := cdb.Scopes(scopes...).Where(conditions).Updates(data).First(&product).Error

	return product, err
}

func (r *ProductRepository) DeleteByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) error {
	cdb := r.DB.WithContext(ctx)
	return cdb.Scopes(scopes...).Where(conditions).Delete(&model.Product{}).Error
}

func (r *ProductRepository) DeleteByConditionsWithTx(
	tx *gorm.DB,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) error {
	return tx.Scopes(scopes...).Where(conditions).Delete(&model.Product{}).Error
}
