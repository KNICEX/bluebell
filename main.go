package main

import (
	"context"
	"errors"
	"fmt"
	"myBulebell/bootstrap"
	"myBulebell/pkg/conf"
	"myBulebell/pkg/logger"
	"myBulebell/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	bootstrap.Init()

	r := routes.Init()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", conf.ServerConf.Port),
		Handler: r,
	}

	go func() {
		fmt.Printf("Server started at :%s%s\n", conf.ServerConf.Port, conf.ServerConf.Prefix)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.L().Error("server shutdown error", err)
	}

	if err := bootstrap.Shutdown(); err != nil {
		logger.L().Error("bootstrap shutdown error", err)
	}

	fmt.Println("Server exiting")

}
