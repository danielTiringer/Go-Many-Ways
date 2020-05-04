package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/danielTiringer/Go-Many-Ways/rest-api/entity"
	"github.com/danielTiringer/Go-Many-Ways/rest-api/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Error getting the posts." }`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "error": "Error unmashalling the request." }`))
		return
	}

	post.ID = rand.Int63()
	repo.Save(&post)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}
