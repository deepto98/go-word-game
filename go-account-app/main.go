package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/deepto98/go-word-game/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := handler.Config{
		Router: router,
	}

	handler.NewHandler(&config)

	// router.GET("/api/account", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"a": "b",
	// 	})
	// })

	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	//Implementing graceful shutdown - https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it

	//Ctrl+C - SIGINT, Ctrl + Z - SIGSTOP
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//Blocks until a signal is received
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
