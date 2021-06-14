package main

import (
	"fmt"

	"jwt-go/src/driver"
	"jwt-go/src/handler"

	"net/http"
)

func main() {
	fmt.Println("hello guy...")
	driver.ConnectMongoDB()


	http.HandleFunc("/login",handler.Login)
	http.HandleFunc("/register",handler.Register)
	fmt.Println("Server running [:8000]")
	http.ListenAndServe(":8000", nil)
}
