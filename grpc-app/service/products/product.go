package products

import (
	"context"
	pb "go/service1/grpc-app/protos/products"
)

type ProductsService interface {
	InsertProduct(context.Context, *pb.Request) (*pb.Response, error)
	GetAllProduct(context.Context, *pb.RequestGetAllProduct) (*pb.ResponseGetAllProduct, error)
	GetProductById(context.Context, *pb.Request) (*pb.Response, error)
	UpdateProduct(context.Context, *pb.Request) (*pb.Response, error)
	DeleteProduct(context.Context, *pb.Request) (*pb.Response, error)
}
