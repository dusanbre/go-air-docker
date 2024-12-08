package main

import (
	"dusanbre/test-1/cmd"
	"dusanbre/test-1/database"
	"dusanbre/test-1/middleware"
	"dusanbre/test-1/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	r := http.NewServeMux()

	router.LoadRoutes(r)

	stack := middleware.Chain(
		middleware.Loging,
	)

	db := database.Connect()

	cmd.InitCli(db)

	server := http.Server{
		Addr:    ":3000",
		Handler: stack(r),
	}

	fmt.Println("Server is running on port " + os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())
}
