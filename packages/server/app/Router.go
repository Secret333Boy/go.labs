package app

import "net/http"

type route struct {
	pattern string
	method  string
	handler http.Handler
}

type Router struct {
	routes []*route
}

func (router *Router) Handle(pattern string, handler http.Handler) {
	router.routes = append(router.routes, &route{pattern, "*", handler})
}

func (router *Router) handleFunc(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	router.routes = append(router.routes, &route{pattern, method, http.HandlerFunc(handler)})
}

func (router *Router) Get(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	router.handleFunc(http.MethodGet, pattern, handler)
}

func (router *Router) Post(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	router.handleFunc(http.MethodPost, pattern, handler)
}

func (router *Router) Put(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	router.handleFunc(http.MethodPut, pattern, handler)
}

func (router *Router) Delete(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	router.handleFunc(http.MethodDelete, pattern, handler)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.routes[0].handler.ServeHTTP(w, r)
}
