package database

import (
	"database/sql"
	"fmt"
	"glow/shared"

	"errors"
	_ "github.com/go-sql-driver/mysql"
)

func findOrderId(userId int) int {
	var order_id int
	status := "cart"
	query := "select order_id from Orders where user_id = ? and status = ? "
	err := Db.QueryRow(query, userId, status).Scan(&order_id)
	if err == sql.ErrNoRows {
		fmt.Println("order id not found")
		return 0
	}
	return order_id

}

func AddThisToCart(pid string, price int, quantity int, userId int) {
	status := "cart"
	price = quantity * price
	order_id := findOrderId(userId)
	if order_id != 0 {
		query := "insert into OrderItems values(?,?,?,?)"
		_, err := Db.Exec(query, order_id, pid, quantity, price)
		if err != nil {
			fmt.Println("Faliure to insert to the cart order items")
		} else {
			fmt.Println("inserted cart to Order items")
		}
	} else {
		query := "insert into Orders(user_id,status) values(?,?)"
		_, err := Db.Exec(query, userId, status)
		if err != nil {
			fmt.Println("Failure to add first cart to orders")
		} else {
			fmt.Println("successfully inserted the record to the orders")
			AddThisToCart(pid, price, quantity, userId)
		}
	}
}

func RemoveFromCart(pid string, userId int) {
	orderId := findOrderId(userId)
	if orderId != 0 {
		query := "delete from OrderItems where order_id = ? and pid = ? "
		_, err := Db.Exec(query, orderId, pid)
		if err != nil {
			fmt.Println("deletion of cart is failure ")
		} else {
			fmt.Println("cart deletion successfull")
		}
	} else {
		fmt.Println("Order id not found to delete")
	}

}

func findProductByPid(pid string) (string, string, string) {
	var productName, poster, product_type_id string
	query := "select name,poster,product_type_id from Products where pid = ? "
	row := Db.QueryRow(query, pid).Scan(&productName, &poster, &product_type_id)
	if row == sql.ErrNoRows {
		fmt.Println("no product id found")
		return "", "", ""
	}
	return productName, poster, product_type_id
}

func DisplayCart(userID int) ([]shared.CartProducts, error) {
	var CartProducts []shared.CartProducts
	order_id := findOrderId(userID)
	if order_id == 0 {
		return nil, errors.New("no active cart found for this user")
	}
	query := "select pid,quantity,price from OrderItems where order_id = ? "
	rows, err := Db.Query(query, order_id)
	if err != nil {
		fmt.Println("this is the problem from db")
		return nil, fmt.Errorf("failed to query cart: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var carproduct shared.CartProducts
		rows.Scan(
			&carproduct.Pid,
			&carproduct.Quantity,
			&carproduct.Price,
		)
		name, poster, product_type_id := findProductByPid(carproduct.Pid)
		carproduct.Name = name
		fmt.Println("carproduct from displayCart - ", carproduct)
		carproduct.Poster = shared.FindPosterFromPid(product_type_id, poster) // find poster from minio
		CartProducts = append(CartProducts, carproduct)
	}
	return CartProducts, nil

}
