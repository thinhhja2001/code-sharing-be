package httpserver

import (
	"code-sharing-be/cmd/config"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	_defaultAddr            = ":80"
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	Engine *gin.Engine
	notify chan error

	address         string
	prefork         bool
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration
}

func New(cfg *config.Config) *Server {
	s := &Server{
		Engine:          nil,
		notify:          make(chan error, 1),
		address:         cfg.HTTP.Port,
		readTimeout:     _defaultReadTimeout,
		writeTimeout:    _defaultWriteTimeout,
		shutdownTimeout: _defaultShutdownTimeout,
	}
	engine := gin.Default()
	s.Engine = engine

	return s
}

// Start -.
func (s *Server) Start() {
	go func() {
		s.notify <- s.Engine.Run(s.address)
		println("s.notify is", s.notify)
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

//// Shutdown -.
//func (s *Server) Shutdown() error {
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//}
