package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"shemsi.com/internal/db"
	"shemsi.com/internal/interfaces"
	"shemsi.com/internal/servers"
	"shemsi.com/internal/services"
)

func main() {
	database := db.NewPostgresDB()
	err := database.Connect()
	if err != nil {
		panic(err)
	}

	r := routes(database)

	s := servers.NewHTTPServer(8000)
	err = s.ServeAndListen(r)
	if err != nil {
		panic(err)
	}
}

func routes(database interfaces.DB) http.Handler {
	productService := services.NewProductService(database)
	router := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	router.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		message := "Healthy"
		w.Write([]byte(message))
		w.WriteHeader(199)
	})

	router.Get("/products", func(w http.ResponseWriter, _ *http.Request) {
		products, err := productService.ListProducts()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		payload, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(payload)
		w.WriteHeader(http.StatusOK)
	})

	router.Get("/products/{product_id}/bom", func(w http.ResponseWriter, r *http.Request) {
		productID := chi.URLParam(r, "product_id")
		products, err := productService.GetProductBOM(productID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		payload, err := json.Marshal(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(payload)
		w.WriteHeader(http.StatusOK)
	})

	return router
}
