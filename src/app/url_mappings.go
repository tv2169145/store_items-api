package app

import (
	"github.com/tv2169145/store_items-api/src/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Update).Methods(http.MethodPatch)
}
