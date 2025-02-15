package controller

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/application/team"
	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestPostTeamController(t *testing.T) {
	t.Parallel()

	t.Run("should create team", func(t *testing.T) {
		teamRepository := new(repository.MockTeamRepository)
		boardRepository := new(repository.MockBoardRepository)
		userRepository := new(repository.MockUserRepository)
		teamCreator := team.NewTeamCreator(*slog.Default(), teamRepository, userRepository)
		teamAddMember := team.NewTeamAddMember(*slog.Default(), teamRepository, userRepository)
		teamRemoveMember := team.NewRemoverMember(*slog.Default(), teamRepository, userRepository)

		teamRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)

		uuid, _ := uuid.NewV7()

		teamController := NewTeamController(*slog.Default(), teamCreator, teamRemoveMember, teamAddMember)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/teams", bytes.NewBuffer([]byte(`{"title":"title","description":"description", "boardId":"` + uuid.String() + `","owner":"` + uuid.String() + `","asigned":"` + uuid.String() + `","statusId":"` + uuid.String() + `"}`)))
		teamController.PostTeam(w, r)

		if w.Code != http.StatusCreated {
			t.Errorf("expected %d, got %d", http.StatusCreated, w.Code)
		}
	})

	t.Run("should return error on bad request", func(t *testing.T) {
		teamRepository := new(repository.MockTeamRepository)
		boardRepository := new(repository.MockBoardRepository)
		userRepository := new(repository.MockUserRepository)
		teamCreator := team.NewTeamCreator(*slog.Default(), teamRepository, userRepository)
		teamAddMember := team.NewTeamAddMember(*slog.Default(), teamRepository, userRepository)
		teamRemoveMember := team.NewRemoverMember(*slog.Default(), teamRepository, userRepository)
		uuid, _ := uuid.NewV7()

		teamRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)

		teamController := NewTeamController(*slog.Default(), teamCreator, teamRemoveMember, teamAddMember)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/teams", bytes.NewBuffer([]byte(`{"a":"title","description":"description","creator":"` + uuid.String() + `","asigned":"` + uuid.String() + `","statusId":"` + uuid.String() + `"}`)))
		teamController.PostTeam(w, r)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("should return error on user not found", func(t *testing.T) {
		teamRepository := new(repository.MockTeamRepository)
		boardRepository := new(repository.MockBoardRepository)
		userRepository := new(repository.MockUserRepository)
		teamCreator := team.NewTeamCreator(*slog.Default(), teamRepository, userRepository)
		teamAddMember := team.NewTeamAddMember(*slog.Default(), teamRepository, userRepository)
		teamRemoveMember := team.NewRemoverMember(*slog.Default(), teamRepository, userRepository)
		uuid, _ := uuid.NewV7()

		teamRepository.On("Save", mock.Anything, mock.Anything).Return(nil)
		boardRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Board{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.User](), errors.ErrNotExist)

		teamController := NewTeamController(*slog.Default(), teamCreator, teamRemoveMember, teamAddMember)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/teams", bytes.NewBuffer([]byte(`{"title":"title","description":"description","owner":"` + uuid.String() + `","asigned":"` + uuid.String() + `","statusId":"` + uuid.String() + `"}`)))
		teamController.PostTeam(w, r)

		if w.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestAddMemberController(t *testing.T) {
	t.Parallel()

	t.Run("should add member", func(t *testing.T) {
		writter := httptest.NewRecorder()
		teamRepository := repository.MockTeamRepository{}
		userRepository := repository.MockUserRepository{}
		teamCreator := team.NewTeamCreator(*slog.Default(), &teamRepository, &userRepository)
		teamAddMember := team.NewTeamAddMember(*slog.Default(), &teamRepository, &userRepository)
		teamRemoveMember := team.NewRemoverMember(*slog.Default(), &teamRepository, &userRepository)
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)
		teamRepository.On("Update", mock.Anything, mock.Anything).Return(nil)

		uuid, _ := uuid.NewV7()

		teamController := NewTeamController(*slog.Default(), teamCreator, teamRemoveMember, teamAddMember)
		r := httptest.NewRequest(http.MethodPost, "/teams/"+uuid.String()+"/members", bytes.NewBuffer([]byte(`{"id":"` + uuid.String() + `"}`)))
		r.SetPathValue("teamId", uuid.String())
		r.SetPathValue("memberId", uuid.String())
		teamController.PostMemberController(writter, r)

		if writter.Code != http.StatusCreated {
			t.Errorf("expected %d, got %d", http.StatusCreated, writter.Code)
		}
	})

	t.Run("should return error on team not found", func(t *testing.T) {
		writter := httptest.NewRecorder()
		teamRepository := repository.MockTeamRepository{}
		userRepository := repository.MockUserRepository{}
		teamCreator := team.NewTeamCreator(*slog.Default(), &teamRepository, &userRepository)
		teamAddMember := team.NewTeamAddMember(*slog.Default(), &teamRepository, &userRepository)
		teamRemoveMember := team.NewRemoverMember(*slog.Default(), &teamRepository, &userRepository)
		uuid, _ := uuid.NewV7()

		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), errors.ErrNotExist)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)

		teamController := NewTeamController(*slog.Default(), teamCreator, teamRemoveMember, teamAddMember)
		r := httptest.NewRequest(http.MethodPost, "/teams/"+uuid.String()+"/members", bytes.NewBuffer([]byte(`{"id":"` + uuid.String() + `"}`)))
		r.SetPathValue("teamId", uuid.String())
		r.SetPathValue("memberId", uuid.String())
		teamController.PostMemberController(writter, r)

		if writter.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, writter.Code)
		}
	})

	t.Run("should return error on user not found", func(t *testing.T) {
		writter := httptest.NewRecorder()
		teamRepository := repository.MockTeamRepository{}
		userRepository := repository.MockUserRepository{}
		teamCreator := team.NewTeamCreator(*slog.Default(), &teamRepository, &userRepository)
		teamAddMember := team.NewTeamAddMember(*slog.Default(), &teamRepository, &userRepository)
		teamRemoveMember := team.NewRemoverMember(*slog.Default(), &teamRepository, &userRepository)
		uuid, _ := uuid.NewV7()

		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.User](), errors.ErrNotExist)

		teamController := NewTeamController(*slog.Default(), teamCreator, teamRemoveMember, teamAddMember)
		r := httptest.NewRequest(http.MethodPost, "/teams/"+uuid.String()+"/members", bytes.NewBuffer([]byte(`{"id":"` + uuid.String() + `"}`)))
		teamController.PostMemberController(writter, r)

		if writter.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, writter.Code)
		}
	})
}

