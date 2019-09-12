package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func InicializeRoutes() {
	router := NewRouter()

	err := http.ListenAndServe(":3001", router)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
