package swagger

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer(port int) error {
	log.Printf("Mock Server started")

	router := NewRouter()

	return http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}
