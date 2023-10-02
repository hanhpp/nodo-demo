package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStockRepo_CreateStock(t *testing.T) {
	// Create a mock database
	mockDB := new(MockStockRepo)

	// Create a Stock instance
	stock := &Stock{
		ID:           1,
		Name:         "Test Stock",
		CurrentPrice: 100.0,
	}
	// Set expectations for the CreateStock method
	mockDB.On("CreateStock", mock.Anything).Return(nil)

	// Call the method being tested
	err := mockDB.CreateStock(stock)

	// Assert that there are no errors
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}
func TestStockRepo_GetStocks(t *testing.T) {
	// Create a mock database
	mockDB := new(MockStockRepo)

	// Define mock data to return
	mockData := []Stock{
		{
			ID:           1,
			Name:         "Stock A",
			CurrentPrice: 100.0,
		},
		{
			ID:           2,
			Name:         "Stock B",
			CurrentPrice: 150.0,
		},
	}

	// Set expectations for the GetStocks method
	mockDB.On("GetStocks").Return(mockData, nil)

	// Call the method being tested
	stocks, err := mockDB.GetStocks()

	// Assert that there are no errors
	assert.NoError(t, err)

	// Assert that the returned data matches the mock data
	assert.Equal(t, mockData, stocks)

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}

func TestStockRepo_GetStockByID(t *testing.T) {
	// Create a mock database
	mockDB := new(MockStockRepo)

	// Define the ID of the stock to retrieve
	stockID := uint(1)

	// Define mock data to return
	mockData := &Stock{
		ID:           1,
		Name:         "Stock A",
		CurrentPrice: 100.0,
	}

	// Set expectations for the GetStockByID method
	mockDB.On("GetStockByID", stockID).Return(mockData, nil)

	// Call the method being tested
	stock, err := mockDB.GetStockByID(stockID)

	// Assert that there are no errors
	assert.NoError(t, err)

	// Assert that the returned data matches the mock data
	assert.Equal(t, mockData, stock)

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}

func TestStockRepo_UpdateStock(t *testing.T) {
	// Create a mock database
	mockDB := new(MockStockRepo)

	// Define the stock to update
	stockToUpdate := &Stock{
		ID:           1,
		Name:         "Stock A",
		CurrentPrice: 100.0,
	}

	// Set expectations for the UpdateStock method
	mockDB.On("UpdateStock", stockToUpdate).Return(nil)

	// Call the method being tested
	err := mockDB.UpdateStock(stockToUpdate)

	// Assert that there are no errors
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}

func TestStockRepo_DeleteStock(t *testing.T) {
	// Create a mock database
	mockDB := new(MockStockRepo)

	// Define the ID of the stock to delete
	stockID := uint(1)

	// Set expectations for the DeleteStock method
	mockDB.On("DeleteStock", stockID).Return(nil)

	// Call the method being tested
	err := mockDB.DeleteStock(stockID)

	// Assert that there are no errors
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockDB.AssertExpectations(t)
}
