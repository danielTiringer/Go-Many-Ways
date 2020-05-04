package main

import (
	"fmt"
	"net/http"
	"os"

	controller "github.com/danielTiringer/Go-Many-Ways/rest-api/controller"
	router "github.com/danielTiringer/Go-Many-Ways/rest-api/http"
	service "github.com/danielTiringer/Go-Many-Ways/rest-api/service"
)

var (
	postService    service.PostService       = service.NewPostService()
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(":" + os.Getenv("PORT"))
}
