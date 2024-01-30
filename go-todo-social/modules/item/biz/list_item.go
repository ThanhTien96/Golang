package biz

import (
	"context"
	"todo-list-social/common"
	"todo-list-social/modules/item/model"
)

type ListItemStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.TodoItem, error)
}

type listItemBiz struct {
	store ListItemStorage
}

func NewListItemBiz(store ListItemStorage) *listItemBiz {
	return &listItemBiz{store: store}
}

func (biz *listItemBiz) GetListItem(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.TodoItem, error) {

	data, err := biz.store.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName ,err)
	}

	return data, nil

}
