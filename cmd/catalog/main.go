package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ejstavares/goapi/internal/database"
	"github.com/ejstavares/goapi/internal/services"
	"github.com/ejstavares/goapi/internal/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main ()  {
	db, err := sql.Open("mysql","root:root@tcp(localhost:3306)/go_api_db")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := services.NewCategoryService(*categoryDB)
	webCategoryHandler:= webserver.NewWebCategoryHandler(*categoryService)

	productDB := database.NewProductDB(db)
	productService := services.NewProductService(*productDB)
	webProductHandler:= webserver.NewWebProductHandler(*productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/categories", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductByCategoryID)
	c.Get("/products", webProductHandler.GetProducts)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Print("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}