package routes

import (
	"net/http"

	"grhamm.com/todo/handler"
)

func RegisterRoute() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.Health)
	mux.HandleFunc("/todo/get", handler.GetTodo)
	mux.HandleFunc("/todo/post", handler.InsertTodo)
	mux.HandleFunc("/todo/set-finished", handler.SetTodoFinished)

	return mux
}
