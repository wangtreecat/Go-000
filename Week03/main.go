package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type handler struct{}

type Server struct {
	addr string
	http *http.Server
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("WEEK03"))
}

func NewServer(addr string) *Server {
	mux := http.NewServeMux()
	mux.Handle("/", new(handler))
	http := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	return &Server{
		addr: addr,
		http: http,
	}
}

func (s *Server) ServerStart() error {
	return s.http.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	log.Printf("server shutdown, addr : %v", s.addr)
	if err := s.http.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func StopAll(s1 *Server, s2 *Server) error {
	g, ctx := errgroup.WithContext(context.Background())

	log.Println("ready to stopall")

	g.Go(func() error {
		return s1.ShutDown(ctx)
	})

	g.Go(func() error {
		return s2.ShutDown(ctx)
	})

	return g.Wait()
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	s1 := NewServer(":8080")
	s2 := NewServer(":8081")

	g.Go(func() error {
		return s1.ServerStart()
	})

	g.Go(func() error {
		return s2.ServerStart()
	})

	g.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-sigs: // 关闭信号
			log.Println("signal exit")
		case <-ctx.Done(): // 服务关闭
			log.Println("server exit")
		}
		return StopAll(s1, s2)
		//exit 出口
	})

	if err := g.Wait(); err != nil {
		log.Println("ShutDown")
		return
	}
}
