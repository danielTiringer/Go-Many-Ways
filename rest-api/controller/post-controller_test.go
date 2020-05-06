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
	ID    int64  = 123
	TITLE string = "Title 1"
	TEXT  string = "Text 1"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postController PostController            = NewPostController(postSrv)
)

func setup() {
	var post entity.Post = entity.Post {
		ID: ID,
		Title: TITLE,
		Text: TEXT,
	}

	postRepo.Save(&post)
}

func tearDown(postID int64) {
	var post entity.Post = entity.Post{
		ID: postID,
	}
	postRepo.Delete(&post)
}

func TestAddPost(t *testing.T) {
	// Create a POST HTTP request
	var jsonStr = []byte(`{"title":"` + TITLE + `","text":"` + TEXT + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postController.AddPost)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v, want %v.",
			status, http.StatusCreated)
	}

	// Decode HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	// Clean up database
	tearDown(post.ID)
}

func TestGetPosts(t *testing.T) {
	// Generate a sample entry in the DB
	setup()

	// Create a GET HTTP request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postController.GetPosts)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v.",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// Assert HTTP response
	assert.NotNil(t, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	// Clean up database
	tearDown(ID)
}

func TestGetPostByID(t *testing.T) {
	// Insert new post
	setup()

	// Create new HTTP request
	req, _ := http.NewRequest("GET", "/posts/123", nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postController.GetPostByID)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v.",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	assert.Equal(t, ID, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	// Cleanup database
	tearDown(ID)
}
