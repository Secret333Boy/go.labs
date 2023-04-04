package posts

import (
	"encoding/json"
	"go.labs/server/app/controllers/api/posts/dtos"
	"go.labs/server/app/middlewares"
	"go.labs/server/app/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func GetPostsRouter() *httprouter.Router {
	router := httprouter.New()
	postsService := services.PostsService
	//get all posts
	router.GET("/api/posts", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		_, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(postsService.GetAllPosts(limit, offset))
		if err != nil {
			http.Error(w, "Failed encoding json", http.StatusInternalServerError)
		}

	})

	//create new post
	router.POST("/api/posts", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		account, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		createPostDto := &dtos.CreatePostDto{}
		err = json.NewDecoder(r.Body).Decode(createPostDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validationErr := createPostDto.Validate()
		if validationErr != nil {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}
		postsService.AddPost(account, createPostDto)

	})

	//get post by id
	router.GET("/api/posts/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		_, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(postsService.GetOnePost(id))
		if err != nil {
			http.Error(w, "Failed encoding json", http.StatusInternalServerError)
		}
	})

	return router

}
