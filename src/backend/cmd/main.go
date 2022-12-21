package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/talkanbaev-artur/poisson-equation-solver/src/backend/config"
	"github.com/talkanbaev-artur/poisson-equation-solver/src/backend/logic"
	"github.com/talkanbaev-artur/poisson-equation-solver/src/backend/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	lg := logrus.New()
	config.Init()

	s := logic.NewAPIService()

	r := mux.NewRouter()
	server.MakeMuxRoutes(s, r, lg)
	r.PathPrefix("/debug/pprof/").Handler(http.DefaultServeMux)
	go func() {
		c := cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowedMethods: []string{"POST", "GET", "OPTIONS"}})
		srv := http.Server{
			Handler:      c.Handler(r),
			ReadTimeout:  time.Minute,
			WriteTimeout: time.Minute,
			Addr:         fmt.Sprintf("0.0.0.0:%d", config.Config().Port),
		}
		lg.WithField("port", config.Config().Port).Info("The web server is up")
		lg.Error(srv.ListenAndServe().Error())
		cancel()
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		lg.Info("Catched shutdown signal - initiating graceful shutdown")
		cancel()
	}()

	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				runtime.GC()
			case <-c:
				t.Stop()
				return
			}
		}
	}()

	lg.Info("Initialised numericals application")
	<-ctx.Done()
	lg.Info("Shutting down...")
}
