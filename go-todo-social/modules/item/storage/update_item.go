package storage

import (
	"context"
	"todo-list-social/common"
	"todo-list-social/modules/item/model"
)

func (s *sqlStore) UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error {

	if err := s.db.Where(cond).Updates(dataUpdate).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
 }
