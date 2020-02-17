package handler

import (
	"fmt"
	"net/http"
	"os"
)

// Run - run server
func Run(server *http.Server) {
	// make channel for errors
	serverErr := make(chan error, 1)
	go func() {
		fmt.Println("listening on ", server.Addr)
		err := server.ListenAndServe()
		serverErr <- err
	}()

	for {
		select {
		case err := <-serverErr:
			fmt.Println("there was an error the server crashed %w", err)
			os.Exit(1)
		}

	}
}
