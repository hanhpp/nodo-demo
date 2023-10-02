package repo

import (
	"github.com/stretchr/testify/mock"
)

type MockStockRepo struct {
	mock.Mock
}

func (m *MockStockRepo) CreateStock(stock *Stock) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *MockStockRepo) GetStocks() ([]Stock, error) {
	args := m.Called()
	return args.Get(0).([]Stock), args.Error(1)
}

func (m *MockStockRepo) GetStockByID(id uint) (*Stock, error) {
	args := m.Called(id)
	return args.Get(0).(*Stock), args.Error(1)
}

func (m *MockStockRepo) UpdateStock(stock *Stock) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *MockStockRepo) DeleteStock(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockStockRepo) GetPaginatedStocks(page, pageSize int) ([]Stock, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]Stock), args.Error(1)
}