func TestDeleteMemberController(t *testing.T) {
	t.Parallel()

	t.Run("should delete member", func(t *testing.T) {
		writter := httptest.NewRecorder()
		teamRepository := repository.MockTeamRepository{}
		userRepository := repository.MockUserRepository{}
		teamCreator := team.NewTeamCreator(*slog.Default(), &teamRepository, &userRepository)
		teamAddMember := team.NewTeamAddMember(*slog.Default(), &teamRepository, &userRepository)
		teamRemoveMember := team.NewRemoverMember(*slog.Default(), &teamRepository, &userRepository)

		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		userRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.User{}), nil)

		uuid, _ := uuid.NewV7()

		teamController := NewTeamController(*slog.Default(), teamCreator, teamRemoveMember, teamAddMember)
		r := httptest.NewRequest(http.MethodDelete, "/teams/"+uuid.String()+"/members/"+uuid.String(), nil)
		r.SetPathValue("teamId", uuid.String())
		r.SetPathValue("memberId", uuid.String())
		teamController.DeleteMemberController(writter, r)

		if writter.Code != http.StatusNoContent {
			t.Errorf("expected %d, got %d", http.StatusNoContent, writter.Code)
		}
	})
	
	t.Run("should return not found", func(t *testing.T) {
		writter := httptest.NewRecorder()
		teamRepositoryMock := repository.MockTeamRepository{}
		teamRepositoryMock.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), errors.ErrNotExist)
		teamRepositoryMock.On("FindByMember", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), errors.ErrNotExist)		
		teamController := NewTeamController(*slog.Default(), team.NewTeamCreator(*slog.Default(), &teamRepositoryMock, nil), team.NewRemoverMember(*slog.Default(), &teamRepositoryMock, nil), team.NewTeamAddMember(*slog.Default(), &teamRepositoryMock, nil))	
		uuid, _ := uuid.NewV7()
		
		r := httptest.NewRequest(http.MethodDelete, "/teams/"+uuid.String()+"/members/"+uuid.String(), nil)
		r.SetPathValue("teamId", uuid.String())
		r.SetPathValue("memberId", uuid.String())
		teamController.DeleteMemberController(writter, r)

		if writter.Code != http.StatusNotFound {
			t.Errorf("expected %d, got %d", http.StatusNotFound, writter.Code)
		}
	})
}