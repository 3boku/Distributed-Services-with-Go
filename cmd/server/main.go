package main

import (
	"fmt"
	"github.com/3boku/proglog/internal/server"
	"log"
)

func main() {
	fmt.Println("Server On")
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
