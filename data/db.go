package data

import (
	"database/sql"
	"log"
	"os"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"grhamm.com/todo/entity"
)

func InitDatabase() {

	var database *sql.DB

	if _, err := os.Stat("database.db"); os.IsNotExist(err) {

		file, err := os.Create("database.db")
		if err != nil {
			log.Fatal(err.Error())
		}

		file.Close()

		log.Print("DB created")

		database, _ = sql.Open("sqlite3", "./database.db")

		createTodoTable(database)

	} else {
		log.Print("DB initialized")

		database, _ = sql.Open("sqlite3", "./database.db")

	}

	defer database.Close()

}

func createTodoTable(db *sql.DB) {
	query := `CREATE TABLE todo(
		id varchar(254) NOT NULL PRIMARY KEY,
		description TEXT,
		finished INT
		);`

	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	statement.Exec()
	log.Print("TODO table created")
}

func InsertTodo(todo entity.Todo) entity.Todo {
	db, _ := sql.Open("sqlite3", "./database.db")

	defer db.Close()

	query := `INSERT INTO todo (id, description, finished) values (?, ?, ?)`
	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	uuid := uuid.New()

	todo.Id = uuid

	_, err = statement.Exec(uuid, todo.Description, todo.Finished)

	if err != nil {
		log.Fatal(err.Error())
	}

	return todo
}

func FindTodo() []entity.Todo {
	db, _ := sql.Open("sqlite3", "./database.db")

	defer db.Close()

	var todoList []entity.Todo

	row, err := db.Query("SELECT * FROM todo")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer row.Close()

	for row.Next() {
		var todo entity.Todo

		row.Scan(
			&todo.Id,
			&todo.Description,
			&todo.Finished,
		)

		todoList = append(todoList, todo)
	}

	return todoList
}

func SetFinishTodo(id uuid.UUID) {
	db, _ := sql.Open("sqlite3", "./database.db")

	defer db.Close()

	query := `UPDATE todo SET finished = (?) WHERE id = (?)`
	statement, err := db.Prepare(query)

	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = statement.Exec(1, id)

	if err != nil {
		log.Fatal(err.Error())
	}

}
