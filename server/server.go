package server

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/jesperkha/w4d/config"
	"github.com/jesperkha/w4d/database"
)

type Server struct {
	server *http.Server
	config config.Config
	db     *database.Database
}

func New(config config.Config) *Server {
	db := database.New(config)
	r := chi.NewRouter()

	r.Get("/", indexHandler(db))

	server := &http.Server{
		Handler: r,
		Addr:    config.Port,
	}

	return &Server{
		server: server,
		config: config,
		db:     db,
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
