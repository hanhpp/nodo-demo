package stock_handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"stock-api/repo"
	"stock-api/util"

	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// DB interface defines the methods for interacting with the database.
type DB interface {
	GetPaginatedStocks(page, pageSize int) ([]repo.Stock, error)
}
type MockDB struct {
	mock.Mock
}

func (m *MockDB) GetPaginatedStocks(page, pageSize int) ([]repo.Stock, error) {
	// return predefined data
	if page == 2 && pageSize == 20 {
		return TempStockList, nil
	}
	return nil, errors.New("Error retrieving stocks")
}

func (m *MockDB) CreateStock(stock *repo.Stock) error {
	if stock.ID > 0 {
		return nil
	}
	return errors.New("Error creating stock")
}

func (m *MockDB) GetStockByID(id uint) (*repo.Stock, error) {
	if id == 1 {
		return &TempStockList[0], nil
	}
	return nil, errors.New("Error retrieving stock")
}

func (m *MockDB) UpdateStock(stock *repo.Stock) error {
	if stock.ID > 0 {
		return nil
	}
	return errors.New("Error updating stock")
}

func (m *MockDB) DeleteStock(id uint) error {
	if id > 100 {
		return errors.New("Stock not found")
	}
	if id > 0 {
		return nil
	}
	return errors.New("Error deleting stock")
}

var (
	now           = time.Now()
	TempStockList = []repo.Stock{
		{
			ID:           1,
			Name:         "Apple",
			CurrentPrice: 10.0,
			LastUpdate:   now,
		},
		{
			ID:           2,
			Name:         "Google",
			CurrentPrice: 20.0,
			LastUpdate:   now,
		},
		{
			ID:           3,
			Name:         "Microsoft",
			CurrentPrice: 30.0,
			LastUpdate:   now,
		},
	}
	tempStockLen = len(TempStockList)
	format = "2006-01-02 15:04:05"
)
func GetStocksWithMockDB(c *gin.Context) {
	// Get query parameters for pagination
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	// Convert query parameters to integers
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BadRequestResponse)
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BadRequestResponseCustom("Invalid page size"))
		return
	}

	// Retrieve paginated stocks from the repository
	m := new(MockDB)
	stocks, err := m.GetPaginatedStocks(pageInt, pageSizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.InternalServerErrorResponse)
		return
	}

	c.JSON(http.StatusOK, stocks)
}

func TestGetStocks(t *testing.T) {
	// Create a Gin router with the handler function
	r := gin.Default()

	r.GET("/api/stocks", GetStocksWithMockDB)

	// Create a mock HTTP request with query parameters
	req, err := http.NewRequest("GET", "/api/stocks?page=2&pageSize=20", nil)
	assert.NoError(t, err)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the request to the Gin router
	r.ServeHTTP(w, req)

	// Check the HTTP response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the JSON response body
	var response []repo.Stock // Replace with the actual stock struct type
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Perform assertions on the response data
	assert.Len(t, response, tempStockLen)
}

func TestGetStocksBadRequest(t *testing.T) {

	// Create a Gin router with the handler function
	r := gin.Default()
	r.GET("/api/stocks", GetStocks)

	// Create a mock HTTP request with invalid query parameters
	req, err := http.NewRequest("GET", "/api/stocks?page=invalid&pageSize=20", nil)
	assert.NoError(t, err)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the request to the Gin router
	r.ServeHTTP(w, req)

	// Check the HTTP response status code for a bad request
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Perform assertions on the response body or error message
	// Add more assertions as needed
}

func CreateStockWithMockDB(c *gin.Context) {
	var stock repo.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a mock database
	m := new(MockDB)
	// Create a new stock in the mock database
	if err := m.CreateStock(&stock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock"})
		return
	}

	c.JSON(http.StatusCreated, stock)
}





func TestCreateStock(t *testing.T) {
	// Create a Gin router with the handler function
	r := gin.Default()
	r.POST("/api/stocks", CreateStockWithMockDB)

	// Create a sample stock to be sent in the request body
	stock := repo.Stock{
		ID:           4,
		Name:         "Amazon",
		CurrentPrice: 40.0,
		LastUpdate:   now,
	}
	stockJSON, _ := json.Marshal(stock)

	// Create a mock HTTP request with the stock JSON in the request body
	req, err := http.NewRequest("POST", "/api/stocks", bytes.NewBuffer(stockJSON))
	assert.NoError(t, err)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the request to the Gin router
	r.ServeHTTP(w, req)

	// Check the HTTP response status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Parse the JSON response body
	var response repo.Stock
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Perform assertions on the response data
	assert.Equal(t, stock.ID, response.ID)
	assert.Equal(t, stock.Name, response.Name)
	assert.Equal(t, stock.CurrentPrice, response.CurrentPrice)
	assert.Equal(t, stock.LastUpdate.UTC().Format(format), response.LastUpdate.UTC().Format(format))
}

