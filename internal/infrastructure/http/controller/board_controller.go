package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/aperezgdev/task-it-api/internal/application/board"
)

type boardPostRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Owner string `json:"owner"`
	Team string `json:"team"`
}

type BoardController struct {
	logger slog.Logger
	creator board.BoardCreator
	remover board.BoardRemover
}

func NewBoardController(logger slog.Logger, creator board.BoardCreator, remover board.BoardRemover) BoardController {
	return BoardController{logger, creator, remover}
}

func (bc *BoardController) PostController(w http.ResponseWriter, r http.Request) {
	var boardRequest boardPostRequest
	err := json.NewDecoder(r.Body).Decode(&boardRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errCreator := bc.creator.Run(r.Context(), boardRequest.Title, boardRequest.Description, boardRequest.Owner, boardRequest.Team)

	if errCreator != nil {
		writeError(w, errCreator)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (bc *BoardController) DeleteController(w http.ResponseWriter, r http.Request) {
	var boardId = r.PathValue("id")
	err := bc.remover.Run(r.Context(), boardId)
	if err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}