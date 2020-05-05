package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/danielTiringer/Go-Many-Ways/rest-api/entity"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty.", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "Test text"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty.", err.Error())
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier int64 = 1
	post := entity.Post{ID: identifier, Title: "Test title", Text: "Test text"}
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "Test title", result[0].Title)
	assert.Equal(t, "Test text", result[0].Text)
	assert.Nil(t, err)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	post := entity.Post{Title: "Test title", Text: "Test text"}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "Test title", result.Title)
	assert.Equal(t, "Test text", result.Text)
	assert.Nil(t, err)
}
