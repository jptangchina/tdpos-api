package main

import (
	"context"
	log "github.com/jptangchina/log4g"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"tdpos/common"
	_ "tdpos/initialize"
	"time"
)

func main() {
	router := common.DefaultEngine()
	srv := &http.Server{
		Addr:    ":" + viper.GetString("port"),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Info("Server exiting")
}
