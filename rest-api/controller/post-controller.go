package controller

import (
	"encoding/json"
	"net/http"

	"github.com/danielTiringer/Go-Many-Ways/rest-api/entity"
	"github.com/danielTiringer/Go-Many-Ways/rest-api/errors"
	"github.com/danielTiringer/Go-Many-Ways/rest-api/service"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(http.ResponseWriter, *http.Request)
	AddPost(http.ResponseWriter, *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error getting the posts."})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (*controller) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error unmarshalling the request."})
		return
	}

	valError := postService.Validate(&post)
	if valError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: valError.Error()})
		return
	}

	result, postError := postService.Create(&post)
	if postError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error saving the post."})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
