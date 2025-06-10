package database

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"glow/shared"
	"time"
)

func insertBuyToOrderItems(pid string, quantity int, price int, order_id int) {
	query := "insert into OrderItems(order_id,pid,quantity,price) values(?,?,?,?)"
	_, err := Db.Exec(query, order_id, pid, quantity, price)
	if err != nil {
		fmt.Println("Faliure to insert to the myorders  Orderitems")
	} else {
		fmt.Println("Myorders inserted to Order items")
	}
}

func BuyItNowDb(pid string, quantity int, price int, userId int) {
	status := "ordered"
	now := time.Now()
	formattedNow := now.Format("2006-01-02 15:04:05")
	fmt.Println("formattedNow - ", formattedNow)
	query := "insert into Orders(user_id,status,ordered_at) values(?,?,?)"
	_, err := Db.Exec(query, userId, status, formattedNow)
	if err != nil {
		fmt.Println("Failure to add buy it now orders")
	} else {
		var order_id int
		query := "select order_id from Orders where ordered_at = ?"
		err := Db.QueryRow(query, formattedNow).Scan(&order_id)
		if err != nil {
			fmt.Println("Failure to find order_id by now time ")
		}
		fmt.Println("successfully inserted the record(buy it now) to the orders")
		insertBuyToOrderItems(pid, quantity, price, order_id)
		name, email, product := FindEmail(userId, pid)
		err = SendOrderConfirmationEmail(email, name, product, quantity, price, findDeliveryDate(findOrderedTime(order_id)))
		if err != nil {
			fmt.Println("email not sending....")
			return
		}
		fmt.Printf("email sent to %v", email)
	}

}

func findOrderedTime(order_id int) string {
	var orderedTime string
	query := "select ordered_at from Orders where order_id = ? and status = ?"
	err := Db.QueryRow(query, order_id, "ordered").Scan(&orderedTime)
	if err != nil {
		fmt.Println("error while finding orderedTime")
		return ""
	}
	fmt.Println("order_id for time - ", order_id)
	fmt.Println("ordered time - ", orderedTime)

	return orderedTime

}
func findDeliveryDate(orderedDate string) string {
	layout := "2006-01-02 15:04:05" //reference layout
	orderDate, err := time.Parse(layout, orderedDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}
	deliveryDate := orderDate.AddDate(0, 0, 7)
	formattedDelivery := deliveryDate.Format("15 January 2006")

	return formattedDelivery
}

func DisplayMyOrders(userID int) ([]shared.OrderedProducts, error) {
	var OrderedProducts []shared.OrderedProducts
	query := "select order_id from Orders where user_id = ? and status = ? order by ordered_at desc "
	order_id_rows, err := Db.Query(query, userID, "ordered")
	if err != nil {
		fmt.Println("order id not found")
		return nil, nil
	}
	for order_id_rows.Next() {
		var order_id int
		order_id_rows.Scan(&order_id)
		if order_id == 0 {
			return nil, errors.New("no orders found for this user")
		}
		query := "select pid,quantity,price from OrderItems where order_id = ? "
		rows, err := Db.Query(query, order_id)
		if err != nil {
			fmt.Println("this is the problem from db(display my orders )")
			return nil, fmt.Errorf("failed to query orders: %v", err)
		}
		defer rows.Close()
		for rows.Next() {
			var orderedproduct shared.OrderedProducts
			rows.Scan(
				&orderedproduct.Pid,
				&orderedproduct.Quantity,
				&orderedproduct.Price,
			)
			name, poster, product_type_id := findProductByPid(orderedproduct.Pid)
			orderedproduct.Name = name
			fmt.Println("ordered products from displayorders - ", orderedproduct)
			orderedproduct.Poster = shared.FindPosterFromPid(product_type_id, poster) // find poster from minio
			orderedproduct.OrderedDate = findOrderedTime(order_id)
			orderedproduct.DeliveryDate = findDeliveryDate(orderedproduct.OrderedDate)
			OrderedProducts = append(OrderedProducts, orderedproduct)
		}
	}
	return OrderedProducts, nil

}
