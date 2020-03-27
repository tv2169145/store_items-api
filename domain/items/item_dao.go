package items

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tv2169145/store_items-api/clients/elasticsearch"
	"github.com/tv2169145/store_utils-go/rest_errors"
	"strings"
)

const (
	indexItems = "items"
	typeItem = "_doc"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}

func (i *Item) Get() rest_errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return rest_errors.NewNotFoundError(fmt.Sprintf("no item found with id %s", i.Id))
		}
		return rest_errors.NewInternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id), errors.New("database error"))
	}
	resultBytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", err)
	}
	fmt.Println(string(resultBytes))
	if err := json.NewDecoder(bytes.NewReader(resultBytes)).Decode(&i); err != nil {
		return rest_errors.NewInternalServerError("error when trying to parse database response", err)
	}
	i.Id = itemId
	return nil
}
