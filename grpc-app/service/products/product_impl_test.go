package products

import (
	"context"
	"errors"
	"go/service1/config"
	"testing"
	"time"

	"go.uber.org/zap"

	pbProducts "go/service1/grpc-app/protos/products"
	"go/service1/shared/models/entities/table"
	pMock "go/service1/shared/usecase/products/mocks"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var dataProduct = &pbProducts.Product{
	UserId:          uuid.New().String(),
	ProductId:       uuid.New().String(),
	ProductName:     "Kue Basi",
	ProductCategory: "electronic",
	ProductInfo:     "Kue Basi tapi enak",
	ProductPrice:    100000,
	ProductSell:     80000,
	ProductStock:    20,
	ProductImage:    "example.com/images/userid/productid/images1.jpg",
}

var product = &table.Product{
	UserId:       dataProduct.UserId,
	ProductId:    dataProduct.ProductId,
	ProductName:  dataProduct.ProductName,
	CategoryName: dataProduct.ProductCategory,
	ProductInfo:  dataProduct.ProductInfo,
	ProductPrice: dataProduct.ProductPrice,
	ProductSell:  dataProduct.ProductSell,
	ProductStock: dataProduct.ProductStock,
	ProductImage: dataProduct.ProductImage,
}

var (
	ctx                 context.Context
	ctrl                *gomock.Controller
	mockProductsUseCase *pMock.MockProductUseCase
	productsService     ProductsService
	any                 gomock.Matcher
)

func TestProductsService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GRPC Products Service")
}

var _ = BeforeSuite(func() {
	config.LoadEnvFile()
	ctx = context.Background()
	ctrl = gomock.NewController(GinkgoT())
	mockProductsUseCase = pMock.NewMockProductUseCase(ctrl)
	productsService = NewProductService(mockProductsUseCase, zap.NewExample())
	any = gomock.Any()

	// mock for usecase insert product
	mockProductsUseCase.EXPECT().InsertProduct(
		ctx,
		&pbProducts.Request{
			Data: dataProduct,
		},
	).Return(product, nil)

	mockProductsUseCase.EXPECT().InsertProduct(any, any).Return(nil, errors.New("failed insert product"))

	// mock for usecase get all product
	mockProductsUseCase.EXPECT().GetAllProduct(
		ctx,
		&pbProducts.RequestGetAllProduct{
			Data: &pbProducts.RequestGetAllProduct_Data{
				Limit: 5,
			},
		},
	).Return([]table.Product{
		*product,
	}, nil)

	mockProductsUseCase.EXPECT().GetAllProduct(any, any).Return(nil, errors.New("failed get all product"))

	// mock for usecase get product by id
	mockProductsUseCase.EXPECT().GetProductById(
		ctx,
		&pbProducts.Request{
			Data: dataProduct,
		},
	).Return(product, nil)

	mockProductsUseCase.EXPECT().GetProductById(any, any).Return(nil, errors.New("failed get product by id"))

	// mock for usecase update product by id
	mockProductsUseCase.EXPECT().UpdateProduct(
		ctx,
		&pbProducts.Request{
			Data: dataProduct,
		},
	).Return(product, nil)

	mockProductsUseCase.EXPECT().UpdateProduct(any, any).Return(nil, errors.New("failed get product by id"))

	// mock for usecase update product by id
	mockProductsUseCase.EXPECT().DeleteProduct(
		ctx,
		&pbProducts.Request{
			Data: dataProduct,
		},
	).Return(product, nil)

	mockProductsUseCase.EXPECT().DeleteProduct(any, any).Return(nil, errors.New("failed get product by id"))
})

var _ = AfterSuite(func() {
	ctrl.Finish()
})

