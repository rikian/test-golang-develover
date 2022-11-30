package common

import (
	pbProduct "go/service1/grpc-app/protos/products"
	pbUsers "go/service1/grpc-app/protos/users"
	"go/service1/shared/models/entities/table"
)

// product
func MapProductToPB(t *table.Product) *pbProduct.Product {
	return &pbProduct.Product{
		UserId:          t.UserId,
		ProductId:       t.ProductId,
		ProductCategory: t.CategoryName,
		ProductName:     t.ProductName,
		ProductStock:    t.ProductStock,
		ProductPrice:    t.ProductPrice,
		CreatedDate:     t.CreatedDate,
		LastUpdate:      t.LastUpdate,
		ProductImage:    t.ProductImage,
		ProductSell:     t.ProductSell,
		ProductInfo:     t.ProductInfo,
	}
}

// products
func MapProductsToPB(t []table.Product) []*pbProduct.Product {
	var produts []*pbProduct.Product

	for _, product := range t {
		produts = append(produts, MapProductToPB(&product))
	}

	return produts
}

// user
func MapUserToPb(t *table.User) *pbUsers.User {
	user := &pbUsers.User{
		UserId:       t.UserId,
		UserEmail:    t.UserEmail,
		UserName:     t.UserName,
		UserImage:    t.UserImage,
		UserPassword: t.UserPassword,
		UserSession:  t.UserSession,
		UserStatus:   t.UserStatus.Status,
		CreatedDate:  t.CreatedDate,
		LastUpdate:   t.LastUpdate,
	}

	if len(t.Products) != 0 {
		products := MapProductsToPB(t.Products)
		user.Products = products
		return user
	} else {
		return user
	}
}
