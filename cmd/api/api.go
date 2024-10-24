package api

import (
	"log"
	"net/http"

	"github.com/Harsh-710/hospital-management/services/appointment"
	"github.com/Harsh-710/hospital-management/services/user"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

// NewAPIServer initializes a new API server with GORM DB connection
func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the API server and sets up routes
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Initialize user store and handler
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// Initialize product store and handler
	patientStore := patient.NewStore(s.db)
	// patientHandler := patient.NewHandler(productStore, userStore)
	patientHandler.RegisterRoutes(subrouter)

	// Initialize order store
	orderStore := appointment.NewStore(s.db)

	// Initialize cart handler with dependencies
	// cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	// cartHandler.RegisterRoutes(subrouter)

	// Serve static files from the "static" directory
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.Println("Server is listening on", s.addr)

	// Start the HTTP server
	return http.ListenAndServe(s.addr, router)
}
