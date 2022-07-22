package main

import (
	"fmt"
	supervisorService "github.com/submaline/supervisor-service"
	"github.com/submaline/supervisor-service/gen/supervisor/v1/supervisorv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

func main() {
	supervisorServiceHandler := &supervisorService.SupervisorService{}

	mux := http.NewServeMux()
	mux.Handle(supervisorv1connect.NewSupervisorServiceHandler(
		supervisorServiceHandler))

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Service listening on %v", port)
	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
