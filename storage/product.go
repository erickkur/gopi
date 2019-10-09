package storage

// Product class for product table
type Product struct {
	ID   int
	SKU  string
	Name string
}

// GetProducts get all products
func GetProducts() []Product {
	var p []Product
	GopiDB.Find(&p)

	return p
}
