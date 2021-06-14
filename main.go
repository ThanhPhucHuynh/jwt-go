package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"jwt-go/src/driver"
	"jwt-go/src/handler"

	"net/http"
)

func main() {
	fmt.Println("hello guy...")
	driver.ConnectMongoDB()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Running Hello Handler")

		// read the body
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		// write the response
		fmt.Fprintf(rw, "Hello %s", b)
	})

	http.HandleFunc("/login",handler.Login)
	http.HandleFunc("/register",handler.Register)
	http.HandleFunc("/user", handler.GetUser)

	fmt.Println("Server running [:8000]")
	http.ListenAndServe(":8000", nil)
}
