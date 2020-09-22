package http

import (
	cache2 "github.com/debuginn/go-implement-your-cache-server/chapter2/server/cache"
	"net/http"
)

type Server struct {
	cache2.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	http.ListenAndServe(":12345", nil)
}

func New(c cache2.Cache) *Server {
	return &Server{c}
}
