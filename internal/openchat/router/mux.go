package router

import (
	"net/http"

	"github.com/geisonbiazus/openchat/internal/openchat/handlers"
	"github.com/geisonbiazus/openchat/internal/openchat/repositories"
	"github.com/geisonbiazus/openchat/internal/openchat/services"
	"github.com/geisonbiazus/openchat/internal/openchat/support"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	idGen := support.NewUUIDGenerator()
	userRep := repositories.NewInMemoryUserRepository()
	createUserService := services.NewCreateUserService(idGen, userRep)
	createUserHandler := handlers.NewCreateUserHandler(createUserService)

	mux.Handle("/users/", createUserHandler)

	return mux
}
