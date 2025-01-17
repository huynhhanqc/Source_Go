package biz

import (
	"context"
	"golandtitorial/modules/item/model"
)

type GetItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
}
type getItemBiz struct {
	store GetItemStorage
}

func NewGetItemBiz(store GetItemStorage) *getItemBiz {
	return &getItemBiz{store: store}
}

func (biz *getItemBiz) GetItembyId(ctx context.Context, id string) (*model.TodoItem, error) {

	data, err := biz.store.GetItem(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	return data, nil
}
