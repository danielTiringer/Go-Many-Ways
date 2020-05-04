package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	chiDispatcher = chi.NewRouter()
)

type chiRouter struct{}

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) SERVE(port string) {
	log.Println("Chi HTTP server running on port", port)
	log.Fatalln(http.ListenAndServe(port, chiDispatcher))
}
