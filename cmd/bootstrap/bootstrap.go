package bootstrap

import (
	"log/slog"

	"github.com/aperezgdev/task-it-api/internal/application/board"
	"github.com/aperezgdev/task-it-api/internal/application/status"
	"github.com/aperezgdev/task-it-api/internal/application/task"
	"github.com/aperezgdev/task-it-api/internal/application/team"
	"github.com/aperezgdev/task-it-api/internal/application/user"
	"github.com/aperezgdev/task-it-api/internal/infrastructure/config"
	"github.com/aperezgdev/task-it-api/internal/infrastructure/http"
	"github.com/aperezgdev/task-it-api/internal/infrastructure/http/controller"
	"github.com/aperezgdev/task-it-api/internal/infrastructure/repository/local"
)

func Run() error {
	logger := slog.Default()
	config := config.NewConfig(logger)
	server := http.NewServer(logger, config)

	userRepository := local.NewUserRepository(*logger)
	boardRepository := local.NewBoardRepository(*logger)
	teamRepository := local.NewTeamRepository(*logger)
	taskRepository := local.NewTaskRepository(*logger)
	statusRepository := local.NewStatusRepository(*logger)

	userCreator := user.NewUserCreator(*logger, userRepository)

	teamCreator := team.NewTeamCreator(*logger, teamRepository, userRepository)
	teamRemoveMember := team.NewRemoverMember(*logger, teamRepository, userRepository)
	teamAddMember := team.NewTeamAddMember(*logger, teamRepository, userRepository)
	
	taskCreator := task.NewTaskCreator(*logger, boardRepository, userRepository, taskRepository)
	taskMover := task.NewTaskMover(*logger, taskRepository, statusRepository)
	taskRemover := task.NewTaskRemover(*logger, taskRepository)
	taskFinderByTeam := task.NewTaskFinderByTeam(*logger, taskRepository, teamRepository)

	statusCreator := status.NewStatusCreator(*logger, statusRepository, boardRepository)
	statusRemover := status.NewStatusRemover(*logger, statusRepository)
	statusFinderByBoard := status.NewStatusFinderByBoard(*logger, statusRepository, boardRepository)

	boardCreator := board.NewBoardCreator(*logger, boardRepository, userRepository, teamRepository)
	boardRemover := board.NewBoardRemover(*logger, boardRepository)
	boardFinderByTeam := board.NewBoardFinderByTeam(*logger, boardRepository, teamRepository)

	healthController := controller.NewHealthController(*logger)
	userController := controller.NewUserController(*logger, userCreator)
	boardController := controller.NewBoardController(*logger, boardCreator, boardRemover, boardFinderByTeam)
	teamController := controller.NewTeamController(*logger, teamCreator, teamRemoveMember, teamAddMember)
	taskController := controller.NewTaskController(*logger, taskCreator, taskMover, taskRemover, taskFinderByTeam)
	statusController := controller.NewStatusController(*logger, statusCreator, statusRemover, statusFinderByBoard)

	server.AddHandler("GET /health", healthController.GetHealth)
	server.AddHandler("POST /users", userController.PostController)
	server.AddHandler("POST /boards", boardController.PostController)
	server.AddHandler("POST /teams", teamController.PostTeam)
	server.AddHandler("POST /teams/{teamId}/members", teamController.PostMemberController)
	server.AddHandler("DELETE /teams/{teamId}/members/{memberId}", teamController.DeleteMemberController)
	server.AddHandler("POST /tasks", taskController.PostController)
	server.AddHandler("PATCH /tasks/{taskId}", taskController.PatchController)
	server.AddHandler("DELETE /tasks/{taskId}", taskController.DeleteController)
	server.AddHandler("GET /tasks/{boardId}", taskController.GetControllerByTeam)
	server.AddHandler("POST /statuses", statusController.PostController)
	server.AddHandler("DELETE /statuses/{statusId}", statusController.DeleteController)
	server.AddHandler("GET /boards/{boardId}/statuses", statusController.GetControllerByBoard)

	return server.Start()
}