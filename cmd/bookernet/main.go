package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/configs"
	"github.com/claudineyveloso/bookernet.git/internal/handler"
	"github.com/gorilla/mux"
)

func main() {

	err := configs.Load()
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Bem-vindo ao BookerNet!"))
		if err != nil {
			http.Error(w, "Erro ao escrever resposta", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	// Endpoint User
	r.HandleFunc("/create_user", handler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/get_users", handler.GetUsersHandler).Methods("GET")
	r.HandleFunc("/get_user", handler.GetUserHandler).Methods("GET")
	r.HandleFunc("/disable_user", handler.DisableUserHandler).Methods("PUT")
	r.HandleFunc("/update_user", handler.UpdateUsersHandler).Methods("PUT")
	r.HandleFunc("/update_password_user", handler.UpdatePasswordUsersHandler).Methods("PUT")

	port := configs.GetServerPort()
	fmt.Printf("Servidor escutando na porta %s\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

}
