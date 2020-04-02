package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/tv2169145/store_items-api/domain/items"
	"github.com/tv2169145/store_items-api/domain/queries"
	"github.com/tv2169145/store_items-api/services"
	"github.com/tv2169145/store_items-api/utils/http_utils"
	"github.com/tv2169145/store_oauth-go/oauth"
	"github.com/tv2169145/store_utils-go/rest_errors"
	"net/http"
	"strings"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {

}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("Unauthorized Error!")
		http_utils.RespondError(w, respErr)
		return
	}

	//requestBody, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	restErr := rest_errors.NewBadRequestError("invalid request body")
	//	http_utils.RespondError(w, restErr.Status(), restErr)
	//	return
	//}
	//defer r.Body.Close()
	var itemRequest items.Item
	//if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
	//	restErr := rest_errors.NewBadRequestError("invalid json body")
	//	http_utils.RespondError(w, restErr.Status(), restErr)
	//	return
	//}

	if err := json.NewDecoder(r.Body).Decode(&itemRequest); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid request json body")
		http_utils.RespondError(w, restErr)
		return
	}

	itemRequest.Seller = sellerId
	result, createErr := services.ItemService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemService.Get(itemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	var query queries.EsQuery
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w, apiErr)
		return
	}
	items, searchErr := services.ItemService.Search(query)
	if searchErr != nil {
		http_utils.RespondError(w, searchErr)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, items)
}

func (c *itemsController) Delete(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("Unauthorized Error!")
		http_utils.RespondError(w, respErr)
		return
	}

	vars := mux.Vars(r)
	targetItemId := vars["id"]

	targetItem, err := services.ItemService.Get(targetItemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}

	if targetItem.Seller != sellerId {
		restErr := rest_errors.NewUnauthorizedError("item owner error")
		http_utils.RespondError(w, restErr)
		return
	}

	if err := services.ItemService.Delete(targetItemId); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, targetItem)
}

func (c *itemsController) Update(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("Unauthorized Error!")
		http_utils.RespondError(w, respErr)
		return
	}
	vars := mux.Vars(r)
	targetItemId := vars["id"]
	targetItem, err := services.ItemService.Get(targetItemId)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}

	var item items.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http_utils.RespondError(w, rest_errors.NewBadRequestError("invalid json body"))
		return
	}

	if strings.TrimSpace(item.Title) != "" && item.Title != targetItem.Title {
		targetItem.Title = item.Title
	}

	if strings.TrimSpace(item.Description.PlainText) != "" && item.Description.PlainText != targetItem.Description.PlainText {
		targetItem.Description.PlainText = item.Description.PlainText
	}

	if strings.TrimSpace(item.Status) != "" && item.Status != targetItem.Status {
		targetItem.Status = item.Status
	}

	if item.AvailableQuantity != targetItem.AvailableQuantity {
		targetItem.AvailableQuantity = item.AvailableQuantity
	}
	responseItem, err := services.ItemService.Update(*targetItem)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, responseItem)
}
