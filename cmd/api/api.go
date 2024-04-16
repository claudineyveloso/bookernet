package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/claudineyveloso/bookernet.git/internal/services/attendance"
	"github.com/claudineyveloso/bookernet.git/internal/services/bucket"
	"github.com/claudineyveloso/bookernet.git/internal/services/customer"
	"github.com/claudineyveloso/bookernet.git/internal/services/healthy"
	"github.com/claudineyveloso/bookernet.git/internal/services/insurance"
	"github.com/claudineyveloso/bookernet.git/internal/services/interval"
	"github.com/claudineyveloso/bookernet.git/internal/services/owner"
	typeservice "github.com/claudineyveloso/bookernet.git/internal/services/type_service"
	"github.com/claudineyveloso/bookernet.git/internal/services/user"
	"github.com/gorilla/handlers"
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
	ownerHandler := owner.NewHandler(ownerStore)
	ownerHandler.RegisterRoutes(r)

	customerStore := customer.NewStore(s.db)
	customerHandler := customer.NewHandler(customerStore)
	customerHandler.RegisterRoutes(r)

	typeServiceStore := typeservice.NewStore(s.db)
	typeServiceHandler := typeservice.NewHandler(typeServiceStore)
	typeServiceHandler.RegisterRoutes(r)

	intervalStore := interval.NewStore(s.db)
	intervalHandler := interval.NewHandler(intervalStore)
	intervalHandler.RegisterRoutes(r)

	attendanceStore := attendance.NewStore(s.db)
	attendanceHandler := attendance.NewHandler(attendanceStore)
	attendanceHandler.RegisterRoutes(r)

	insuranceStore := insurance.NewStore(s.db)
	insuranceHandler := insurance.NewHandler(insuranceStore)
	insuranceHandler.RegisterRoutes(r)

	fmt.Println("Server started on http://localhost:8080")
	//return http.ListenAndServe("localhost:8080", r)
	return http.ListenAndServe("localhost:8080",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(r))
}
