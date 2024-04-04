package api

import (
	"database/sql"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/services/bucket"
	"github.com/claudineyveloso/bookernet.git/internal/services/customer"
	"github.com/claudineyveloso/bookernet.git/internal/services/healthy"
	"github.com/claudineyveloso/bookernet.git/internal/services/owner"
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

	bucketStore := bucket.NewStore(s.db)
	bucketHandler := bucket.NewHandler(bucketStore)
	bucketHandler.RegisterRoutes(r)

	ownerStore := owner.NewStore(s.db)
	//personStore := person.NewStore(s.db)
	//addressStore := address.NewStore(s.db)
	ownerHandler := owner.NewHandler(ownerStore)
	ownerHandler.RegisterRoutes(r)

	customerStore := customer.NewStore(s.db)
	//personStore = person.NewStore(s.db)
	//addressStore = address.NewStore(s.db)
	customerHandler := customer.NewHandler(customerStore)
	customerHandler.RegisterRoutes(r)

	return http.ListenAndServe("localhost:8080", r)
}
