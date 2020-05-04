package main

import (
	"fmt"
	"net/http"
	"os"

	router "github.com/danielTiringer/Go-Many-Ways/rest-api/http"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	// httpRouter.GET("/posts", getPosts)
	// httpRouter.POST("/posts", addPost)

	httpRouter.SERVE(":" + os.Getenv("PORT"))
}
