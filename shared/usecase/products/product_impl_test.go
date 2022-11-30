package products

import (
	"context"
	"errors"
	"go/service1/config"
	pbProducts "go/service1/grpc-app/protos/products"
	"go/service1/shared/models/entities/table"
	prMock "go/service1/shared/repository/products/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
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
	ctx                    context.Context
	ctrl                   *gomock.Controller
	mockProductsRepository *prMock.MockProductRepository
	productsUseCase        ProductUseCase
	any                    gomock.Matcher
)

func TestProductsUseCase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "UseCase Products Suite")
}

var _ = BeforeSuite(func() {
	config.LoadEnvFile()
	ctx = context.Background()
	ctrl = gomock.NewController(GinkgoT())
	mockProductsRepository = prMock.NewMockProductRepository(ctrl)
	productsUseCase = NewProductUseCaseImpl(zap.NewExample(), mockProductsRepository)
	any = gomock.Any()

	// mock for repository insert product
	mockProductsRepository.EXPECT().InsertProduct(
		ctx,
		&pbProducts.Request{
			Data: dataProduct,
		},
	).Return(product, nil)

	mockProductsRepository.EXPECT().InsertProduct(any, any).Return(nil, errors.New("failed insert product"))

	// mock for repository get all product
	mockProductsRepository.EXPECT().GetAllProduct(
		ctx,
		&pbProducts.RequestGetAllProduct{
			Data: &pbProducts.RequestGetAllProduct_Data{
				Limit: 5,
			},
		},
	).Return([]table.Product{
		*product,
	}, nil)

	mockProductsRepository.EXPECT().GetAllProduct(any, any).Return(nil, errors.New("failed get all product"))

	// mock for repository get product by id
	mockProductsRepository.EXPECT().GetProductById(
		ctx,
		&pbProducts.Request{
			Data: dataProduct,
		},
	).Return(product, nil)

	mockProductsRepository.EXPECT().GetProductById(any, any).Return(nil, errors.New("failed get product by id"))

	// mock for repository update product by id
	mockProductsRepository.EXPECT().UpdateProduct(
		ctx,
		&pbProducts.Request{
			Data: dataProduct,
		},
	).Return(product, nil)

	mockProductsRepository.EXPECT().UpdateProduct(any, any).Return(nil, errors.New("failed get product by id"))

	// mock for repository update product by id
	mockProductsRepository.EXPECT().DeleteProduct(
		ctx,
		&pbProducts.Request{
			Data: dataProduct,
		},
	).Return(product, nil)

	mockProductsRepository.EXPECT().DeleteProduct(any, any).Return(nil, errors.New("failed get product by id"))

})

var _ = AfterSuite(func() {
	ctrl.Finish()
})

var _ = Describe("Checking Products Usecase", Ordered, func() {
	Context("NewProductsUseCase", func() {
		It("SUCCESS", func() {
			Expect(productsUseCase).NotTo(BeNil())
		})
	})
	Context("Insert Product", func() {
		It("SUCCESS", func() {
			result, err := productsUseCase.InsertProduct(ctx, &pbProducts.Request{
				Data: dataProduct,
			})

			Expect(err).To(BeNil())
			Expect(result.ProductName).To(Equal(dataProduct.ProductName))
			Expect(result.CategoryName).To(Equal(dataProduct.ProductCategory))
			Expect(result.ProductInfo).To(Equal(dataProduct.ProductInfo))
			Expect(result.ProductPrice).To(Equal(dataProduct.ProductPrice))
			Expect(result.ProductSell).To(Equal(dataProduct.ProductSell))
			Expect(result.ProductStock).To(Equal(dataProduct.ProductStock))
		})

		It("FAIL", func() {
			_, err := productsUseCase.InsertProduct(ctx, &pbProducts.Request{
				Data: dataProduct,
			})

			Expect(err).NotTo(BeNil())
		})

		It("CANCEL", func() {
			c, cancel := context.WithTimeout(ctx, 2*time.Second)
			cancel()

			_, err := productsUseCase.InsertProduct(c, &pbProducts.Request{
				Data: dataProduct,
			})

			Expect(err).NotTo(BeNil())
			Expect(err).To(Equal(c.Err()))
		})
	})

	Context("Get All Product", func() {
		It("SUCCESS", func() {
			result, err := productsUseCase.GetAllProduct(
				ctx,
				&pbProducts.RequestGetAllProduct{
					Data: &pbProducts.RequestGetAllProduct_Data{
						Limit: 5,
					},
				},
			)

			Expect(err).To(BeNil())
			Expect(len(result)).NotTo(Equal(0))

			Expect(len(result)).ToNot(Equal(0))

			for _, product := range result {
				if product.ProductId == dataProduct.ProductId {
					Expect(product.UserId).To(Equal(dataProduct.UserId))
					return
				}
			}

			Fail("test failed. err : product not found")
		})

		It("FAIL", func() {
			_, err := productsUseCase.GetAllProduct(
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

			_, err := productsUseCase.GetAllProduct(
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
			result, err := productsUseCase.GetProductById(
				ctx,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).To(BeNil())
			Expect(result.ProductName).To(Equal(dataProduct.ProductName))
			Expect(result.CategoryName).To(Equal(dataProduct.ProductCategory))
			Expect(result.ProductInfo).To(Equal(dataProduct.ProductInfo))
			Expect(result.ProductPrice).To(Equal(dataProduct.ProductPrice))
			Expect(result.ProductSell).To(Equal(dataProduct.ProductSell))
			Expect(result.ProductStock).To(Equal(dataProduct.ProductStock))
		})

		It("FAIL", func() {
			_, err := productsUseCase.GetProductById(
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

			_, err := productsUseCase.GetProductById(
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
			result, err := productsUseCase.UpdateProduct(
				ctx,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).To(BeNil())
			Expect(result.ProductName).To(Equal(dataProduct.ProductName))
			Expect(result.CategoryName).To(Equal(dataProduct.ProductCategory))
			Expect(result.ProductInfo).To(Equal(dataProduct.ProductInfo))
			Expect(result.ProductPrice).To(Equal(dataProduct.ProductPrice))
			Expect(result.ProductSell).To(Equal(dataProduct.ProductSell))
			Expect(result.ProductStock).To(Equal(dataProduct.ProductStock))
		})

		It("FAIL", func() {
			_, err := productsUseCase.UpdateProduct(
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

			_, err := productsUseCase.UpdateProduct(
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
			result, err := productsUseCase.DeleteProduct(
				ctx,
				&pbProducts.Request{
					Data: dataProduct,
				},
			)

			Expect(err).To(BeNil())
			Expect(result.UserId).To(Equal(dataProduct.UserId))
			Expect(result.ProductId).To(Equal(dataProduct.ProductId))
			Expect(result.ProductInfo).To(Equal(dataProduct.ProductInfo))
			Expect(result.ProductPrice).To(Equal(dataProduct.ProductPrice))
			Expect(result.ProductSell).To(Equal(dataProduct.ProductSell))
			Expect(result.ProductStock).To(Equal(dataProduct.ProductStock))
		})

		It("FAIL", func() {
			_, err := productsUseCase.DeleteProduct(
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

			_, err := productsUseCase.DeleteProduct(
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
