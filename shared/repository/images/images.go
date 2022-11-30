package images

import (
	pb "go/service1/grpc-app/protos/images"
)

type ImageRepository interface {
	InsertProductId(input *pb.DataProduct) (*pb.ResponseInsertProductID, error)
}
