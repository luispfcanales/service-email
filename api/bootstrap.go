package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/luispfcanales/service-email/api/handlers"
)

// Run start server
func Run() error {
	http.HandleFunc("/send-email", handlers.SendEmail)
	return http.ListenAndServe(getPort(), nil)
}

func getPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = "3000"
	}
	return fmt.Sprintf(":%s", p)
}
