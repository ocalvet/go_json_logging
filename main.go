package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"ping": "pong"}`))
	})

	// RESTy routes for "todo" resource
	r.Route("/todo", func(r chi.Router) {
		r.Post("/", createTodo) // POST /todo
		r.Get("/", getTodos)    // GET /todo/search
		// // Subrouters:
		// r.Route("/{todoID}", func(r chi.Router) {
		// 	r.Get("/", getTodo)       // GET /todo/123
		// 	r.Put("/", updateTodo)    // PUT /todo/123
		// 	r.Delete("/", deleteTodo) // DELETE /todo/123
		// })
	})

	http.ListenAndServe(":2018", r)
}

var todos = []string{}

func createTodo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{ "message": "Error reading body: %s" }`, err.Error())
	}
	todo := string(body)
	todos = append(todos, todo)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{ "success": true}`)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	ts, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{ "message": "reading todos: %s" }`, err.Error())
	}
	w.WriteHeader(http.StatusOK)
	w.Write(ts)
}
