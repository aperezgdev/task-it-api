package task

import (
	"context"
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/domain/model"
	"github.com/aperezgdev/task-it-api/internal/domain/repository"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/board"
	"github.com/aperezgdev/task-it-api/internal/domain/use_case/user"
)

type TaskCreator struct {
	logger slog.Logger
	boardFinder board.BoardFinder
	userFinder user.UserFinder
	taskRepository repository.TaskRepository
}

func NewTaskCreator(logger slog.Logger, boardRepository repository.BoardRepository, userRepository repository.UserRepository, taskRepository repository.TaskRepository) TaskCreator {
	return TaskCreator{
		logger: logger, 
		boardFinder: board.NewBoardFinder(logger, boardRepository),
		userFinder: user.NewUserFinder(logger, userRepository),
		taskRepository: taskRepository,
	}
}

func (tc *TaskCreator) Run(ctx context.Context,title, description, creator, asigned, statusId, boardId string) error {
	tc.logger.Info(
			"TaskCreator - Run - Params: ", 
			slog.Any("title", title), 
			slog.Any("description", description),
			slog.Any("creator", creator),
			slog.Any("asigned", asigned),
			slog.Any("statusId", statusId),
		)
	_, errBoard := tc.boardFinder.Run(ctx, boardId)
	if errBoard != nil {
		tc.logger.Info("TaskCreator - Run - Board not found")
		return errBoard
	}

	_, errUser := tc.userFinder.Run(ctx, creator)
	if errUser != nil {
		tc.logger.Info("TaskCreator - Run - User not found")
		return errUser
	}

	task, err := model.NewTask(title, description, creator, asigned, statusId, boardId)
	if err != nil {
		tc.logger.Info("TaskCreator - Run - Error trying to create task")
		return err
	}

	return tc.taskRepository.Save(ctx, task)	
}