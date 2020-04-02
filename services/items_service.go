package services

import (
	"github.com/tv2169145/store_items-api/domain/items"
	"github.com/tv2169145/store_items-api/domain/queries"
	"github.com/tv2169145/store_utils-go/rest_errors"
)

var (
	ItemService ItemsServiceInterface = &itemService{}
)

type ItemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(queries.EsQuery) ([]items.Item, rest_errors.RestErr)
	Delete(string) rest_errors.RestErr
	Update(items.Item) (*items.Item, rest_errors.RestErr)
}

type itemService struct {

}

func(s *itemService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	//return nil, rest_errors.NewRestError("implement it", http.StatusNotImplemented, "not_implement", nil)
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id:id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Search(query queries.EsQuery) ([]items.Item, rest_errors.RestErr) {
	dao := items.Item{}
	return dao.Search(query)
}

func (s *itemService) Delete(id string) rest_errors.RestErr {
	dao := items.Item{}

	return dao.Delete(id)
}

func (s *itemService) Update(updateItem items.Item) (*items.Item, rest_errors.RestErr) {
	if err := updateItem.Update(); err != nil {
		return nil, err
	}
	return &updateItem, nil
}
