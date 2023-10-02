package stock_handler

import (
	"net/http"
	"stock-api/repo"
	"stock-api/util"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @title Stock API
// @description This is a sample stock API.
// @version 1
// @host localhost:8080
// @BasePath /api

// @Summary Get a list of stocks
// @Description Retrieves a list of stocks with pagination.
// @Accept json
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param pageSize query int false "Number of stocks per page (default is 10)"
// @Success 200 {array} repo.Stock
// @Failure 400 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /api/stocks [get]
func GetStocks(c *gin.Context) {
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
	stocks, err := repo.Server.StockRepo.GetPaginatedStocks(pageInt, pageSizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.InternalServerErrorResponse)
		return
	}

	c.JSON(http.StatusOK, stocks)
}

// @Summary Create a new stock
// @Description Creates a new stock.
// @Accept json
// @Produce json
// @Param stock body repo.Stock true "Stock object to create"
// @Success 201 {object} repo.Stock
// @Failure 400 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /api/stocks [post]
func CreateStock(c *gin.Context) {
	var stock repo.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := repo.Server.StockRepo.CreateStock(&stock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock"})
		return
	}

	c.JSON(http.StatusCreated, stock)
}

// @Summary Get a stock by ID
// @Description Retrieves a single stock by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Stock ID"
// @Success 200 {object} repo.Stock
// @Failure 400 {object} util.ErrorResponse
// @Failure 404 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /api/stocks/{id} [get]
func GetStockByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	stock, err := repo.Server.StockRepo.GetStockByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, *stock)
}

// @Summary Update a stock's price
// @Description Updates the price of a single stock.
// @Accept json
// @Produce json
// @Param id path int true "Stock ID"
// @Param updatedStock body repo.Stock true "Updated Stock object"
// @Success 200 {object} repo.Stock
// @Failure 400 {object} util.ErrorResponse
// @Failure 404 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /api/stocks/{id} [patch]
func UpdateStock(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	var updatedStock repo.Stock
	if err := c.ShouldBindJSON(&updatedStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the stock with the given ID exists
	_, err = repo.Server.StockRepo.GetStockByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	updatedStock.ID = uint(id)

	if err := repo.Server.StockRepo.UpdateStock(&updatedStock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
		return
	}

	c.JSON(http.StatusOK, updatedStock)
}

// @Summary Delete a stock by ID
// @Description Deletes a single stock by its ID.
// @Accept json
// @Produce json
// @Param id path int true "Stock ID"
// @Success 200 {object} util.Response
// @Failure 400 {object} util.ErrorResponse
// @Failure 404 {object} util.ErrorResponse
// @Failure 500 {object} util.ErrorResponse
// @Router /api/stocks/{id} [delete]
func DeleteStock(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	// Check if the stock with the given ID exists
	_, err = repo.Server.StockRepo.GetStockByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	if err := repo.Server.StockRepo.DeleteStock(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete stock"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Stock deleted successfully",
	})
}
