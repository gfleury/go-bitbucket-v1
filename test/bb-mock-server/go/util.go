package swagger

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	handlerName := "UNKNOWN"
	if route := mux.CurrentRoute(r); route != nil {
		routeName := route.GetName()
		if routeName != "" {
			handlerName = routeName
		}
	}
	fileName := fmt.Sprintf("test/bb-mock-server/mocked_responses/%s.json", handlerName)
	response, err := os.ReadFile(fileName)
	if err == nil {
		w.Write(response)
	} else {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Write([]byte("{}"))
		file.Close()
	}
}
