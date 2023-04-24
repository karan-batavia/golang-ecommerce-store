package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/domain"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/repository/interfaces"
	service "github.com/Noush-012/Project-eCommerce-smart_gads/pkg/useCase/interfaces"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/request"
	"github.com/Noush-012/Project-eCommerce-smart_gads/pkg/utils/response"
)

type productUseCase struct {
	ProductRepository interfaces.ProductRepository
}

func NewProductUseCase(ProdRepo interfaces.ProductRepository) service.ProductService {
	return &productUseCase{ProductRepository: ProdRepo}
}

func (p *productUseCase) AddProduct(ctx context.Context, product domain.Product) error {
	// Check the product already exists in databse
	if dbProd, err := p.ProductRepository.FindProduct(ctx, product); err != nil {
		return err
	} else if dbProd.ID != 0 {
		return fmt.Errorf("product already exist with %s product name", dbProd.ProductName)
	}
	return p.ProductRepository.SaveProduct(ctx, product)

}
func (p *productUseCase) AddBrand(ctx context.Context, brand domain.Brand) error {
	// check if req brand already exists in db
	dbBrand, _ := p.ProductRepository.FindBrand(ctx, brand)
	if dbBrand.ID != 0 {
		return fmt.Errorf("brand already exist with %s name", brand.BrandName)
	}
	if err := p.ProductRepository.SaveBrand(ctx, brand); err != nil {
		return err
	}

	return nil

}

// to get all product
func (p *productUseCase) GetProducts(ctx context.Context, page request.ReqPagination) (products []response.ResponseProduct, err error) {
	return p.ProductRepository.GetAllProducts(ctx, page)
}

// to update product
func (p *productUseCase) UpdateProduct(ctx context.Context, product domain.Product) error {
	// validate the product_id
	checkProduct, err := p.ProductRepository.FindProductByID(ctx, product.ID)
	if err != nil {
		return err
	} else if checkProduct.ProductName == "" {
		return errors.New("invalid product_id")
	}

	// check the given product_name already exist or not
	checkProduct, err = p.ProductRepository.FindProduct(ctx, domain.Product{ProductName: product.ProductName})
	if err != nil {
		return err
	} else if checkProduct.ID != 0 && checkProduct.ID != product.ID {
		return errors.New("can't update the product \nthere is alread a product exist with this product_name")
	}

	return p.ProductRepository.UpdateProduct(ctx, product)
}
