package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/aperezgdev/task-it-api/internal/application/status"
)

type statusPostRequest struct {
	Title string `json:"title"`
	Board string `json:"board"`
	NextStatus []string `json:"nextStatus"`
	PreviousStatus []string `json:"previousStatus"`
}

type StatusController struct {
	logger slog.Logger
	creator status.StatusCreator
	remover status.StatusRemover
	finderByBoard status.StatusFinderByBoard
}

func NewStatusController(logger slog.Logger, creator status.StatusCreator, remover status.StatusRemover, finderByBoard status.StatusFinderByBoard) StatusController {
	return StatusController{logger, creator, remover, finderByBoard}
}

func (sc *StatusController) PostController(w http.ResponseWriter, r http.Request) {
	var statusRequest statusPostRequest
	err := json.NewDecoder(r.Body).Decode(&statusRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errCreator := sc.creator.Run(r.Context(), statusRequest.Title, statusRequest.Board, statusRequest.NextStatus, statusRequest.PreviousStatus)

	if errCreator != nil {
		writeError(w, errCreator)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (sc *StatusController) DeleteController(w http.ResponseWriter, r http.Request) {
	var statusId = r.PathValue("id")
	err := sc.remover.Run(r.Context(), statusId)
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (sc *StatusController) GetControllerByBoard(w http.ResponseWriter, r http.Request) {
	var boardId = r.PathValue("boardId")
	statuses, err := sc.finderByBoard.Run(r.Context(), boardId)
	if err != nil {
		writeError(w, err)
		return
	}

	err = json.NewEncoder(w).Encode(statuses)
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}