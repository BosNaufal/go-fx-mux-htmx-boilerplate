package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/BosNaufal/go-fx-mux-htmx-boilerplate/internal/routes"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
)

// NewHTTPServer builds an HTTP server that will begin serving requests
// when the Fx application starts.
func NewHTTPServer(lc fx.Lifecycle, router *mux.Router) *http.Server {
	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("Starting HTTP server at", srv.Addr)
			// ref: https://github.com/uber-go/fx/issues/627#issuecomment-399235227
			go func() {
				err := srv.ListenAndServe()
				if err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("shutting down...")
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

func main() {
	app := fx.New(
		fx.Provide(routes.NewRouter),
		fx.Provide(NewHTTPServer),
		fx.Invoke(func(*http.Server) {}),
	)
	app.Run()
}
