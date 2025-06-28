package main

import (
	"fmt"
	"log"

	"net/http"

	"github.com/SitaGomes/coins-exchange/internal/user"
)

func main() {

	userController := user.UserController{}
	userController.Listen()

	fmt.Println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
