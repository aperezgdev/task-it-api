package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/aperezgdev/task-it-api/internal/application/task"
)

type taskPostRequest struct {
    Title string `json:"title"`
    Description string `json:"description"`
    Creator string `json:"creator"`
    Asigned string `json:"asigned"`
    StatusId string `json:"statusId"`
    BoardId string `json:"boardId"`
}

type taskPatchRequest struct {
    TaskId string `json:"taskId"`
    StatusId string `json:"statusId"`
}

type TaskController struct {
    logger slog.Logger
    creator task.TaskCreator
    mover task.TaskMover
}

func NewTaskController(logger slog.Logger, creator task.TaskCreator, mover task.TaskMover) TaskController {
    return TaskController{logger, creator, mover}
}

func (tc *TaskController) PostController(w http.ResponseWriter, r http.Request) {
    var taskRequest taskPostRequest
    err := json.NewDecoder(r.Body).Decode(&taskRequest)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    errCreator := tc.creator.Run(
        r.Context(), 
        taskRequest.Title, 
        taskRequest.Description, 
        taskRequest.Creator, 
        taskRequest.Asigned, 
        taskRequest.StatusId,
        taskRequest.BoardId,
    )

    if errCreator != nil {
        writeError(w, errCreator)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (tc *TaskController) PatchController(w http.ResponseWriter, r http.Request) {
    var taskPatchRequest taskPatchRequest
    taskId := r.PathValue("taskId")
    err := json.NewDecoder(r.Body).Decode(&taskPatchRequest)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    
    if taskId != taskPatchRequest.TaskId {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    errMover := tc.mover.Run(
        r.Context(), 
        taskPatchRequest.TaskId, 
        taskPatchRequest.StatusId,
    )

    if errMover != nil {
        writeError(w, errMover)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}