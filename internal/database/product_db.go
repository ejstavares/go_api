package database

import (
	"database/sql"
	"github.com/ejstavares/goapi/internal/entity"
)

type ProductDB struct {
	db *sql.DB // native lib of go
}
func NewProductDB (db *sql.DB) *ProductDB{
	return &ProductDB{db: db}
}

func (cd *ProductDB)  GetProducts() ([]*entity.Product, error){
	rows, err := cd.db.Query("select id, name, price, category_id, image_url, description from products")

	if err != nil {
		return nil, err
	}
	defer rows.Close() // depois de todo o código a baixo for executado, feicha a conexão

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (cd *ProductDB) CreateProduct(product *entity.Product) (*entity.Product, error) {
	_, err := cd.db.Exec("insert into products (id, name, price, category_id, image_url, description) values(?,?,?,?,?,?)",product.ID, product.Name, product.Price, product.CategoryID, product.ImageURL, product.Description)

	if err != nil {
		return nil, err
	}

	return product, nil
}
func (cd *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product

	err := cd.db.QueryRow("select id, name, price, category_id, image_url, description from products where id = ?", id).
			Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL, &product.Description)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
func (cd *ProductDB) GetProductByCategoryID(categoryID string) ([]*entity.Product, error) {

	rows, err := cd.db.Query("select id, name, price, category_id, image_url, description from products where category_id = ?", categoryID)

	if err != nil {
		return nil, err
	}

	defer rows.Close() // depois de todo o código a baixo for executado, feicha a conexão

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CategoryID, &product.ImageURL, &product.Description); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}