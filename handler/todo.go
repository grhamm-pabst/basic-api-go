package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"grhamm.com/todo/entity"
	"grhamm.com/todo/repo"
)

func Health(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "SERVER RUNNING NORMALLY")
}

func GetTodo(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {

		todoList := repo.Index()

		jsonBody, err := json.Marshal(todoList)
		if err != nil {
			http.Error(responseWriter, "Error parsing todo to json", http.StatusInternalServerError)
		}

		responseWriter.Write(jsonBody)
	} else {
		http.Error(responseWriter, "", http.StatusMethodNotAllowed)
	}

}

func InsertTodo(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(responseWriter, "Error reading request body", http.StatusInternalServerError)

		}

		var todo entity.Todo
		err = json.Unmarshal(body, &todo)
		if err != nil {
			http.Error(responseWriter, "Error parsing request body", http.StatusInternalServerError)

		}

		todo = repo.Insert(todo)

		jsonBody, err := json.Marshal(todo)
		if err != nil {
			http.Error(responseWriter, "Error parsing response body", http.StatusInternalServerError)

		}

		responseWriter.Write(jsonBody)
	} else {
		http.Error(responseWriter, "", http.StatusMethodNotAllowed)
	}

}

func SetTodoFinished(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method == "PUT" {
		query := request.URL.Query()
		todoId := query.Get("todo-id")

		todoUUID, err := uuid.Parse(todoId)
		if err != nil {
			http.Error(responseWriter, "Error parsing uuid", http.StatusBadRequest)

		}

		updatedTodo := repo.UpdateToFinished(todoUUID)

		jsonBody, err := json.Marshal(updatedTodo)
		if err != nil {
			http.Error(responseWriter, "Error parsing response", http.StatusInternalServerError)

		}

		responseWriter.Write(jsonBody)

	} else {
		http.Error(responseWriter, "", http.StatusMethodNotAllowed)
	}
}
