package products

import (
	"context"
	"errors"
	pb "go/service1/grpc-app/protos/products"
	"go/service1/shared/models/entities/table"
	"log"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewProductsRepositoryImpl(db *gorm.DB, logger *zap.Logger) ProductRepository {
	return &ProductRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (p *ProductRepositoryImpl) InsertProduct(c context.Context, i *pb.Request) (*table.Product, error) {
	select {
	case <-c.Done():
		p.logger.Info("Canceled Insert Product")
		log.Print("canceled request insert product")
		return nil, c.Err()
	default:
		var time string = time.Now().Format("20060102150405")
		var productId string = uuid.New().String()
		var product *table.Product = &table.Product{
			UserId:       i.Data.UserId,
			CategoryName: i.Data.ProductCategory,
			ProductId:    productId,
			ProductName:  i.Data.ProductName,
			ProductImage: i.Data.ProductImage,
			ProductInfo:  i.Data.ProductInfo,
			ProductStock: i.Data.ProductStock,
			ProductPrice: i.Data.ProductPrice,
			ProductSell:  i.Data.ProductSell,
			CreatedDate:  time,
			LastUpdate:   time,
		}

		var createProduct *gorm.DB = p.db.Create(product)

		if createProduct.Error != nil {
			p.logger.Info(createProduct.Error.Error())
			return nil, createProduct.Error
		}

		if createProduct.RowsAffected != 1 {
			p.logger.Info("row affected insert product not equal 1")
			return nil, errors.New("row affected insert product not equal 1")
		}

		return product, nil
	}
}

func (p *ProductRepositoryImpl) GetAllProduct(c context.Context, i *pb.RequestGetAllProduct) ([]table.Product, error) {
	select {
	case <-c.Done():
		p.logger.Info("Canceled Insert Product")
		log.Print("canceled request get all product")
		return nil, c.Err()
	default:
		var products = []table.Product{}
		var getAllProduct *gorm.DB = p.db.Find(&products)

		if getAllProduct.Error != nil {
			p.logger.Info(getAllProduct.Error.Error())
			return nil, getAllProduct.Error
		}

		if getAllProduct.RowsAffected == 0 {
			p.logger.Info("row affected getAllProduct equal 0")
			return nil, errors.New("row affected getAllProduct equal 0")
		}

		return products, nil
	}
}

func (p *ProductRepositoryImpl) GetProductById(c context.Context, i *pb.Request) (*table.Product, error) {
	select {
	case <-c.Done():
		p.logger.Info("Canceled Insert Product")
		log.Print("canceled request get product by id")
		return nil, c.Err()
	default:
		var product *table.Product = &table.Product{}

		getProductById := p.db.Model(product).
			Where("product_id = ? AND product_name = ?", i.Data.ProductId, i.Data.ProductName).
			First(product)

		if getProductById.Error != nil {
			p.logger.Info(getProductById.Error.Error())
			return nil, getProductById.Error
		}

		if getProductById.RowsAffected != 1 {
			p.logger.Info("row affected getProductById not equal 1")
			return nil, errors.New("row affected getProductById not equal 1")
		}

		return product, nil
	}
}

func (p *ProductRepositoryImpl) UpdateProduct(c context.Context, i *pb.Request) (*table.Product, error) {
	select {
	case <-c.Done():
		p.logger.Info("canceled request update product")
		log.Print("canceled request update product")
		return nil, c.Err()
	default:
		var time string = time.Now().Format("20060102150405")
		var product *table.Product = &table.Product{}

		p.db.First(product)
		product.CategoryName = i.Data.ProductCategory
		product.ProductName = i.Data.ProductName
		product.ProductImage = i.Data.ProductImage
		product.ProductInfo = i.Data.ProductInfo
		product.ProductStock = i.Data.ProductStock
		product.ProductPrice = i.Data.ProductPrice
		product.ProductSell = i.Data.ProductSell
		product.LastUpdate = time

		updateProduct := p.db.Save(product)

		if updateProduct.Error != nil {
			p.logger.Info(updateProduct.Error.Error())
			return nil, updateProduct.Error
		}

		if updateProduct.RowsAffected != 1 {
			p.logger.Info("row affected updateProduct not equal 1")
			return nil, errors.New("row affected updateProduct not equal 1")
		}

		return product, nil
	}
}

func (p *ProductRepositoryImpl) DeleteProduct(c context.Context, i *pb.Request) (*table.Product, error) {
	tx := p.db.Begin()

	defer func() {
		r := recover()
		if r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		p.logger.Info(tx.Error.Error())
		return nil, tx.Error
	}

	select {
	case <-c.Done():
		p.logger.Info("Canceled Insert Product")
		log.Print("canceled request delete product")
		tx.Rollback()
		return nil, c.Err()
	default:
		// logic
		deleteProduct := tx.Where("user_id = ? AND product_id=?", i.Data.UserId, i.Data.ProductId).
			Delete(&table.Product{})

		if deleteProduct.Error != nil {
			p.logger.Info(deleteProduct.Error.Error())
			return nil, deleteProduct.Error
		}

		if deleteProduct.RowsAffected != 1 {
			tx.Rollback()
			p.logger.Info("row affected deleteProduct not equal 1")
			return nil, errors.New("row affected deleteProduct not equal 1")
		}

		// delete image product from directory server image
		// deleteDirImage, err := p.h.DeleteImage(i.UserId, i.ProductId)

		// if err != nil {
		// 	p.logger.Info(err.Error())
		// 	tx.Rollback()
		// 	return nil, err
		// }

		// if deleteDirImage.Message != "ok" {
		// 	p.logger.Info(deleteDirImage.Message)
		// 	tx.Rollback()
		// 	return nil, errors.New(deleteDirImage.Message)
		// }

		// comit
		comit := tx.Commit()

		if comit.Error != nil {
			p.logger.Info(comit.Error.Error())
			return nil, comit.Error
		}

		return &table.Product{
			UserId:    i.Data.UserId,
			ProductId: i.Data.ProductId,
		}, nil
	}
}
