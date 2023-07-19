package repository

import (
	"context"
	"gin-graphql/app/model"
	"gin-graphql/pkg/utils"

	"gorm.io/gorm"
)

type OrderRepositoryInterface interface {
	Create(
		ctx context.Context,
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Order, error)
	CreateWithTx(
		tx *gorm.DB,
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Order, error)
	TakeByConditionsWithScopes(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Order, error)
	FindByConditionsWithScopes(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) ([]model.Order, error)
	UpdateByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Order, error)
	UpdateByConditionsWithTx(
		tx *gorm.DB,
		conditions map[string]interface{},
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Order, error)
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

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepositoryInterface {
	return &OrderRepository{db}
}

func (r *OrderRepository) Create(
	ctx context.Context,
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Order, error) {
	var order model.Order
	err := utils.MapToStruct(data, &order)
	if err != nil {
		return order, err
	}

	cdb := r.DB.WithContext(ctx)
	err = cdb.Scopes(scopes...).Create(&order).Error

	return order, err
}

func (r *OrderRepository) CreateWithTx(
	tx *gorm.DB,
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Order, error) {
	var order model.Order
	err := utils.MapToStruct(data, &order)
	if err != nil {
		return order, err
	}

	err = tx.Scopes(scopes...).Create(&order).Error
	return order, err
}

func (r *OrderRepository) TakeByConditionsWithScopes(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Order, error) {
	var order model.Order
	cdb := r.DB.WithContext(ctx)
	err := cdb.Scopes(scopes...).Where(conditions).Take(&order).Error

	return order, err
}

func (r *OrderRepository) FindByConditionsWithScopes(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) ([]model.Order, error) {
	var orders []model.Order
	cdb := r.DB.WithContext(ctx)
	err := cdb.Scopes(scopes...).Where(conditions).Find(&orders).Error

	return orders, err
}

func (r *OrderRepository) UpdateByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Order, error) {
	var order model.Order
	cdb := r.DB.WithContext(ctx)
	err := cdb.Model(&model.Order{}).Scopes(scopes...).Where(conditions).Updates(data).First(&order).Error

	return order, err
}

func (r *OrderRepository) UpdateByConditionsWithTx(
	tx *gorm.DB,
	conditions map[string]interface{},
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Order, error) {
	var order model.Order
	cdb := tx.Model(&model.Order{})
	err := cdb.Scopes(scopes...).Where(conditions).Updates(data).First(&order).Error

	return order, err
}

func (r *OrderRepository) DeleteByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) error {
	cdb := r.DB.WithContext(ctx)
	return cdb.Scopes(scopes...).Where(conditions).Delete(&model.Order{}).Error
}

func (r *OrderRepository) DeleteByConditionsWithTx(
	tx *gorm.DB,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) error {
	return tx.Scopes(scopes...).Where(conditions).Delete(&model.Order{}).Error
}
