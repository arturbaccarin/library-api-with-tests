package main

import "github.com/arturbaccarin/library-api-with-tests/handler"

func main() {

	router := SetupRouter(handler.Book{})

	router.Run()
}
