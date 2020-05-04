package service

import (
	"errors"
	"math/rand"

	"github.com/danielTiringer/Go-Many-Ways/rest-api/entity"
	"github.com/danielTiringer/Go-Many-Ways/rest-api/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	FindAll() ([]entity.Post, error)
	Create(post *entity.Post) (*entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

func NewPostService() PostService {
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post is empty.")
		return err
	}

	if post.Title == "" {
		err := errors.New("The post title is empty.")
		return err
	}

	return nil
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}
