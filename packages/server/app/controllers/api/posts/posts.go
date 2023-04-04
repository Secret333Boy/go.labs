package posts

import (
	"encoding/json"
	"go.labs/server/app/controllers/api/posts/dtos"
	"go.labs/server/app/middlewares"
	"go.labs/server/app/router"
	"go.labs/server/app/services"
	"net/http"
)

func GetPostsRouter() *router.Router {
	router := router.NewRouter()
	postsService := services.PostsService

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		err = json.NewEncoder(w).Encode(postsService.GetAllPosts())
		if err != nil {
			http.Error(w, "Failed encoding json", http.StatusInternalServerError)
		}

	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
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

	return router

}
