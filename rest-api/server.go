package main

import (
	"fmt"
	"net/http"
	"os"

	controller "github.com/danielTiringer/Go-Many-Ways/rest-api/controller"
	router "github.com/danielTiringer/Go-Many-Ways/rest-api/http"
)

var (
	httpRouter     router.Router             = router.NewChiRouter()
	postController controller.PostController = controller.NewPostController()
)

func main() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(":" + os.Getenv("PORT"))
}
