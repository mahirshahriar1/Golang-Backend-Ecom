package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	cart "github.com/mahirshahriar1/Golang-Backend-API/service/carts"
	"github.com/mahirshahriar1/Golang-Backend-API/service/order"
	"github.com/mahirshahriar1/Golang-Backend-API/service/product"
	"github.com/mahirshahriar1/Golang-Backend-API/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore, userStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(productStore, orderStore, userStore)
	cartHandler.RegisterRoutes(subrouter)


	log.Println("Server is running on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
