package repo

import (
	"time"
)

// Stock represents the stock entity.
type Stock struct {
	ID           uint      `gorm:"primarykey"`
	Name         string    `json:"name"`
	CurrentPrice float64   `json:"currentPrice"`
	LastUpdate   time.Time `json:"lastUpdate"`
}

// StockRepository represents the repository containing GORM instance.
type StockRepo struct {
	db *Database
}

// NewStockRepository initializes a new StockRepository with a GORM instance.
func NewStockRepo(db *Database) *StockRepo {
	return &StockRepo{db}
}

// CreateStock inserts a new stock into the database.
func (s *StockRepo) CreateStock(stock *Stock) error {
	result := s.db.db.Create(stock)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetStocks retrieves a list of stocks from the database.
func (s *StockRepo) GetStocks() ([]Stock, error) {
	var stocks []Stock
	result := s.db.db.Find(&stocks)
	if result.Error != nil {
		return nil, result.Error
	}
	return stocks, nil
}

// GetStockByID retrieves a single stock by ID from the database.
func (repo *StockRepo) GetStockByID(id uint) (*Stock, error) {
	var stock Stock
	result := repo.db.db.First(&stock, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &stock, nil
}

// UpdateStock updates the price of a single stock in the database.
func (repo *StockRepo) UpdateStock(stock *Stock) error {
	result := repo.db.db.Save(stock)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteStock deletes a single stock from the database.
func (repo *StockRepo) DeleteStock(id uint) error {
	result := repo.db.db.Delete(&Stock{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetPaginatedStocks retrieves a paginated list of stocks from the database.
func (repo *StockRepo) GetPaginatedStocks(page, pageSize int) ([]Stock, error) {
	var stocks []Stock
	offset := (page - 1) * pageSize

	result := repo.db.db.Offset(offset).Limit(pageSize).Find(&stocks)
	if result.Error != nil {
		return nil, result.Error
	}
	return stocks, nil
}
