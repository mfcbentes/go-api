package usecase

import (
	"github.com/mfcbentes/go-api/model"
	"github.com/mfcbentes/go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductuseCase(repository repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		repository: repository,
	}
}

func (pu *ProductUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUseCase) GetProductById(id int) (model.Product, error) {
	return pu.repository.GetProductById(id)
}
