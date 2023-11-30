package model

import "time"

type Post struct {
	ID        string
	CreatedAt time.Time
	Title     string
	Content   string
	Author    string
	Likes     int64
	Comments  []Comment
}

type Comment struct {
	Author  string
	Content string
}
