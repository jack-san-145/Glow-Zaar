package database

import (
	"database/sql"
	"fmt"
	"glow/shared"
)

func GetAllCategory() []shared.Category {
	var totalProductType []shared.Category
	query := "select * from Category"
	row, err := Db.Query(query)
	if err == sql.ErrNoRows {
		fmt.Println("no product type found in category table")
		return nil
	}
	for row.Next() {
		var HomeCard shared.Category
		err = row.Scan(
			&HomeCard.Product_type_id,
			&HomeCard.Poster,
			&HomeCard.IsProduct,
		)
		if err != nil {
			fmt.Println("error while accessing product_type")
		}
		totalProductType = append(totalProductType, HomeCard)
	}
	return totalProductType

}
