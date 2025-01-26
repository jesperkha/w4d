package server

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/jesperkha/w4d/config"
)

type Server struct {
	server *http.Server
	config config.Config
}

func New(config config.Config) *Server {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app/index.html")
	})

	server := &http.Server{
		Handler: r,
		Addr:    config.Port,
	}

	return &Server{
		server: server,
		config: config,
	}
}

func (s *Server) ListenAndServe(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		<-ctx.Done()
		s.server.Shutdown(ctx)
	}()

	log.Printf("Listening on port %s\n", s.config.Port)

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

	wg.Done()
}
