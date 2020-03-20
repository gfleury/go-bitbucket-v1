package swagger

import (
	"fmt"
	"io/ioutil"
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
		path, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println(path)
		}
	}
	fmt.Println(handlerName)
	fileName := fmt.Sprintf("mocked_responses/%s.json", handlerName)
	response, err := ioutil.ReadFile(fileName)
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
