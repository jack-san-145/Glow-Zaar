package database

import (
	// "database/sql"
	"fmt"
	"glow/shared"

	_ "github.com/go-sql-driver/mysql"
)

func InsertProduct(productList []shared.Product) {
	query := " INSERT INTO Products (product_type_id, pid, poster, name, sku, price, quantity, brand,category, color, material, weight, size, original_price,sale, discount) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	for _, product := range productList {
		result, err := Db.Exec(query,
			product.Product_type_id,
			product.Pid,
			product.Poster,
			product.Name,
			product.SKU,
			product.Price,
			product.Quantity,
			product.Brand,
			product.Category,
			product.Color,
			product.Material,
			product.Weight,
			product.Size,
			product.OriginalPrice,
			product.Sale,
			product.Discount,
			product.IsProduct,
		)

		if err != nil {
			fmt.Println("something error while inserting products")
		} else {
			fmt.Println("successfully added - ", result)
		}
	}
}

func InsertHomeProduct(product_type []shared.Category) {
	query := "insert into Category(product_type,poster) values(?,?)"
	for _, EachProductType := range product_type {
		_, err := Db.Exec(query, EachProductType.Product_type_id, EachProductType.Poster)
		if err != nil {
			fmt.Println("error while inserting products category")
		}
	}
}
