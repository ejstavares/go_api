package database

import (
	"database/sql"

	"github.com/ejstavares/goapi/internal/entity"
)

type CategoryDB struct {
	db *sql.DB // native lib of go
}
func NewCategoryDB (db *sql.DB) *CategoryDB{
	return &CategoryDB{db: db}
}

func (cd *CategoryDB)  GetCategories() ([]*entity.Category, error){
	rows, err := cd.db.Query("select id, name from categories")

	if err != nil {
		return nil, err
	}
	defer rows.Close() // depois de todo o código a baixo for executado, feicha a conexão

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Exec("insert into categories (id, name) values(?,?)",category.ID, category.Name)

	if err != nil {
		return "", err
	}

	return category.ID, nil
}
func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category

	err := cd.db.QueryRow("select id, name from categories where id = ?", id).Scan(&category.ID, &category.Name)
	if err != nil {
		return nil, err
	}

	return &category, nil
}