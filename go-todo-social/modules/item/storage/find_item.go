package storage

import (
	"context"
	"todo-list-social/common"
	"todo-list-social/modules/item/model"
	"gorm.io/gorm"
)

func (s *sqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecourdNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
