package repo

// server is a struct that contains all the repositories
type server struct {
	StockRepo *StockRepo
}


func NewServer(stockRepo *StockRepo) *server {
	return &server{
		StockRepo: stockRepo,
	}
}
