package posts

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go.labs/server/app/middlewares"
	"go.labs/server/app/services/posts"

	"github.com/julienschmidt/httprouter"
)

type PostsHandler struct {
	service           posts.PostsAndMessages
	useAuthMiddleware *middlewares.UseAuthMiddleware
}

func NewPostsHandler(service posts.PostsAndMessages, useAuthMiddleware *middlewares.UseAuthMiddleware) *PostsHandler {
	return &PostsHandler{service: service, useAuthMiddleware: useAuthMiddleware}
}

func (h *PostsHandler) RegisterHandler(router *httprouter.Router) {
	//get all posts
	router.GET("/api/posts", h.getAllPosts)
	//create new post
	router.POST("/api/posts", h.createPost)
	//get post by id
	router.GET("/api/posts/:id", h.getPost)
	//update post by id
	router.PATCH("/api/posts/:id", h.updatePost)
	//delete post by id
	router.DELETE("/api/posts/:id", h.deletePost)
	//get all messages
	router.GET("/api/posts/:id/messages", h.getAllMessages)
	//create new message
	router.POST("/api/posts/:id/messages", h.createMessage)
	//get message by id
	router.GET("/api/posts/:id/messages/:messageId", h.getMessage)
	//update message by id
	router.PATCH("/api/posts/:id/messages/:messageId", h.updateMessage)
	//delete message by id
	router.DELETE("/api/posts/:id/messages/:messageId", h.deleteMessage)
}

func (h *PostsHandler) getAllPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := h.useAuthMiddleware.UseAuth(w, r)
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
	err = json.NewEncoder(w).Encode(h.service.GetAllPosts(limit, offset))
	if err != nil {
		http.Error(w, "Failed encoding json", http.StatusInternalServerError)
	}
	//TODO: count pages
}

func (h *PostsHandler) createPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	account, err := h.useAuthMiddleware.UseAuth(w, r)
	if err != nil {
		return
	}

	createPostDto := &CreatePostDto{}
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
	err = h.service.AddPost(account, createPostDto.Title, createPostDto.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *PostsHandler) getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := h.useAuthMiddleware.UseAuth(w, r)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(h.service.GetOnePost(id))
	if err != nil {
		http.Error(w, "Failed encoding json", http.StatusInternalServerError)
	}
}

func (h *PostsHandler) updatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	account, err := h.useAuthMiddleware.UseAuth(w, r)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createPostDto := &CreatePostDto{}
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
	err = h.service.UpdatePost(account, id, createPostDto.Title, createPostDto.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *PostsHandler) deletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	account, err := h.useAuthMiddleware.UseAuth(w, r)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.RemovePost(account, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *PostsHandler) getAllMessages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := h.useAuthMiddleware.UseAuth(w, r)
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
		limit = 10
	}
	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		offset = 0
	}

	resultMessages, err := h.service.GetAllMessagesByPostId(id, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = json.NewEncoder(w).Encode(resultMessages)
	if err != nil {
		http.Error(w, "Failed encoding json", http.StatusInternalServerError)
	}
	//TODO: count pages
}

func (h *PostsHandler) createMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	account, err := h.useAuthMiddleware.UseAuth(w, r)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addMessageDto := &AddMessageDto{}
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

	err = h.service.AddMessageByPostId(account, id, addMessageDto.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h *PostsHandler) getMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := h.useAuthMiddleware.UseAuth(w, r)
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
	resultMessages, err := h.service.GetOneMessageByPostId(id, messageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = json.NewEncoder(w).Encode(resultMessages)
	if err != nil {
		http.Error(w, "Failed encoding json", http.StatusInternalServerError)
	}
}

func (h *PostsHandler) updateMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	account, err := h.useAuthMiddleware.UseAuth(w, r)
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

	addMessageDto := &AddMessageDto{}
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
	err = h.service.UpdateMessageByPostId(account, id, messageId, addMessageDto.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *PostsHandler) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	account, err := h.useAuthMiddleware.UseAuth(w, r)
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

	err = h.service.RemoveMessageByPostId(account, id, messageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
