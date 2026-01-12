package main

import (
	"fmt"
	"library/http"
	"library/lib"
)

func main() {
	lib := lib.NewLib()

	handle := http.NewHTTPHAndlers(lib)
	server := http.NewHTTPServer(handle)

	if err := server.StartServer(); err != nil {
		fmt.Println("failed to start server")
	}
}
