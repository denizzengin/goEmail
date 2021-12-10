package main

import (
	"crypto/tls"
	"fmt"
	"os"

	template "github.com/denizzengin/goTemplate"
	"gopkg.in/gomail.v2"
)

func main() {
	body := getTemplate()
	from := os.Getenv("EMAIL_ADDRESS")
	to := os.Getenv("EMAIL_ADDRESS")
	sendEmail(from, to, "", "Dummy Subject", body)
}

func sendEmail(from, to, cc, subject, body string) {	
	e := gomail.NewMessage(func(es *gomail.Message) {
		es.SetHeader("From", from)
		es.SetHeader("To", to)
		if cc != "" {
			es.SetAddressHeader("Cc", cc, "")
		}
		es.SetHeader("Subject", subject)
		es.SetBody("text/html", body)		
	})

	// Send with gmail
	// if you get error then it may be required to create and email app password and use it
	dial := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_ADDRESS"), os.Getenv("EMAIL_APP_PASSWORD"))
	dial.TLSConfig = &tls.Config{InsecureSkipVerify: true} // pass some cert error
	if err := dial.DialAndSend(e); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Email sent successfully...")
}

func getTemplate() string {
	template.GetAllTemplate()
	orders := []template.Order{
		{OrderId: 1, OrderDesc: "iphone 5s", CustomerId: 100, CustomerName: "John"},
		{OrderId: 2, OrderDesc: "iphone 6", CustomerId: 101, CustomerName: "Alice"},
		{OrderId: 3, OrderDesc: "iphone 13Pro", CustomerId: 102, CustomerName: "Sarı Cizmeli"},
	}
	wo := template.WrappedOrder{Description: "You can find all orders gave in last two days.", Orders: orders}
	t := template.ParseTemplate(template.EmailTemplate, wo)
	return t
}
