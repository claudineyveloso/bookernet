package handler

import (
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/controller"
	"github.com/claudineyveloso/bookernet.git/internal/db"
	"github.com/claudineyveloso/bookernet.git/internal/infra/database"
)

func CreateBucketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.OpenConnection()
	if err != nil {
		// If an error occurs when opening the connection, send a 500 Internal Server Error to the client
		http.Error(w, "Erro ao abrir a conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	dbConn := db.New(conn)
	controller.CreateBucketController(w, r, dbConn)
}

func GetBucketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.OpenConnection()
	if err != nil {
		// If an error occurs when opening the connection, send a 500 Internal Server Error to the client
		http.Error(w, "Erro ao abrir a conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	dbConn := db.New(conn)
	controller.GetBucketController(w, r, dbConn)
}

func GetBucketsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.OpenConnection()
	if err != nil {
		// If an error occurs when opening the connection, send a 500 Internal Server Error to the client
		http.Error(w, "Erro ao abrir a conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	dbConn := db.New(conn)
	controller.GetBucketsController(w, r, dbConn)
}

func UpdateBucketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := database.OpenConnection()
	if err != nil {
		// If an error occurs when opening the connection, send a 500 Internal Server Error to the client
		http.Error(w, "Erro ao abrir a conexão com o banco de dados", http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	dbConn := db.New(conn)
	controller.UpdateBucketController(w, r, dbConn)
}
