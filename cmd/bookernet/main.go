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

	// Endpoint Login
	r.HandleFunc("/login", handler.LoginHandler).Methods("POST")

	// Endpoint User
	r.HandleFunc("/create_user", handler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/get_users", handler.GetUsersHandler).Methods("GET")
	r.HandleFunc("/get_user", handler.GetUserHandler).Methods("GET")
	r.HandleFunc("/disable_user", handler.DisableUserHandler).Methods("PUT")
	r.HandleFunc("/update_user", handler.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/update_password_user", handler.UpdatePasswordUserHandler).Methods("PUT")

	// Endpoint Bucket
	r.HandleFunc("/create_bucket", handler.CreateBucketHandler).Methods("POST")
	r.HandleFunc("/get_buckets", handler.GetBucketsHandler).Methods("GET")
	r.HandleFunc("/get_bucket", handler.GetBucketHandler).Methods("GET")
	r.HandleFunc("/update_bucket", handler.UpdateBucketHandler).Methods("PUT")

	// Endpoint Owner
	r.HandleFunc("/create_owner", handler.CreateOwnerHandler).Methods("POST")

	port := configs.GetServerPort()
	fmt.Printf("Servidor escutando na porta %s\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

}
