package controllers

import (
	"encoding/json"
	"github.com/tv2169145/store_items-api/domain/items"
	"github.com/tv2169145/store_items-api/services"
	"github.com/tv2169145/store_items-api/utils/http_utils"
	"github.com/tv2169145/store_oauth-go/oauth"
	"github.com/tv2169145/store_utils-go/rest_errors"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {

}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.RespondError(w, err.Status, err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("Unauthorized Error!")
		http_utils.RespondError(w, respErr.Status(), respErr)
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
		http_utils.RespondError(w, restErr.Status(), restErr)
		return
	}

	itemRequest.Seller = sellerId
	result, createErr := services.ItemService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr.Status(), createErr)
		return
	}
	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
