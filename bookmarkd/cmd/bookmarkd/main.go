package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dolldot/scrapyard/bookmarkd/config"
	"github.com/dolldot/scrapyard/bookmarkd/internal/clients"
	"github.com/dolldot/scrapyard/bookmarkd/internal/logger"
	"github.com/dolldot/scrapyard/bookmarkd/internal/routes"
)

func main() {
	logger.InitializeLogger()
	config.InitializeConfig()

	os.Setenv("TZ", config.Load.Server.TimeZone)

	db := clients.NewEntclient()

	server := routes.SetupRouter(db)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer db.Close()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown: %s", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("server exiting")
	}
}
