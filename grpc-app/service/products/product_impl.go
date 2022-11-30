package products

import (
	"context"
	"go/service1/common"
	protos "go/service1/grpc-app/protos/auth"
	pb "go/service1/grpc-app/protos/products"
	"go/service1/shared/usecase/products"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

type ProductsServiceImpl struct {
	pb.UnimplementedProductRPCServer
	productUC products.ProductUseCase
	logger    *zap.Logger
}

func NewProductService(productUC products.ProductUseCase, logger *zap.Logger) pb.ProductRPCServer {
	return &ProductsServiceImpl{
		productUC: productUC,
		logger:    logger,
	}
}

func (p *ProductsServiceImpl) InsertProduct(c context.Context, i *pb.Request) (*pb.Response, error) {
	select {
	case <-c.Done():
		log.Print("canceled from service insert product")
		return nil, c.Err()
	default:
		res, err := p.productUC.InsertProduct(c, i)

		if err != nil {
			p.logger.Info(err.Error())
			err = status.Errorf(3, err.Error())
			return nil, err
		}

		i.Data.ProductId = res.ProductId
		i.Data.CreatedDate = res.CreatedDate
		i.Data.LastUpdate = res.LastUpdate

		return &pb.Response{
			Info: &protos.Info{
				Code:   0,
				Status: "ok",
			},

			Data: i.Data,
		}, nil
	}
}

func (p *ProductsServiceImpl) GetAllProduct(c context.Context, i *pb.RequestGetAllProduct) (*pb.ResponseGetAllProduct, error) {
	select {
	case <-c.Done():
		log.Print("canceled from service get all product")
		return nil, c.Err()
	default:
		products, err := p.productUC.GetAllProduct(c, i)

		if err != nil {
			p.logger.Info(err.Error())
			err = status.Errorf(3, err.Error())
			return nil, err
		}

		pbProdcuts := common.MapProductsToPB(products)

		return &pb.ResponseGetAllProduct{
			Info: &protos.Info{
				Code:   0,
				Status: "ok",
			},
			Data: pbProdcuts,
		}, nil
	}
}

func (p *ProductsServiceImpl) GetProductById(c context.Context, i *pb.Request) (*pb.Response, error) {
	select {
	case <-c.Done():
		log.Print("canceled from service get product by id")
		return nil, c.Err()
	default:
		product, err := p.productUC.GetProductById(c, i)

		if err != nil {
			p.logger.Info(err.Error())
			err = status.Errorf(3, err.Error())
			return nil, err
		}

		pbProduct := common.MapProductToPB(product)

		return &pb.Response{
			Info: &protos.Info{
				Code:   0,
				Status: "ok",
			},
			Data: pbProduct,
		}, nil
	}
}

func (p *ProductsServiceImpl) UpdateProduct(c context.Context, i *pb.Request) (*pb.Response, error) {
	select {
	case <-c.Done():
		log.Print("canceled from service update product")
		return nil, c.Err()
	default:
		updateProduct, err := p.productUC.UpdateProduct(c, i)

		if err != nil {
			p.logger.Info(err.Error())
			err = status.Errorf(3, err.Error())
			return nil, err
		}

		i.Data.LastUpdate = updateProduct.LastUpdate

		return &pb.Response{
			Info: &protos.Info{
				Code:   0,
				Status: "ok",
			},
			Data: i.Data,
		}, nil
	}
}

func (p *ProductsServiceImpl) DeleteProduct(c context.Context, i *pb.Request) (*pb.Response, error) {
	select {
	case <-c.Done():
		log.Print("canceled from service delete product")
		return nil, c.Err()
	default:
		deleteProduct, err := p.productUC.DeleteProduct(c, i)

		if err != nil {
			p.logger.Info(err.Error())
			err = status.Errorf(3, err.Error())
			return nil, err
		}

		return &pb.Response{
			Info: &protos.Info{
				Code:   0,
				Status: "ok",
			},
			Data: &pb.Product{
				UserId:    deleteProduct.UserId,
				ProductId: deleteProduct.ProductId,
			},
		}, nil
	}
}
