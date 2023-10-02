package stock_handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"stock-api/repo"

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
	if page == 1 && pageSize == 10 {
		return TempStockList, nil
	}
	return nil, errors.New("Mock database error")
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
)

func TestGetStocks(t *testing.T) {
	// Create a Gin router with the handler function
	r := gin.Default()
	r.GET("/api/stocks", GetStocks)

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
	assert.Len(t, response, 20) // Assuming 20 items per page for this test
	// Add more assertions as needed
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
