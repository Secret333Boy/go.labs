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
			limit = 10
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			offset = 0
		}
		err = json.NewEncoder(w).Encode(postsService.GetAllPosts(limit, offset))
		if err != nil {
			http.Error(w, "Failed encoding json", http.StatusInternalServerError)
		}
		//TODO: count pages and set default limit and offset
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

	//update post by id
	router.PATCH("/api/posts/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		account, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
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
		err = postsService.UpdatePost(account, id, createPostDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	//delete post by id
	router.DELETE("/api/posts/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		account, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = postsService.RemovePost(account, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	//get all messages
	router.GET("/api/posts/:id/messages", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		_, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
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
		err = json.NewEncoder(w).Encode(postsService.GetAllMessagesByPostId(id, limit, offset))
		if err != nil {
			http.Error(w, "Failed encoding json", http.StatusInternalServerError)
		}
		//TODO: count pages and set default limit and offset
	})

	//create new message
	router.POST("/api/posts/:id/messages", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		account, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		addMessageDto := &dtos.AddMessageDto{}
		err = json.NewDecoder(r.Body).Decode(addMessageDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validationErr := addMessageDto.Validate()
		if validationErr != nil {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}
		postsService.AddMessageByPostId(account, id, addMessageDto)

	})

	//get message by id
	router.GET("/api/posts/:id/messages/:messageId", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		_, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		messageId, err := strconv.Atoi(ps.ByName("messageId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(postsService.GetOneMessageByPostId(id, messageId))
		if err != nil {
			http.Error(w, "Failed encoding json", http.StatusInternalServerError)
		}
	})

	//update message by id
	router.PATCH("/api/posts/:id/messages/:messageId", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		account, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		messageId, err := strconv.Atoi(ps.ByName("messageId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		addMessageDto := &dtos.AddMessageDto{}
		err = json.NewDecoder(r.Body).Decode(addMessageDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validationErr := addMessageDto.Validate()
		if validationErr != nil {
			http.Error(w, validationErr.Error(), http.StatusBadRequest)
			return
		}
		err = postsService.UpdateMessageByPostId(account, id, messageId, addMessageDto)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	//delete message by id
	router.DELETE("/api/posts/:id/messages/:messageId", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		account, err := middlewares.UseAuth(w, r)
		if err != nil {
			return
		}

		id, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		messageId, err := strconv.Atoi(ps.ByName("messageId"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = postsService.RemoveMessageByPostId(account, id, messageId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})

	return router

}
