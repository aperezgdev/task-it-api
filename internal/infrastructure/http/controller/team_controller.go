package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/aperezgdev/task-it-api/internal/application/team"
)

type TeamController struct {
	logger slog.Logger
	creator team.TeamCreator
	removeMember team.RemoverMember
	addMember team.TeamAddMember
}

type teamPostRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Owner string `json:"owner"`
}

type teamPostMemberRequest struct {
	Id string `json:"id"`
}

func NewTeamController(logger slog.Logger, creator team.TeamCreator, removeMember team.RemoverMember, addMember team.TeamAddMember) TeamController {
	return TeamController{logger, creator, removeMember, addMember}
}

func (tc *TeamController) PostTeam(w http.ResponseWriter, r *http.Request) {
	var teamRequest teamPostRequest
	err := json.NewDecoder(r.Body).Decode(&teamRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errCreator := tc.creator.Run(
		r.Context(), 
		teamRequest.Title, 
		teamRequest.Owner, 
		teamRequest.Description,
	)
	
	if errCreator != nil {
		writeError(w, errCreator)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (tc *TeamController) PostMemberController(w http.ResponseWriter, r http.Request) {
	var memberRequest teamPostMemberRequest
	var teamId = r.PathValue("teamId")
	err := json.NewDecoder(r.Body).Decode(&memberRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	errCreator := tc.addMember.Run(
		r.Context(),
		teamId,
		memberRequest.Id,
	)
	
	if errCreator != nil {
		writeError(w, errCreator)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (tc *TeamController) DeleteMemberController(w http.ResponseWriter, r http.Request) {
	var teamId = r.PathValue("teamId")
	var memberId = r.PathValue("memberId")

	errCreator := tc.removeMember.Run(
		r.Context(),
		teamId,
		memberId,
	)
	
	if errCreator != nil {
		writeError(w, errCreator)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}