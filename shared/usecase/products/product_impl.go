package products

import (
	"context"
	pb "go/service1/grpc-app/protos/products"
	"go/service1/shared/models/entities/table"
	"go/service1/shared/repository/products"
	"log"

	"go.uber.org/zap"
)

type ProductUseCaseImpl struct {
	logger      *zap.Logger
	productRepo products.ProductRepository
}

func NewProductUseCaseImpl(logger *zap.Logger, productRepo products.ProductRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		logger:      logger,
		productRepo: productRepo,
	}
}

func (p *ProductUseCaseImpl) InsertProduct(c context.Context, i *pb.Request) (*table.Product, error) {
	select {
	case <-c.Done():
		log.Print("canceled from usecase insert product")
		return nil, c.Err()
	default:
		res, err := p.productRepo.InsertProduct(c, i)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func (p *ProductUseCaseImpl) GetAllProduct(c context.Context, i *pb.RequestGetAllProduct) ([]table.Product, error) {
	select {
	case <-c.Done():
		log.Print("canceled from usecase get all product")
		return nil, c.Err()
	default:
		res, err := p.productRepo.GetAllProduct(c, i)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func (p *ProductUseCaseImpl) GetProductById(c context.Context, i *pb.Request) (*table.Product, error) {
	select {
	case <-c.Done():
		log.Print("canceled from usecase get product by id")
		return nil, c.Err()
	default:
		res, err := p.productRepo.GetProductById(c, i)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func (p *ProductUseCaseImpl) UpdateProduct(c context.Context, i *pb.Request) (*table.Product, error) {
	select {
	case <-c.Done():
		log.Print("canceled from usecase update product")
		return nil, c.Err()
	default:
		res, err := p.productRepo.UpdateProduct(c, i)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}

func (p *ProductUseCaseImpl) DeleteProduct(c context.Context, i *pb.Request) (*table.Product, error) {
	select {
	case <-c.Done():
		log.Print("canceled from usecase delete product")
		return nil, c.Err()
	default:
		res, err := p.productRepo.DeleteProduct(c, i)

		if err != nil {
			return nil, err
		}

		return res, nil
	}
}
