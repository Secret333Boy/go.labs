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

func (model *PostModel) FindAll(limit int, offset int) []Post {

	if len(model.posts) > 0 {
		if offset > len(model.posts) {
			offset = len(model.posts)
		}

		end := offset + limit
		if end > len(model.posts) {
			end = len(model.posts)
		}
		return model.posts[offset:end]
	} else {
		return model.posts
	}

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
