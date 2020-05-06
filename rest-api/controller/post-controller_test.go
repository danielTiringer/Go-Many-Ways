package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	entity "github.com/danielTiringer/Go-Many-Ways/rest-api/entity"
	repository "github.com/danielTiringer/Go-Many-Ways/rest-api/repository"
	service "github.com/danielTiringer/Go-Many-Ways/rest-api/service"
)

const (
	TITLE string = "test title"
	TEXT  string = "test text"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postSrv        service.PostService       = service.NewPostService(postRepository)
	postController PostController            = NewPostController(postSrv)
)

func TestGetPosts(t *testing.T) {

}

func TestAddPost(t *testing.T) {
	var jsonReq = []byte(`{"title": "` + TITLE + `", "text": "` + TEXT + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonReq))

	handler := http.HandlerFunc(postController.AddPost)

	response := httptest.NewRecorder()

	handler.ServeHTTP(response, req)

	status := response.Code

	if status != http.StatusCreated {
		t.Errorf("Handler returned incorrect status code: got %v, want %v", status, http.StatusCreated)
	}

	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)
}
