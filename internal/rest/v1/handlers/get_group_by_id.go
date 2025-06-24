package handlers

import (
	"log/slog"
	"net/http"

	"github.com/chriss-de/ssshare/internal/helpers"
)

// GetGroupByID returns all roles filtered by query parameters
/*
//@Summary Returns all roles filtered by query parameters
//@Schemes
//@Description Returns all roles filtered by query parameters
//@Produce json
//@Param roleFilters query  restModel.RolesQueryParams false "filter"
//@Success 200 {object} restModel.Roles
//@Failure 401
//@Failure 500
//@Router /api/v5/roles [get]
//@Security ApiKeyAuth
*/
func GetGroupByID(w http.ResponseWriter, r *http.Request) {
	var (
		//err        error
		groupID string
	)

	groupID = r.PathValue("groupID")

	if _, iErr := helpers.WriteJSON(w, http.StatusOK, groupID); iErr != nil {
		slog.ErrorContext(r.Context(), "error writing response", "error", iErr)
	}
}
