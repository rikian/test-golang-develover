// Code generated by MockGen. DO NOT EDIT.
// Source: shared/repository/products/product.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	products "go/service1/grpc-app/protos/products"
	table "go/service1/shared/models/entities/table"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProductRepository is a mock of ProductRepository interface.
type MockProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryMockRecorder
}

// MockProductRepositoryMockRecorder is the mock recorder for MockProductRepository.
type MockProductRepositoryMockRecorder struct {
	mock *MockProductRepository
}

// NewMockProductRepository creates a new mock instance.
func NewMockProductRepository(ctrl *gomock.Controller) *MockProductRepository {
	mock := &MockProductRepository{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepository) EXPECT() *MockProductRepositoryMockRecorder {
	return m.recorder
}

// DeleteProduct mocks base method.
func (m *MockProductRepository) DeleteProduct(arg0 context.Context, arg1 *products.Request) (*table.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0, arg1)
	ret0, _ := ret[0].(*table.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductRepositoryMockRecorder) DeleteProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProductRepository)(nil).DeleteProduct), arg0, arg1)
}

// GetAllProduct mocks base method.
func (m *MockProductRepository) GetAllProduct(arg0 context.Context, arg1 *products.RequestGetAllProduct) ([]table.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProduct", arg0, arg1)
	ret0, _ := ret[0].([]table.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProduct indicates an expected call of GetAllProduct.
func (mr *MockProductRepositoryMockRecorder) GetAllProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProduct", reflect.TypeOf((*MockProductRepository)(nil).GetAllProduct), arg0, arg1)
}

// GetProductById mocks base method.
func (m *MockProductRepository) GetProductById(arg0 context.Context, arg1 *products.Request) (*table.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductById", arg0, arg1)
	ret0, _ := ret[0].(*table.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductById indicates an expected call of GetProductById.
func (mr *MockProductRepositoryMockRecorder) GetProductById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductById", reflect.TypeOf((*MockProductRepository)(nil).GetProductById), arg0, arg1)
}

// InsertProduct mocks base method.
func (m *MockProductRepository) InsertProduct(arg0 context.Context, arg1 *products.Request) (*table.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertProduct", arg0, arg1)
	ret0, _ := ret[0].(*table.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertProduct indicates an expected call of InsertProduct.
func (mr *MockProductRepositoryMockRecorder) InsertProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertProduct", reflect.TypeOf((*MockProductRepository)(nil).InsertProduct), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockProductRepository) UpdateProduct(arg0 context.Context, arg1 *products.Request) (*table.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(*table.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductRepositoryMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductRepository)(nil).UpdateProduct), arg0, arg1)
}