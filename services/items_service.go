package services

import (
	"github.com/tv2169145/store_items-api/domain/items"
	"github.com/tv2169145/store_utils-go/rest_errors"
	"net/http"
)

var (
	ItemService ItemsServiceInterface = &itemService{}
)

type ItemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
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
	return nil, rest_errors.NewRestError("implement it", http.StatusNotImplemented, "not_implement", nil)
}
