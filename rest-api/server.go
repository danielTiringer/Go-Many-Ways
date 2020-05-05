package main

import (
	"fmt"
	"net/http"
	"os"

	controller "github.com/danielTiringer/Go-Many-Ways/rest-api/controller"
	router "github.com/danielTiringer/Go-Many-Ways/rest-api/http"
	repository "github.com/danielTiringer/Go-Many-Ways/rest-api/repository"
	service "github.com/danielTiringer/Go-Many-Ways/rest-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(":" + os.Getenv("PORT"))
}
