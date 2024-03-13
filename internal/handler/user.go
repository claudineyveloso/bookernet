package handler

import (
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/controller"
	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/infra/database"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.OpenConnection()
	if err != nil {
		// If an error occurs when opening the connection, send a 500 Internal Server Error to the client
		http.Error(w, "Erro ao abrir a conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	dbConn := db.New(conn)
	controller.CreateUserController(w, r, dbConn)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.OpenConnection()
	if err != nil {
		// If an error occurs when opening the connection, send a 500 Internal Server Error to the client
		http.Error(w, "Erro ao abrir a conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	dbConn := db.New(conn)
	controller.GetUserController(w, r, dbConn)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.OpenConnection()
	if err != nil {
		// If an error occurs when opening the connection, send a 500 Internal Server Error to the client
		http.Error(w, "Erro ao abrir a conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	dbConn := db.New(conn)
	controller.GetUsersController(w, r, dbConn)
}

func UpdateUsersHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.OpenConnection()
	if err != nil {
		// If an error occurs when opening the connection, send a 500 Internal Server Error to the client
		http.Error(w, "Erro ao abrir a conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	dbConn := db.New(conn)
	controller.UpdateUserController(w, r, dbConn)
}

func UpdatePasswordUsersHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.OpenConnection()
	if err != nil {
		// If an error occurs when opening the connection, send a 500 Internal Server Error to the client
		http.Error(w, "Erro ao abrir a conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	dbConn := db.New(conn)
	controller.UpdatePasswordController(w, r, dbConn)
}

func DisableUserHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.OpenConnection()
	if err != nil {
		// If an error occurs when opening the connection, send a 500 Internal Server Error to the client
		http.Error(w, "Erro ao abrir a conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	dbConn := db.New(conn)
	controller.DisableUserController(w, r, dbConn)
}
