package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/danielTiringer/Go-Many-Ways/rest-api/entity"
)

var (
	dbHost   string = os.Getenv("MYSQL_HOST")
	dbPort   string = os.Getenv("MYSQL_PORT")
	dbName   string = os.Getenv("MYSQL_DATABASE")
	dbUser   string = os.Getenv("MYSQL_USER")
	dbPass   string = os.Getenv("MYSQL_PASSWORD")
	dbSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
)

type mysqlRepo struct{}

func NewMySQLRepository() PostRepository {
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sqlStmt := `
		CREATE table IF NOT EXISTS posts (id integer NOT NULL PRIMARY KEY, title text, txt text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
	return &mysqlRepo{}
}

func (*mysqlRepo) FindAll() ([]entity.Post, error) {
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var posts []entity.Post
	for rows.Next() {
		var id int64
		var title string
		var text string
		err = rows.Scan(&id, &title, &text)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		post := entity.Post{
			ID:    id,
			Title: title,
			Text:  text,
		}
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return posts, nil
}

func (*mysqlRepo) FindByID(id string) (*entity.Post, error) {
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM posts WHERE id = ?", id)

	var post entity.Post
	if row != nil {
		var id int64
		var title string
		var text string
		err := row.Scan(&id, &title, &text)
		if err != nil {
			return nil, err
		} else {
			post = entity.Post{
				ID:    id,
				Title: title,
				Text:  text,
			}
		}
	}

	return &post, nil
}

func (*mysqlRepo) Save(post *entity.Post) (*entity.Post, error) {
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	stmt, err := tx.Prepare("INSERT INTO posts(id, title, txt) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()
	_, err = stmt.Exec(post.ID, post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx.Commit()
	return post, nil
}

func (*mysqlRepo) Delete(post *entity.Post) error {
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal(err)
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := tx.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(post.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	tx.Commit()
	return nil
}
