package board

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aperezgdev/task-it-api/internal/domain/errors"
	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/pkg"
	"github.com/stretchr/testify/mock"
)

func TestBoardFinderByTeam(t *testing.T) {
	t.Run("should return board by team", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("FindByTeam", mock.Anything, mock.Anything).Return(pkg.NewOptional([]model.Board{}), nil)
		teamRepository := &repository.MockTeamRepository{}
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		boardFinder := NewBoardFinderByTeam(*slog.Default(), boardRepository, teamRepository)

		boards, err := boardFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}

		if len(boards) != 0 {
			t.Errorf("Expected empty array, got %v", boards)
		}

		boardRepository.AssertNumberOfCalls(t, "FindByTeam", 1)
		teamRepository.AssertNumberOfCalls(t, "Find", 1)
	})

	t.Run("should return error on invalid team id", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("FindByTeam", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[[]model.Board](), nil)
		teamRepository := &repository.MockTeamRepository{}
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[model.Team](), nil)
		boardFinder := NewBoardFinderByTeam(*slog.Default(), boardRepository, teamRepository)

		boards, err := boardFinder.Run(context.Background(), "")
		if err == nil {
			t.Errorf("Expected error on invalid team id, got nil")
		}

		if len(boards) != 0 {
			t.Errorf("Expected empty array, got %v", boards)
		}

		boardRepository.AssertNumberOfCalls(t, "FindByTeam", 0)
		teamRepository.AssertNumberOfCalls(t, "Find", 0)
	})

	t.Run("should return error on board not found", func(t *testing.T) {
		boardRepository := &repository.MockBoardRepository{}
		boardRepository.On("FindByTeam", mock.Anything, mock.Anything).Return(pkg.EmptyOptional[[]model.Board](), errors.ErrNotExist)
		teamRepository := &repository.MockTeamRepository{}
		teamRepository.On("Find", mock.Anything, mock.Anything).Return(pkg.NewOptional(model.Team{}), nil)
		boardFinder := NewBoardFinderByTeam(*slog.Default(), boardRepository, teamRepository)

		boards, err := boardFinder.Run(context.Background(), "01946ba3-ee73-76e6-83a9-33f87a35d6e9")
		if err == nil {
			t.Errorf("Expected error on board not found, got nil")
		}

		if len(boards) != 0 {
			t.Errorf("Expected empty array, got %v", boards)
		}

		boardRepository.AssertNumberOfCalls(t, "FindByTeam", 1)
		teamRepository.AssertNumberOfCalls(t, "Find", 1)
	})
}