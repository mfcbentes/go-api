package repository

import (
	"database/sql"
	"fmt"

	"github.com/mfcbentes/go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) *ProductRepository {
	return &ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "select id, product_name, product_price from product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var product model.Product

	for rows.Next() {
		err = rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, product)
	}

	rows.Close()

	return productList, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int

	query, err := pr.connection.Prepare("insert into product (product_name, product_price) values ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(&product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (pr *ProductRepository) GetProductById(id_product int) (model.Product, error) {
	query, err := pr.connection.Prepare("select id, product_name, product_price from product where id = $1")
	if err != nil {
		fmt.Println(err)
		return model.Product{}, err
	}

	var product model.Product

	err = query.QueryRow(id_product).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Product{}, nil
		}

		fmt.Println(err)
		return model.Product{}, err
	}

	query.Close()

	return product, nil
}
