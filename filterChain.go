package crucible

import (
	"net/http"
)

type Filter interface {
	DoFilter(writer http.ResponseWriter, request *http.Request, filterChain chan Filter)
}

func Handle() *filterHandler {
	return handler
}

func AddFilter(f Filter) {
	Handle().addFilter(f)
}

func AddHandler(h http.Handler) {
	Handle().addHandler(h)
}

type filterHandler struct {
	filterQueue  chan Filter
	handlerQueue chan http.Handler
}

var handler *filterHandler = &filterHandler{
	filterQueue:  make(chan Filter, 8),
	handlerQueue: make(chan http.Handler, 4)}

func (handler *filterHandler) addFilter(f Filter) {
	handler.filterQueue <- f
}

func (handler *filterHandler) addHandler(h http.Handler) {
	handler.handlerQueue <- h
}

func (handler *filterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for filter := range handler.filterQueue {
		filter.DoFilter(writer, request, handler.filterQueue)
	}

	for handler := range handler.handlerQueue {
		handler.ServeHTTP(writer, request)
	}

	return
}
