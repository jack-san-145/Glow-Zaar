package database

import (
	"fmt"
	"net/smtp"
)

func SendOrderConfirmationEmail(toEmail, userName string, product string, quantity int, price int, delivery string) error {
	from := "glowzaar145@gmail.com" //  app's Gmail
	password := "csqsaylkszluhdog"  // Gmail password

	// Gmail SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message body
	subject := "Order Confirmation - Glow Zaar"
	body := fmt.Sprintf("Hello %s,\n\nYou have ordered - %s -> %dPcs (₹%d)\nDelivery Date - %s \n\nThank you for your order! We are processing it now.\n\n- Glow Zaar Team", userName, product, quantity, price, delivery)

	msg := []byte("Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, msg)
	if err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
	return err
}

func FindEmail(userId int, pid string) (string, string, string) {
	var email, name, productName string
	query := "select user_name,user_email from User where user_id = ? "
	err := Db.QueryRow(query, userId).Scan(&name, &email)
	if err != nil {
		fmt.Println("failed to find the user email")
		return "", "", ""
	}
	query = "select name from Products where pid = ? "
	err = Db.QueryRow(query, pid).Scan(&productName)
	if err != nil {
		fmt.Println("failed to find the order products")
		return "", "", ""
	}
	fmt.Println("name, email, productName", name, email, productName)
	return name, email, productName
}
