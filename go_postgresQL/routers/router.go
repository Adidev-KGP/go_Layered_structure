package routers

import (
    "go-postgres/service"

    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    router.HandleFunc("/api/user/{id}", service.GetUser).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/user", service.GetAllUser).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newuser", service.CreateUser).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/user/{id}", service.UpdateUser).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deleteuser/{id}", service.DeleteUser).Methods("DELETE", "OPTIONS")

    return router
}

