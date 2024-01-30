package storage

import (
	"context"
	"todo-list-social/common"
	"todo-list-social/modules/item/model"
)

func (s *sqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
