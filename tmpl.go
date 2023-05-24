// Package tmpl builds a middleware that FIXME.
package traefik_plugin_tmpl

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct{}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// Tmpl a plugin to FIXME.
type Tmpl struct {
	next   http.Handler
	name   string
	config *Config
}

// New create a new Tmpl plugin.
func New(_ context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &Tmpl{
		next:   next,
		name:   name,
		config: config,
	}, nil
}

func (a *Tmpl) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// noop
	a.next.ServeHTTP(rw, req)
}
