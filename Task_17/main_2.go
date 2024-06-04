package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	l := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	logHandler := logMiddleware(l)
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: authHandler(
			logHandler(mux),
			l,
		),
	}
	if err := httpServer.ListenAndServe(); err != nil {
		l.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	log.Println("resp:", msg)
	fmt.Fprint(res, msg)
}

func authHandler(h http.Handler, l *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			msg := "пользователь не авторизован"
			l.Println(msg)
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func logMiddleware(l *log.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
