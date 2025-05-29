package database

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"glow/shared"
)

func GetProductsDb(product_type_id string) ([]shared.Product, error) {

	fmt.Println("GetProductsDb working from database folder")

	var AllProducts []shared.Product

	query := "select * from Products where product_type_id = ?"
	rows, _ := Db.Query(query, product_type_id)

	if !rows.Next() {
		return nil, errors.New("no products found")
	} else {
		var err error
		for rows.Next() {
			var product shared.Product
			err = rows.Scan(
				&product.Product_type_id,
				&product.Pid,
				&product.Poster,
				&product.Name,
				&product.SKU,
				&product.Price,
				&product.Quantity,
				&product.Brand,
				&product.Category,
				&product.Color,
				&product.Material,
				&product.Weight,
				&product.Size,
				&product.OriginalPrice,
				&product.Sale,
				&product.Discount,
			)
			if err != nil {
				fmt.Println("error when fetching products from db")
			}
			AllProducts = append(AllProducts, product)
		}

	}
	return AllProducts, nil

}
