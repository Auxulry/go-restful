package config

import (
	"net/http"
	"time"
)

const (
	Addr           = "localhost:5000"
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
	MaxHeaderBytes = 1 << 20
)

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:           Addr,
		Handler:        handler,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: MaxHeaderBytes,
	}
}
