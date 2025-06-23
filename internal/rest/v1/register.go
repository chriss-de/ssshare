package v1

import (
	"net/http"

	"github.com/chriss-de/ssshare/internal/rest/v1/handlers"

	"github.com/chriss-de/grouter/v1"
)

// RegisterEndpoints registers endpoints to chi-router
// We do all setup stuff for those endpoints here - AUTH, cache, req-param injection
func RegisterEndpoints(r *grouter.Router) {
	//r.AddMiddlewares(setContentType)

	r.Get("/s/{groupID}/{shareID}").DoFunc(handlers.GetFile)
	r.Get("/s/{groupID}/{shareID}/").DoFunc(handlers.GetFile)
	// .With(httpin.NewInput(model.XXQueryParams{}))

}

func setContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
