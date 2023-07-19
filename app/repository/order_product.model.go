package repository

import (
	"context"
	"gin-graphql/app/model"
	"gin-graphql/pkg/utils"

	"gorm.io/gorm"
)

type OrderProductRepositoryInterface interface {
	Create(
		ctx context.Context,
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.OrderProduct, error)
	CreateWithTx(
		tx *gorm.DB,
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.OrderProduct, error)
	TakeByConditionsWithScopes(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.OrderProduct, error)
	FindByConditionsWithScopes(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) ([]model.OrderProduct, error)
	UpdateByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.OrderProduct, error)
	UpdateByConditionsWithTx(
		tx *gorm.DB,
		conditions map[string]interface{},
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.OrderProduct, error)
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

type OderProductRepository struct {
	DB *gorm.DB
}

func NewOrderProductRepository(db *gorm.DB) OrderProductRepositoryInterface {
	return &OderProductRepository{db}
}

func (r *OderProductRepository) Create(
	ctx context.Context,
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.OrderProduct, error) {
	var orderProduct model.OrderProduct
	err := utils.MapToStruct(data, &orderProduct)
	if err != nil {
		return orderProduct, err
	}

	cdb := r.DB.WithContext(ctx)
	err = cdb.Scopes(scopes...).Create(&orderProduct).Error

	return orderProduct, err
}

func (r *OderProductRepository) CreateWithTx(
	tx *gorm.DB,
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.OrderProduct, error) {
	var orderProduct model.OrderProduct
	err := utils.MapToStruct(data, &orderProduct)
	if err != nil {
		return orderProduct, err
	}

	err = tx.Scopes(scopes...).Create(&orderProduct).Error
	return orderProduct, err
}

func (r *OderProductRepository) TakeByConditionsWithScopes(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.OrderProduct, error) {
	var orderProduct model.OrderProduct
	cdb := r.DB.WithContext(ctx)
	err := cdb.Scopes(scopes...).Where(conditions).Take(&orderProduct).Error

	return orderProduct, err
}

func (r *OderProductRepository) FindByConditionsWithScopes(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) ([]model.OrderProduct, error) {
	var orderProducts []model.OrderProduct
	cdb := r.DB.WithContext(ctx)
	err := cdb.Scopes(scopes...).Where(conditions).Find(&orderProducts).Error

	return orderProducts, err
}

func (r *OderProductRepository) UpdateByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.OrderProduct, error) {
	var orderProduct model.OrderProduct
	cdb := r.DB.WithContext(ctx)
	err := cdb.Model(&model.OrderProduct{}).Scopes(scopes...).Where(conditions).Updates(data).First(&orderProduct).Error

	return orderProduct, err
}

func (r *OderProductRepository) UpdateByConditionsWithTx(
	tx *gorm.DB,
	conditions map[string]interface{},
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.OrderProduct, error) {
	var orderProduct model.OrderProduct
	cdb := tx.Model(&model.OrderProduct{})
	err := cdb.Scopes(scopes...).Where(conditions).Updates(data).First(&orderProduct).Error

	return orderProduct, err
}

func (r *OderProductRepository) DeleteByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) error {
	cdb := r.DB.WithContext(ctx)
	return cdb.Scopes(scopes...).Where(conditions).Delete(&model.OrderProduct{}).Error
}

func (r *OderProductRepository) DeleteByConditionsWithTx(
	tx *gorm.DB,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) error {
	return tx.Scopes(scopes...).Where(conditions).Delete(&model.OrderProduct{}).Error
}
