package main

import (
	"fmt"
	"go-webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting Server ...")
	// do important stuff
	fmt.Println("Server started at port", port)
	fmt.Println("Number of retries", numberOfRetries)
	fmt.Println("Server Started!")
	return port, nil
}