var _ = Describe("Test Grpc Service Products", Ordered, func() {
	Context("Insert Product", func() {
		It("SUCCESS", func() {
			result, err := productsService.InsertProduct(ctx, &pbProducts.Request{
				Data: dataProduct,
			})

			Expect(err).To(BeNil())
			Expect(result.Info.Status).To(Equal("ok"))
			Expect(result.Data.ProductName).To(Equal(dataProduct.ProductName))
			Expect(result.Data.ProductCategory).To(Equal(dataProduct.ProductCategory))
			Expect(result.Data.ProductInfo).To(Equal(dataProduct.ProductInfo))
			Expect(result.Data.ProductPrice).To(Equal(dataProduct.ProductPrice))
			Expect(result.Data.ProductSell).To(Equal(dataProduct.ProductSell))
			Expect(result.Data.ProductStock).To(Equal(dataProduct.ProductStock))
		})

		It("FAIL", func() {
			_, err := productsService.InsertProduct(ctx, &pbProducts.Request{
				Data: dataProduct,
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsService.InsertProduct(c, &pbProducts.Request{
				Data: dataProduct,
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Get All Product", func() {
		It("SUCCESS", func() {
			result, err := productsService.GetAllProduct(
				ctx,
				&pbProducts.RequestGetAllProduct{
					Data: &pbProducts.RequestGetAllProduct_Data{
						Limit: 5,
					},
				},
			)

			Expect(err).To(BeNil())
			Expect(result.Info.Status).To(Equal("ok"))
			Expect(len(result.Data)).NotTo(Equal(0))

			Expect(len(result.Data)).ToNot(Equal(0))

			for _, product := range result.Data {
				if product.ProductId == dataProduct.ProductId {
					Expect(product.UserId).To(Equal(dataProduct.UserId))
					return
				}
			}

			Fail("test failed. err : product not found")
		})

		It("FAIL", func() {
			_, err := productsService.GetAllProduct(
				ctx,
				&pbProducts.RequestGetAllProduct{
					Data: &pbProducts.RequestGetAllProduct_Data{
						Limit: 5,
					},
				},
			)

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsService.GetAllProduct(
				c,
				&pbProducts.RequestGetAllProduct{
					Data: &pbProducts.RequestGetAllProduct_Data{
						Limit: 5,
					},
				},
			)

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Get Product By Id", func() {
		It("SUCCESS", func() {
			result, err := productsService.GetProductById(
				ctx,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).To(BeNil())
			Expect(result.Info.Status).To(Equal("ok"))
			Expect(result.Data.ProductName).To(Equal(dataProduct.ProductName))
			Expect(result.Data.ProductCategory).To(Equal(dataProduct.ProductCategory))
			Expect(result.Data.ProductInfo).To(Equal(dataProduct.ProductInfo))
			Expect(result.Data.ProductPrice).To(Equal(dataProduct.ProductPrice))
			Expect(result.Data.ProductSell).To(Equal(dataProduct.ProductSell))
			Expect(result.Data.ProductStock).To(Equal(dataProduct.ProductStock))
		})

		It("FAIL", func() {
			_, err := productsService.GetProductById(
				ctx,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsService.GetProductById(
				c,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Update Product", func() {
		It("SUCCESS", func() {
			result, err := productsService.UpdateProduct(
				ctx,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).To(BeNil())
			Expect(result.Info.Status).To(Equal("ok"))
			Expect(result.Data.ProductName).To(Equal(dataProduct.ProductName))
			Expect(result.Data.ProductCategory).To(Equal(dataProduct.ProductCategory))
			Expect(result.Data.ProductInfo).To(Equal(dataProduct.ProductInfo))
			Expect(result.Data.ProductPrice).To(Equal(dataProduct.ProductPrice))
			Expect(result.Data.ProductSell).To(Equal(dataProduct.ProductSell))
			Expect(result.Data.ProductStock).To(Equal(dataProduct.ProductStock))
		})

		It("FAIL", func() {
			_, err := productsService.UpdateProduct(
				ctx,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsService.UpdateProduct(
				c,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Delete Product", func() {
		It("SUCCESS", func() {
			result, err := productsService.DeleteProduct(
				ctx,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).To(BeNil())
			Expect(result.Info.Status).To(Equal("ok"))
			Expect(result.Data.UserId).To(Equal(dataProduct.UserId))
			Expect(result.Data.ProductId).To(Equal(dataProduct.ProductId))
			Expect(result.Data.ProductInfo).To(Equal(""))
			Expect(result.Data.ProductPrice).To(Equal(uint32(0)))
			Expect(result.Data.ProductSell).To(Equal(uint32(0)))
			Expect(result.Data.ProductStock).To(Equal(uint32(0)))
		})

		It("FAIL", func() {
			_, err := productsService.DeleteProduct(
				ctx,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsService.DeleteProduct(
				c,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})
})
