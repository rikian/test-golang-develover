package interceptor

import (
	"context"
	"fmt"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		select {
		case <-ctx.Done():
			log.Print("canceled from auth iterceptor")
			return nil, ctx.Err()
		default:
			log.Print("incoming request : " + info.FullMethod)
			log.Print("processing...")
			logger.Info("incoming request : " + info.FullMethod)

			md, ok := metadata.FromIncomingContext(ctx)

			if !ok {
				logger.Info("meta data not found")
				return nil, status.Error(codes.PermissionDenied, "meta data not found")
			}

			// we can add authentication in metadata before hits the entry point.
			// jwt etc...
			var authenticationMetaData []string = md.Get("hello")

			if len(authenticationMetaData) == 0 || authenticationMetaData[0] != "world" {
				logger.Info("invalid meta data")
				return nil, status.Error(codes.InvalidArgument, "invalid meta data")
			}

			// check if the payload not empty so we can process it
			if fmt.Sprint(req) == "" {
				logger.Info("invalid argument request")
				return nil, status.Error(codes.InvalidArgument, "param cannot be empty")
			}

			// next if authentication valid
			data, err := handler(ctx, req)

			log.Print(data)

			if err != nil {
				logger.Error(err.Error())
				return nil, err
			}

			return data, nil
		}
	}
}
