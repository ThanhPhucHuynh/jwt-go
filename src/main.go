package main

import (
	"fmt"
	"jwt-go/src/driver"
)

func main() {
	fmt.Println("hello guy...")
	driver.ConnectMongoDB()
}