func GetStockByIDWithMockDB(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BadRequestResponse)
		return
	}

	// Create a mock database
	m := new(MockDB)
	// Retrieve a single stock from the mock database
	stock, err := m.GetStockByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, *stock)

}

func TestGetStockByID(t *testing.T) {
	// Create a Gin router with the handler function
	r := gin.Default()
	r.GET("/api/stocks/:id", GetStockByIDWithMockDB)

	// Create a mock HTTP request with a valid stock ID
	id := "1"
	req, err := http.NewRequest("GET", "/api/stocks/"+id, nil)
	assert.NoError(t, err)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the request to the Gin router
	r.ServeHTTP(w, req)

	// Check the HTTP response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the JSON response body
	var response repo.Stock
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Perform assertions on the response data
	assert.Equal(t, uint(1), response.ID)
	// Add more assertions as needed
}

func UpdateStockWithMockDB(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BadRequestResponse)
		return
	}

	var updatedStock repo.Stock
	if err := c.ShouldBindJSON(&updatedStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a mock database
	m := new(MockDB)
	// Update the stock in the mock database
	updatedStock.ID = uint(id)
	if err := m.UpdateStock(&updatedStock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
		return
	}

	c.JSON(http.StatusOK, updatedStock)
}

func TestUpdateStock(t *testing.T) {
	// Create a Gin router with the handler function
	r := gin.Default()
	r.PATCH("/api/stocks/:id", UpdateStockWithMockDB)

	// Create a sample updated stock to be sent in the request body
	updatedStock := repo.Stock{
		ID:           1,
		Name:         "Updated Apple",
		CurrentPrice: 15.0,
		LastUpdate:   now,
	}
	updatedStockJSON, _ := json.Marshal(updatedStock)

	// Create a mock HTTP request with the updated stock JSON in the request body
	id := "1"
	req, err := http.NewRequest("PATCH", "/api/stocks/"+id, bytes.NewBuffer(updatedStockJSON))
	assert.NoError(t, err)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the request to the Gin router
	r.ServeHTTP(w, req)

	// Check the HTTP response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse the JSON response body
	var response repo.Stock
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Perform assertions on the response data
	assert.Equal(t, updatedStock.ID, response.ID)
	assert.Equal(t, updatedStock.Name, response.Name)
	assert.Equal(t, updatedStock.CurrentPrice, response.CurrentPrice)
	assert.Equal(t, updatedStock.LastUpdate.Format(format), response.LastUpdate.Format(format))
}

func DeleteStockWithMockDB(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, util.BadRequestResponse)
		return
	}

	// Create a mock database
	m := new(MockDB)
	// Delete the stock from the mock database
	if err := m.DeleteStock(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete stock : " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock deleted"})
}

func TestDeleteStock(t *testing.T) {
	// Create a Gin router with the handler function
	r := gin.Default()
	r.DELETE("/api/stocks/:id", DeleteStockWithMockDB)

	// Create a mock HTTP request with a valid stock ID
	id := "1"
	req, err := http.NewRequest("DELETE", "/api/stocks/"+id, nil)
	assert.NoError(t, err)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()

	// Serve the request to the Gin router
	r.ServeHTTP(w, req)

	// Check the HTTP response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Perform assertions on the response data or error message
	// Add more assertions as needed
}

func TestDeleteStockNotFound(t *testing.T) {
	// Create a Gin router with the handler function
	r := gin.Default()
	r.DELETE("/api/stocks/:id", DeleteStockWithMockDB)

	// Create a mock HTTP request with a stock ID that doesn't exist
	id := "1000" // An ID that doesn't exist in TempStockList
	req, err := http.NewRequest("DELETE", "/api/stocks/"+id, nil)
	assert.NoError(t, err)

	// Create a mock HTTP response recorder
	w := httptest.NewRecorder()
	// Serve the request to the Gin router
	r.ServeHTTP(w, req)

	// Check the HTTP response status code for a not found error
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Parse the JSON response body
	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Perform assertions on the response data
	expectedError := "Failed to delete stock : Stock not found"
	actualError, ok := response["error"].(string)
	assert.True(t, ok)
	assert.Equal(t, expectedError, actualError)
	// Add more assertions as needed
}