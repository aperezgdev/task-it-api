package controller

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/application/task"
	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestPostController(t *testing.T) {
	t.Parallel()

	t.Run("should create task", func(t *testing.T) {
		taskRepository := new(repository.MockTaskRepository)
		boardRepository := new(repository.MockBoardRepository)
		userRepository := new(repository.MockUserRepository)
		statusRepository := new(repository.MockStatusRepository)
		taskCreator := task.NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)
		taskMover := task.NewTaskMover(*slog.Default(), taskRepository,statusRepository)
		remover := task.NewTaskRemover(*slog.Default(), taskRepository)

		taskRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)

		uuid, _ := uuid.NewV7()

		taskController := NewTaskController(*slog.Default(), taskCreator, taskMover, remover, task.TaskFinderByTeam{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(`{"title":"title","description":"description", "boardId":"` + uuid.String() + `","creator":"` + uuid.String() + `","asigned":"` + uuid.String() + `","statusId":"` + uuid.String() + `"}`)))
		taskController.PostController(w, *r)

		if w.Code != http.StatusCreated {
			t.Errorf("expected %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("should return error on bad request", func(t *testing.T) {
		taskRepository := new(repository.MockTaskRepository)
		boardRepository := new(repository.MockBoardRepository)
		userRepository := new(repository.MockUserRepository)
		statusRepository := new(repository.MockStatusRepository)
		taskCreator := task.NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)
		taskMover := task.NewTaskMover(*slog.Default(), taskRepository,statusRepository)
		taskRemover := task.NewTaskRemover(*slog.Default(), taskRepository)
		uuid, _ := uuid.NewV7()

		taskRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)

		taskController := NewTaskController(*slog.Default(), taskCreator, taskMover, taskRemover, task.TaskFinderByTeam{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(`{"a":"title","description":"description","creator":"` + uuid.String() + `","asigned":"` + uuid.String() + `","statusId":"` + uuid.String() + `"}`)))
		taskController.PostController(w, *r)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("should return error on board not found", func(t *testing.T) {
		taskRepository := new(repository.MockTaskRepository)
		boardRepository := new(repository.MockBoardRepository)
		userRepository := new(repository.MockUserRepository)
		statusRepository := new(repository.MockStatusRepository)
		taskCreator := task.NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)
		taskMover := task.NewTaskMover(*slog.Default(), taskRepository,statusRepository)
		taskRemover := task.NewTaskRemover(*slog.Default(), taskRepository)
		uuid, _ := uuid.NewV7()

		taskRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Board](), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)

		taskController := NewTaskController(*slog.Default(), taskCreator, taskMover, taskRemover, task.TaskFinderByTeam{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(`{"title":"title","description":"description", "boardId":"` + uuid.String() + `","creator":"` + uuid.String() + `","asigned":"` + uuid.String() + `","statusId":"` + uuid.String() + `"}`)))
		taskController.PostController(w, *r)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})

	t.Run("should return error on user not found", func(t *testing.T) {
		taskRepository := new(repository.MockTaskRepository)
		boardRepository := new(repository.MockBoardRepository)
		userRepository := new(repository.MockUserRepository)
		statusRepository := new(repository.MockStatusRepository)
		taskCreator := task.NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)
		taskMover := task.NewTaskMover(*slog.Default(), taskRepository,statusRepository)
		taskRemover := task.NewTaskRemover(*slog.Default(), taskRepository)
		uuid, _ := uuid.NewV7()

		taskRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.User](), errors.ErrNotExist)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)

		taskController := NewTaskController(*slog.Default(), taskCreator, taskMover, taskRemover, task.TaskFinderByTeam{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer([]byte(`{"title":"title","description":"description", "boardId":"` + uuid.String() + `","creator":"` + uuid.String() + `","asigned":"` + uuid.String() + `","statusId":"` + uuid.String() + `"}`)))
		taskController.PostController(w, *r)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestPatchController(t *testing.T) {
	t.Parallel()

	t.Run("should create a task on valid", func(t *testing.T) {
		taskRepository := new(repository.MockTaskRepository)
		boardRepository := new(repository.MockBoardRepository)
		userRepository := new(repository.MockUserRepository)
		statusRepository := new(repository.MockStatusRepository)
		taskCreator := task.NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)
		taskMover := task.NewTaskMover(*slog.Default(), taskRepository,statusRepository)
		taskRemover := task.NewTaskRemover(*slog.Default(), taskRepository)
		uuid, _ := uuid.NewV7()

		taskRepository.On("Update", mock.Anything, mock.Anything).Return(nil)
		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Task{}), nil)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)

		taskController := NewTaskController(*slog.Default(), taskCreator, taskMover, taskRemover, task.TaskFinderByTeam{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPatch, "/tasks/", bytes.NewBuffer([]byte(`{"taskId":"` + uuid.String() + `","statusId":"` + uuid.String() + `"}`)))
		r.SetPathValue("taskId", uuid.String())
		taskController.PatchController(w, *r)

		if w.Code != http.StatusNoContent {
			t.Errorf("expected %d, got %d", http.StatusNoContent, w.Code)
		}

		taskRepository.AssertNumberOfCalls(t, "Update", 1)
		taskRepository.AssertNumberOfCalls(t, "Find", 1)
		statusRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return error on task not found", func(t *testing.T) {
		taskRepository := new(repository.MockTaskRepository)
		boardRepository := new(repository.MockBoardRepository)
		userRepository := new(repository.MockUserRepository)
		statusRepository := new(repository.MockStatusRepository)
		taskCreator := task.NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)
		taskMover := task.NewTaskMover(*slog.Default(), taskRepository,statusRepository)
		taskRemover := task.NewTaskRemover(*slog.Default(), taskRepository)
		uuid, _ := uuid.NewV7()

		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Task](), errors.ErrNotExist)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Status{}), nil)

		taskController := NewTaskController(*slog.Default(), taskCreator, taskMover, taskRemover, task.TaskFinderByTeam{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPatch, "/tasks", bytes.NewBuffer([]byte(`{"taskId":"` + uuid.String() + `","statusId":"` + uuid.String() + `"}`)))
		r.SetPathValue("taskId", uuid.String())
		taskController.PatchController(w, *r)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})

	t.Run("should return error on status not found", func(t *testing.T) {
		taskRepository := new(repository.MockTaskRepository)
		boardRepository := new(repository.MockBoardRepository)
		userRepository := new(repository.MockUserRepository)
		statusRepository := new(repository.MockStatusRepository)
		taskCreator := task.NewTaskCreator(*slog.Default(), boardRepository, userRepository, taskRepository)
		taskMover := task.NewTaskMover(*slog.Default(), taskRepository,statusRepository)
		taskRemover := task.NewTaskRemover(*slog.Default(), taskRepository)
		uuid, _ := uuid.NewV7()

		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Task{}), nil)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		statusRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Status](), errors.ErrNotExist)

		taskController := NewTaskController(*slog.Default(), taskCreator, taskMover, taskRemover, task.TaskFinderByTeam{})
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPatch, "/tasks", bytes.NewBuffer([]byte(`{"taskId":"` + uuid.String() + `","statusId":"` + uuid.String() + `"}`)))
		r.SetPathValue("taskId", uuid.String())
		taskController.PatchController(w, *r)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestDeleteController(t *testing.T) {
	t.Parallel()

	t.Run("should delete task", func(t *testing.T) {
		taskRepository := new(repository.MockTaskRepository)
		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Task{}), nil)
		taskRepository.On("Delete", mock.Anything, mock.Anything).Return(nil)
		taskController := NewTaskController(*slog.Default(), task.NewTaskCreator(*slog.Default(), nil, nil, nil), task.NewTaskMover(*slog.Default(), nil, nil), task.NewTaskRemover(*slog.Default(), taskRepository), task.TaskFinderByTeam{})
		uuid, _ := uuid.NewV7()

		req := httptest.NewRequest(http.MethodDelete, "/"+uuid.String(), nil)
		req.SetPathValue("id", uuid.String())
		w := httptest.NewRecorder()
		taskController.DeleteController(w, *req)

		if w.Code != http.StatusNoContent {
			t.Errorf("expected %d, got %d", http.StatusNoContent, w.Code)
		}
	})

	t.Run("should return not found", func(t *testing.T) {
		taskRepository := new(repository.MockTaskRepository)
		taskRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Task](), nil)
		taskController := NewTaskController(*slog.Default(), task.NewTaskCreator(*slog.Default(), nil, nil, nil), task.NewTaskMover(*slog.Default(), nil, nil), task.NewTaskRemover(*slog.Default(), taskRepository), task.TaskFinderByTeam{})
		uuid, _ := uuid.NewV7()

		req := httptest.NewRequest(http.MethodDelete, "/"+uuid.String(), nil)
		req.SetPathValue("id", uuid.String())
		w := httptest.NewRecorder()
		taskController.DeleteController(w, *req)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestGetControllerByTeam(t *testing.T) {
	t.Parallel()
	
	t.Run("should return task by team", func(t *testing.T) {
		writter := httptest.NewRecorder()
		taskRepository := repository.MockTaskRepository{}
		taskRepository.On("FindByTeam", mock.Anything, mock.Anything).Return(pkg.NewOptional([]model.Task{}), nil)
		boardRepository := repository.MockBoardRepository{}
		userRepository := repository.MockUserRepository{}
		statusRepository := repository.MockStatusRepository{}
		teamRepository := repository.MockTeamRepository{}
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		taskCreator := task.NewTaskCreator(*slog.Default(), &boardRepository, &userRepository, &taskRepository)
		taskMover := task.NewTaskMover(*slog.Default(), &taskRepository, &statusRepository)
		taskRemover := task.NewTaskRemover(*slog.Default(), &taskRepository)
		taskFinderByTeam := task.NewTaskFinderByTeam(*slog.Default(), &taskRepository, &teamRepository)
		taskController := NewTaskController(*slog.Default(), taskCreator, taskMover, taskRemover, taskFinderByTeam)
		uuid, _ := uuid.NewV7()

		req := httptest.NewRequest(http.MethodGet, "/"+uuid.String(), nil)
		req.SetPathValue("boardId", uuid.String())
		taskController.GetControllerByTeam(writter, *req)

		if writter.Code != http.StatusOK {
			t.Errorf("expected %d, got %d", http.StatusOK, writter.Code)
		}
	})

	t.Run("should return not found", func(t *testing.T) {
		writter := httptest.NewRecorder()
		taskRepositoryMock := repository.MockTaskRepository{}
		boardRepositoryMock := repository.MockBoardRepository{}
		userRepositoryMock := repository.MockUserRepository{}
		statusRepositoryMock := repository.MockStatusRepository{}
		teamRepositoryMock := repository.MockTeamRepository{}
		teamRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), errors.ErrNotExist)
		taskCreator := task.NewTaskCreator(*slog.Default(), &boardRepositoryMock, &userRepositoryMock, &taskRepositoryMock)
		taskMover := task.NewTaskMover(*slog.Default(), &taskRepositoryMock, &statusRepositoryMock)
		taskRemover := task.NewTaskRemover(*slog.Default(), &taskRepositoryMock)
		taskFinderByTeam := task.NewTaskFinderByTeam(*slog.Default(), &taskRepositoryMock, &teamRepositoryMock)
		taskController := NewTaskController(*slog.Default(), taskCreator, taskMover, taskRemover, taskFinderByTeam)
		uuid, _ := uuid.NewV7()

		req := httptest.NewRequest(http.MethodGet, "/"+uuid.String(), nil)
		req.SetPathValue("boardId", uuid.String())
		taskController.GetControllerByTeam(writter, *req)

		if writter.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, writter.Code)
		}
	})
}