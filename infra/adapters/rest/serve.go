package rest

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

func (s *restServer) Serve() {
	router := s.router()

	fmt.Println("Http  Server running at 8080...")
	if err := fasthttp.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
