package v1

import (
	"net/http"

	"github.com/chriss-de/ssshare/internal/rest/v1/handlers"
	"github.com/chriss-de/ssshare/internal/rest/v1/model"

	"github.com/chriss-de/grouter/v1"
	"github.com/ggicci/httpin"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

// RegisterEndpoints registers endpoints to chi-router
// We do all setup stuff for those endpoints here - AUTH, cache, req-param injection
func RegisterEndpoints(r *grouter.Router) {
	_cors := cors.New(cors.Options{
		AllowedOrigins:   viper.GetStringSlice("server.cors.allowed_origins"),
		AllowedMethods:   viper.GetStringSlice("server.cors.allowed_methods"),
		AllowedHeaders:   viper.GetStringSlice("server.cors.allowed_headers"),
		ExposedHeaders:   viper.GetStringSlice("server.cors.exposed_headers"),
		MaxAge:           viper.GetInt("server.cors.max_age"),
		AllowCredentials: viper.GetBool("server.cors.allow_credentials"),
		Debug:            viper.GetBool("server.cors.debug"),
	})

	r.AddMiddlewares(setContentType, _cors.Handler)

	// routes
	r.Get("/groups").With(httpin.NewInput(model.GroupsQueryParams{})).DoFunc(handlers.GetGroups)
	r.Get("/groups/{groupID}").DoFunc(handlers.GetGroupByID)
}

func setContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
