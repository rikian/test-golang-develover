package products

import (
	"context"
	pb "go/service1/grpc-app/protos/products"
	"go/service1/shared/models/entities/table"
)

type ProductRepository interface {
	InsertProduct(context.Context, *pb.Request) (*table.Product, error)
	GetAllProduct(context.Context, *pb.RequestGetAllProduct) ([]table.Product, error)
	GetProductById(context.Context, *pb.Request) (*table.Product, error)
	UpdateProduct(context.Context, *pb.Request) (*table.Product, error)
	DeleteProduct(context.Context, *pb.Request) (*table.Product, error)
}
