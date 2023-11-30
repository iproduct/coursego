package mysql

import (
	"database/sql"
	"fmt"
	"github.com/iproduct/coursego/fmi-2023-05-my-blogs/model"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	opts   MySQLOptions
	client *sql.DB
}

type MySQLOptions struct {
	URI string
}

func New(opts MySQLOptions) *MySQLRepository {
	return &MySQLRepository{
		opts:   opts,
		client: nil,
	}
}

func (r *MySQLRepository) Init() error {
	var err error
	r.client, err = sql.Open("mysql", r.opts.URI)
	return err
}

func (r *MySQLRepository) GetAll() ([]model.Post, error) {
	if r.client == nil {
		return nil, fmt.Errorf("mysql repository is not initilized")
	}
	var posts []model.Post

	rows, err := r.client.Query("select id, title, created_at, author, content, likes from posts")
	if err != nil {
		return nil, fmt.Errorf("mysql query failure: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var result model.Post
		rows.Scan(&result.ID, &result.Title, &result.CreatedAt, &result.Author, &result.Content, &result.Likes)
		posts = append(posts, result)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating posts: %w", err)
	}
	return posts, nil
}

func (r *MySQLRepository) Insert(post *model.Post) error {
	if r.client == nil {
		return fmt.Errorf("mysql repository is not initilized")
	}
	_, err := r.client.Exec("INSERT INTO posts(id, title, author, content, likes, created_at) VALUES (?, ?, ?, ?, ?, ?)",
		post.ID, post.Title, post.Author, post.Content, post.Likes, post.CreatedAt)
	return err
}

func (r *MySQLRepository) Delete(id string) error {
	if r.client == nil {
		return fmt.Errorf("mysql repository is not initilized")
	}
	_, err := r.client.Exec("DELETE FROM posts WHERE id=?", id)
	return err
}
