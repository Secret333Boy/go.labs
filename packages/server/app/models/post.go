package models

import (
	"time"
)

type Post struct {
	Id          int
	Account     *Account
	Title       string
	Description string
	Tags        []string
	PublishedAt time.Time
	model       *PostModel
}

type PostModel struct {
	posts  []Post
	lastId int
}

func NewPostModel() *PostModel {
	var model = new(PostModel)
	model.posts = make([]Post, 0)
	return model
}

func (model *PostModel) FindAll() []Post {
	return model.posts
}

func (model *PostModel) FindOne(id int) *Post {
	for _, post := range model.posts {
		if post.Id == id {
			return &post
		}
	}

	return nil
}

func (model *PostModel) Add(post *Post) {
	post.model = model
	post.Id = model.lastId + 1
	model.posts = append(model.posts, *post)
	model.lastId++
}

func (model *PostModel) Delete(id int) {
	for i, post := range model.posts {
		if post.Id == id {
			post.model = nil
			model.posts = append(model.posts[:i], model.posts[i+1:]...)
			return
		}
	}
}

func (model *PostModel) Exists(id int) bool {
	for _, post := range model.posts {
		if post.Id == id {
			return true
		}
	}
	return false
}
