package database

// var productList []Product // small letter for private variable

var ProductList []Product // capital letter for public variable

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

func init() {
	ProductList = append(ProductList, Product{
		ID:          1,
		Name:        "Product 1",
		Description: "Description of product 1",
		Price:       10.99,
		ImgURL:      "https://example.com/product1.jpg",
	}, Product{
		ID:          2,
		Name:        "Product 2",
		Description: "Description of product 2",
		Price:       19.99,
		ImgURL:      "https://example.com/product2.jpg",
	})
}

/*
	When the variable is named with small letter it is private,it will use in the same package, but if we want to access it from other package then we need to declare as capital letter
*/
