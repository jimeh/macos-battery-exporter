//go:build darwin

package prombat

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const DefaultNamespace = "macos"

type ServerOptions struct {
	Bind         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type Server struct {
	*http.Server
	registry *prometheus.Registry
	mux      *http.ServeMux
}

func NewServer(registry *prometheus.Registry, options ServerOptions) *Server {
	if options.Bind == "" {
		options.Bind = "127.0.0.1"
	}
	if options.Port == 0 {
		options.Port = 9108
	}
	if options.ReadTimeout == 0 {
		options.ReadTimeout = 5 * time.Second
	}
	if options.WriteTimeout == 0 {
		options.WriteTimeout = 10 * time.Second
	}
	if options.IdleTimeout == 0 {
		options.IdleTimeout = 30 * time.Second
	}

	mux := http.NewServeMux()
	mux.Handle(
		"/metrics",
		promhttp.HandlerFor(registry, promhttp.HandlerOpts{}),
	)

	return &Server{
		mux:      mux,
		registry: registry,
		Server: &http.Server{
			Addr:         fmt.Sprintf("%s:%d", options.Bind, options.Port),
			ReadTimeout:  options.ReadTimeout,
			WriteTimeout: options.WriteTimeout,
			IdleTimeout:  options.IdleTimeout,
			Handler:      mux,
		},
	}
}

func (s *Server) Register(
	collectors ...prometheus.Collector,
) error {
	for _, c := range collectors {
		err := s.registry.Register(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) ListenAndServe() error {
	slog.Info(
		"starting prometheus server",
		slog.String("addr", s.Addr),
	)

	return s.Server.ListenAndServe()
}

func RunServer(
	namespace string,
	registry *prometheus.Registry,
	options ServerOptions,
) error {
	if namespace == "" {
		namespace = DefaultNamespace
	}

	s := NewServer(registry, options)

	collector := NewCollector(namespace)
	err := s.Register(collector)
	if err != nil {
		return err
	}

	return s.ListenAndServe()
}
