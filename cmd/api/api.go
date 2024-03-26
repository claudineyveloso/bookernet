package api

import (
	"database/sql"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/services/healthy"
	"github.com/claudineyveloso/bookernet.git/internal/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	r := mux.NewRouter()
	healthy.RegisterRoutes(r)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(r)

	bucketStore := user.NewStore(s.db)
	bucketHandler := user.NewHandler(bucketStore)
	bucketHandler.RegisterRoutes(r)

	return http.ListenAndServe("localhost:8080", r)
}
