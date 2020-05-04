package repository

import (
	"github.com/danielTiringer/Go-Many-Ways/rest-api/entity"
)

type PostRepository interface {
	FindAll() ([]entity.Post, error)
	Save(post *entity.Post) (*entity.Post, error)
}
