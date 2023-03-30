package router

import (
	"fmt"
	"net/http"
	"regexp"
)

type route struct {
	pattern string
	method  string
	handler http.Handler
}

type Router struct {
	routes      []*route
	root        string
	usedRouters []*Router
}

func NewRouter() *Router {
	router := &Router{}
	return router
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

func (router *Router) All(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	router.handleFunc("*", pattern, handler)
}

func (router *Router) Use(pattern string, router2 *Router) {
	unifiedPattern := pattern

	if len(unifiedPattern) > 0 && unifiedPattern[len(unifiedPattern)-1] != '*' {
		unifiedPattern = unifiedPattern + "*"
	}

	if len(unifiedPattern) > 1 && unifiedPattern[len(unifiedPattern)-2] != '/' {
		unifiedPattern = unifiedPattern[:len(unifiedPattern)-1] + "/*"
	}

	router.Handle(unifiedPattern, router2)

	router2.root = pattern

	if pattern[len(pattern)-1] == '/' {
		router2.root = pattern[:len(pattern)-1]
	}

	for _, usedRouter := range router2.usedRouters {
		usedRouter.root = router2.root + usedRouter.root
	}

	router.usedRouters = append(router.usedRouters, router2)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	for _, route := range router.routes {
		routerRegExpEnd := "/?$"
		if route.pattern[len(route.pattern)-1] == '/' {
			routerRegExpEnd = "?$"
		}

		if route.pattern[len(route.pattern)-1] == '*' {
			routerRegExpEnd = ""
		}

		patternRegexp, err := regexp.Compile("^" + router.root + route.pattern + routerRegExpEnd)

		if err != nil {
			continue
		}

		/*Only for debugging!*/
		// fmt.Println(w, "\nMethod: "+r.Method+"\nURL path: "+r.URL.Path+"\nRegexp: "+patternRegexp.String()+"\nRouter root: "+router.root)
		if (r.Method == route.method || route.method == "*") && patternRegexp.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Not found "+r.URL.Path)
}
