package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(
	host string,
	port string,
	readTimeout time.Duration,
	writeTimeout time.Duration,
	router *gin.Engine,
) error {
	address := host + ":" + port
	s.httpServer = &http.Server{
		Addr:           address,
		Handler:        router,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
