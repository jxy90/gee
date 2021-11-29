package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: map[string]HandlerFunc{}}
}

func (g *Engine) AddRoute(method, param string, handler HandlerFunc) {
	key := method + "-" + param
	g.router[key] = handler
}

func (g *Engine) Get(param string, handler HandlerFunc) {
	g.AddRoute("GET", param, handler)
}
func (g *Engine) Post(param string, handler HandlerFunc) {
	g.AddRoute("POST", param, handler)
}

func (g *Engine) Run(port string) {
	http.ListenAndServe(":"+port, g)
}

func (g *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := g.router[key]; ok {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 Not Found %v \n", key)
	}
}
