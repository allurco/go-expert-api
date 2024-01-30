package main

import (
	"net/http"

	"github.com/allurco/go-expert-api/internal/entity"
	"github.com/allurco/go-expert-api/internal/infra/database"
	"github.com/allurco/go-expert-api/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	/* _, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	} */

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/product/{id}", productHandler.GetProduct)
	r.Put("/product/{id}/update", productHandler.UpdateProduct)
	r.Delete("/product/{id}/delete", productHandler.DeleteProduct)
	r.Get("/find/products", productHandler.GetProductsPaged)
	http.ListenAndServe(":8000", r)

}
