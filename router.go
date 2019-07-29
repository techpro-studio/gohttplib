package gohttplib

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router interface {
	Get(path string, handler http.Handler)
	Post(path string, handler http.Handler)
	Put(path string, handler http.Handler)
	Delete(path string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type DefaultRouter struct {
	DefaultRouter *httprouter.Router
}

func (self *DefaultRouter) Get(path string, handler http.Handler) {
	self.DefaultRouter.GET(path, wrapHandler(handler))
}

func (self *DefaultRouter) Post(path string, handler http.Handler) {
	self.DefaultRouter.POST(path, wrapHandler(handler))
}

func (self *DefaultRouter) Put(path string, handler http.Handler) {
	self.DefaultRouter.PUT(path, wrapHandler(handler))
}

func (self *DefaultRouter) Delete(path string, handler http.Handler) {
	self.DefaultRouter.DELETE(path, wrapHandler(handler))
}

func (self *DefaultRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	self.DefaultRouter.ServeHTTP(w, req)
}

func NewDefaultRouter() *DefaultRouter {
	return &DefaultRouter{httprouter.New()}
}
