package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {

	var stopSrv chan struct{}
	stopSrv = make(chan struct{})

	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	m.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		stopSrv <- struct{}{}
	})
	srv := http.Server{
		Addr:    ":7000",
		Handler: m,
	}

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return srv.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("abcde")
			return ctx.Err()
		case <-stopSrv:
			fmt.Println("abcdef")
			srv.Shutdown(ctx)
			return errors.Errorf("server stop.")
		}
	})

	g.Go(func() error {
		signChan := make(chan os.Signal, 0)
		signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-ctx.Done():
			fmt.Println("abc")
			return ctx.Err()
		case s := <-signChan:
			fmt.Println("abcd")
			return errors.Errorf("receive signal: %v.", s)
		}
	})

	fmt.Printf("errgroup exit: %+v.", g.Wait())
}
