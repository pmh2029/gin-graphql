package repository

import (
	"context"
	"gin-graphql/app/model"
	"gin-graphql/pkg/utils"

	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(
		ctx context.Context,
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.User, error)
	CreateWithTx(
		tx *gorm.DB,
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.User, error)
	TakeByConditionsWithScopes(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.User, error)
	FindByConditionsWithScopes(
		ctx context.Context,
		conditions map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) ([]model.User, error)
	UpdateByConditions(
		ctx context.Context,
		conditions map[string]interface{},
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.User, error)
	UpdateByConditionsWithTx(
		tx *gorm.DB,
		conditions map[string]interface{},
		data map[string]interface{},
		scopes ...func(db *gorm.DB) *gorm.DB,
	) (model.User, error)
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

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db}
}

func (r *UserRepository) Create(
	ctx context.Context,
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.User, error) {
	var user model.User
	err := utils.MapToStruct(data, &user)
	if err != nil {
		return user, err
	}

	cdb := r.DB.WithContext(ctx)
	err = cdb.Scopes(scopes...).Create(&user).Error

	return user, err
}

func (r *UserRepository) CreateWithTx(
	tx *gorm.DB,
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.User, error) {
	var user model.User
	err := utils.MapToStruct(data, &user)
	if err != nil {
		return user, err
	}

	err = tx.Scopes(scopes...).Create(&user).Error
	return user, err
}

func (r *UserRepository) TakeByConditionsWithScopes(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.User, error) {
	var user model.User
	cdb := r.DB.WithContext(ctx)
	err := cdb.Scopes(scopes...).Where(conditions).Take(&user).Error

	return user, err
}

func (r *UserRepository) FindByConditionsWithScopes(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) ([]model.User, error) {
	var products []model.User
	cdb := r.DB.WithContext(ctx)
	err := cdb.Scopes(scopes...).Where(conditions).Find(&products).Error

	return products, err
}

func (r *UserRepository) UpdateByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.User, error) {
	var user model.User
	cdb := r.DB.WithContext(ctx)
	err := cdb.Model(&model.User{}).Scopes(scopes...).Where(conditions).Updates(data).First(&user).Error

	return user, err
}

func (r *UserRepository) UpdateByConditionsWithTx(
	tx *gorm.DB,
	conditions map[string]interface{},
	data map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) (model.User, error) {
	var user model.User
	cdb := tx.Model(&model.User{})
	err := cdb.Scopes(scopes...).Where(conditions).Updates(data).First(&user).Error

	return user, err
}

func (r *UserRepository) DeleteByConditions(
	ctx context.Context,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) error {
	cdb := r.DB.WithContext(ctx)
	return cdb.Scopes(scopes...).Where(conditions).Delete(&model.User{}).Error
}

func (r *UserRepository) DeleteByConditionsWithTx(
	tx *gorm.DB,
	conditions map[string]interface{},
	scopes ...func(db *gorm.DB) *gorm.DB,
) error {
	return tx.Scopes(scopes...).Where(conditions).Delete(&model.User{}).Error
}
