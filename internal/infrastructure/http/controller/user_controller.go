package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/aperezgdev/task-it-api/internal/application/user"
)

type userPostRequest struct {
	Email string `json:"email"`
}

type UserController struct {
	logger slog.Logger
	creator user.UserCreator
}

func NewUserController(logger slog.Logger, creator user.UserCreator) UserController {
	return UserController{logger, creator}
}

func (uc *UserController) PostController(w http.ResponseWriter, r http.Request) {
	var userRequest userPostRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errCreator := uc.creator.Run(r.Context(), userRequest.Email)

	if errCreator != nil {
		writeError(w, errCreator)
		return
	}

	w.WriteHeader(http.StatusCreated)
}