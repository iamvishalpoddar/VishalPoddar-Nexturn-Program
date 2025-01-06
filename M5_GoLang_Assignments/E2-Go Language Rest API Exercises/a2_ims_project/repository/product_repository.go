package repository

import (
	"InventoryManagementSystem/model"
	"database/sql"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (pr *ProductRepository) CreateProduct(product *model.Product) (*model.Product, error) {
	stmt, err := pr.DB.Prepare("INSERT INTO product(name, price, quantity, category) VALUES(?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(product.Name, product.Price, product.Quantity, product.Category)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	product.ID = int(id)
	return product, nil
}

func (pr *ProductRepository) GetProduct(id int) (*model.Product, error) {
	row := pr.DB.QueryRow("SELECT * FROM product WHERE id = ?", id)
	product := &model.Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Category)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pr *ProductRepository) GetAllProducts() ([]model.Product, error) {
	rows, err := pr.DB.Query("SELECT * FROM product")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Category)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	return products, nil
}

func (pr *ProductRepository) UpdateProduct(product *model.Product) (*model.Product, error) {
	stmt, err := pr.DB.Prepare(`UPDATE product SET name = ?, price = ?, quantity = ?, category = ? WHERE id = ?`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Quantity, product.Category, product.ID)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pr *ProductRepository) DeleteProduct(id int) error {
	stmt, err := pr.DB.Prepare("DELETE FROM product WHERE id = ?")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
