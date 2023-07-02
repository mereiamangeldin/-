package transport

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mereiamangeldin/OnePlusTask2/transport/handler"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	handler    *handler.Manager
	router     *echo.Echo
}

func NewServer(handler *handler.Manager) *Server {
	return &Server{handler: handler}
}

func (s *Server) Run(ctx context.Context) error {
	s.router = s.BuildEngine()
	s.InitRoutes()
	go func() {
		if err := s.router.Start(fmt.Sprintf(":%d", 8080)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.router.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	return e
}
