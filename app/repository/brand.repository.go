package repository

import (
	"context"
	"gin-graphql/app/model"
	"gin-graphql/pkg/utils"

	"gorm.io/gorm"
)

type BrandRepositoryInterface interface {
	Create(
		ctx context.Context,
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Brand, error)
	CreateWithTx(
		tx *gorm.DB,
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Brand, error)
	TakeByConditionsWithScopes(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Brand, error)
	FindByConditionsWithScopes(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) ([]model.Brand, error)
	UpdateByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Brand, error)
	UpdateByConditionsWithTx(
		tx *gorm.DB,
		conditions map[string]interface{},
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.Brand, error)
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

type BrandRepository struct {
	DB *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepositoryInterface {
	return &BrandRepository{db}
}

func (r *BrandRepository) Create(
	ctx context.Context,
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Brand, error) {
	var brand model.Brand
	err := utils.MapToStruct(data, &brand)
	if err != nil {
		return brand, err
	}

	cdb := r.DB.WithContext(ctx)
	err = cdb.Scopes(scopes...).Create(&brand).Error

	return brand, err
}

func (r *BrandRepository) CreateWithTx(
	tx *gorm.DB,
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Brand, error) {
	var brand model.Brand
	err := utils.MapToStruct(data, &brand)
	if err != nil {
		return brand, err
	}

	err = tx.Scopes(scopes...).Create(&brand).Error
	return brand, err
}

func (r *BrandRepository) TakeByConditionsWithScopes(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Brand, error) {
	var brand model.Brand
	cdb := r.DB.WithContext(ctx)
	err := cdb.Scopes(scopes...).Where(conditions).Take(&brand).Error

	return brand, err
}

func (r *BrandRepository) FindByConditionsWithScopes(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) ([]model.Brand, error) {
	var brands []model.Brand
	cdb := r.DB.WithContext(ctx)
	err := cdb.Scopes(scopes...).Where(conditions).Find(&brands).Error

	return brands, err
}

func (r *BrandRepository) UpdateByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Brand, error) {
	var brand model.Brand
	cdb := r.DB.WithContext(ctx)
	err := cdb.Model(&model.Brand{}).Scopes(scopes...).Where(conditions).Updates(data).First(&brand).Error

	return brand, err
}

func (r *BrandRepository) UpdateByConditionsWithTx(
	tx *gorm.DB,
	conditions map[string]interface{},
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.Brand, error) {
	var brand model.Brand
	cdb := tx.Model(&model.Brand{})
	err := cdb.Scopes(scopes...).Where(conditions).Updates(data).First(&brand).Error

	return brand, err
}

func (r *BrandRepository) DeleteByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) error {
	cdb := r.DB.WithContext(ctx)
	return cdb.Scopes(scopes...).Where(conditions).Delete(&model.Brand{}).Error
}

func (r *BrandRepository) DeleteByConditionsWithTx(
	tx *gorm.DB,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) error {
	return tx.Scopes(scopes...).Where(conditions).Delete(&model.Brand{}).Error
}
