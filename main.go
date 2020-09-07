package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yyh-gl/go-flat-package/app"
	"github.com/yyh-gl/go-flat-package/app/repository"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	taskRepo := repository.NewTaskRepository()
	taskHandler := app.NewTaskHandler(taskRepo)
	
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/tasks/{id}", taskHandler.Get).Methods(http.MethodGet)

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		fmt.Println("========================")
		fmt.Println("Server Start >> http://localhost" + s.Addr)
		fmt.Println("========================")
		errCh <- s.ListenAndServe()
	}()

	select {
	case err := <-errCh:
		fmt.Println("Error happened:", err.Error())
	case sig := <-sigCh:
		fmt.Println("Signal received:", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		fmt.Println("Graceful shutdown failed:", err.Error())
	}
	fmt.Println("Server shutdown")
}
