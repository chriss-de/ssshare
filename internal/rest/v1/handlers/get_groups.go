package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/chriss-de/ssshare/internal/helpers"
	restModel "github.com/chriss-de/ssshare/internal/rest/v1/model"

	"github.com/ggicci/httpin"
)

// GetGroups returns all roles filtered by query parameters
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
func GetGroups(w http.ResponseWriter, r *http.Request) {
	var (
		//err        error
		rolesQuery *restModel.GroupsQueryParams
	)

	rolesQuery = r.Context().Value(httpin.Input).(*restModel.GroupsQueryParams)

	restGroups := restModel.Groups{
		Groups: []restModel.Group{},
		Links: &restModel.GroupsLinks{
			SelfLink: restModel.SelfLink{
				Self: restModel.Href{Href: fmt.Sprintf("%s://%s%s?%s", helpers.GetScheme(r), helpers.GetHost(r), r.URL.Path, r.URL.Query().Encode())},
			},
		},
		Paging: restModel.NewPaging(0, rolesQuery.First, rolesQuery.Skip),
	}

	if _, iErr := helpers.WriteJSON(w, http.StatusOK, restGroups); iErr != nil {
		slog.ErrorContext(r.Context(), "error writing response", "error", iErr)
	}
}
