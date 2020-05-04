package repository

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/danielTiringer/Go-Many-Ways/rest-api/entity"
)

type repo struct{}

func NewFirestoreRepository() PostRepository {
	return &repo{}
}

var (
	projectId      string = os.Getenv("FIRESTORE_PROJECT_ID")
	collectionName string = os.Getenv("FIRESTORE_COLLECTION_NAME")
)

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})
	if err != nil {
		log.Fatalf("Failed to save the new post: %v", err)
		return nil, err
	}

	return post, nil
}
