package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/geisonbiazus/openchat/internal/openchat"
	"github.com/geisonbiazus/openchat/internal/openchat/services"
)

type CreateUserService interface {
	Run(input services.CreateUserInput) services.CreateUserOutput
}

type CreateUserHandler struct {
	service CreateUserService
}

func NewCreateUserHandler(service CreateUserService) *CreateUserHandler {
	return &CreateUserHandler{
		service: service,
	}
}

func (h *CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	output := h.service.Run(h.serviceInput(r))

	w.Header().Set("Content-Type", "application/json")

	if output.Status == services.StatusSuccess {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(h.buildSuccessJSON(output))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(h.buildErrorJSON(output))
	}
}

func (h *CreateUserHandler) serviceInput(r *http.Request) services.CreateUserInput {
	body := struct {
		Username string `json:"username"`
		Password string `json:"password"`
		About    string `json:"about"`
	}{}
	json.NewDecoder(r.Body).Decode(&body)

	return services.CreateUserInput{
		Username: body.Username,
		Password: body.Password,
		About:    body.About,
	}
}

type createUserSuccessJSON struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	About    string `json:"about"`
}

func (h *CreateUserHandler) buildSuccessJSON(output services.CreateUserOutput) createUserSuccessJSON {
	return createUserSuccessJSON{
		ID:       output.ID,
		Username: output.Username,
		About:    output.About,
	}
}

type createUserErrorJSON struct {
	Errors []openchat.Error `json:"errors"`
}

func (h *CreateUserHandler) buildErrorJSON(output services.CreateUserOutput) createUserErrorJSON {
	return createUserErrorJSON{
		Errors: output.Errors,
	}
}
