package categorystorage

import (
	"context"
	"errors"
	categorymodel "golang-training/modules/category/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(context context.Context, condition map[string]interface{}) (*categorymodel.Category, error) {
	var data categorymodel.Category

	if err := s.db.WithContext(context).Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &data, nil
}

func (s *sqlStore) FindCategoryExist(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error) {
	var data categorymodel.Category

	if err := s.db.WithContext(ctx).Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &data, nil
}
