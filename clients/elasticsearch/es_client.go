package elasticsearch

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/tv2169145/store_utils-go/logger"
	"time"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface{
	setClient(*elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
	Delete(string, string, string) error
	Update(string, string, string, interface{}) error
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().Index(index).Type(docType).BodyJson(doc).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index , docType, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().Index(index).Type(docType).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search(index).Query(query).RestTotalHitsAsInt(true).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Delete(index, docType, id string) error {
	ctx := context.Background()
	result, err := c.client.Delete().Index(index).Type(docType).Id(id).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to delete document in index %s and id %s", index, id), err)
		return err
	}
	if result.Result != "deleted" {
		err := errors.New("database error")
		logger.Error("error when trying to delete item by id", err)
		return err
	}

	return nil
}

func (c *esClient) Update(index, docType, id string, item interface{}) error {
	ctx := context.Background()
	//updateData := make(map[string]interface{})
	//if strings.TrimSpace(item.Status) != "" {
	//	updateData["status"] = item.Status
	//}
	//if strings.TrimSpace(item.Description.PlainText) != "" {
	//	updateData["description"] = item.Description
	//}
	//if strings.TrimSpace(item.Title) != "" {
	//	updateData["title"] = item.Title
	//}
	//updateData["available_quantity"] = item.AvailableQuantity

	result, err := c.client.Update().Index(index).Type(docType).Id(id).Doc(item).Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to update document in index %s and id %s", index, id), err)
		return err
	}
	fmt.Println(result)

	return nil
}
