package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"

	"github.com/chriss-de/ssshare/internal/helpers"
	restV1 "github.com/chriss-de/ssshare/internal/rest/v1"

	"github.com/chriss-de/grouter/v1"
	"github.com/chriss-de/mux-middlewares/middlewares"
	"github.com/spf13/viper"
)

var (
	singletonOnce sync.Once
	httpServer    *http.Server
)

func Initialize() (err error) {
	singletonOnce.Do(func() {

		// http router
		router := grouter.NewRouter(viper.GetString("server.base_url"),
			middlewares.RealIP,
			middlewares.Logging,
			//localMiddleware.Recovery,
		)

		// healthz
		router.Get("/healthz").DoFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
		})

		//router.Get("{$}").DoFunc(func(writer http.ResponseWriter, request *http.Request) {
		//	slog.Info("TEST")
		//})
		router.Get(fmt.Sprintf("%s/{groupID}/{shareID}", viper.GetString("shares.url_path_prefix"))).DoFunc(getFile)
		router.Get(fmt.Sprintf("%s/{groupID}/{shareID}/", viper.GetString("shares.url_path_prefix"))).DoFunc(getFile)

		// REST API
		restV1.RegisterEndpoints(router.AddSubRouter("/api/v1"))

		// start http server
		httpServer = &http.Server{
			ErrorLog:       slog.NewLogLogger(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{}), slog.LevelError),
			Addr:           viper.GetString("server.listen_addr"),
			Handler:        router,
			MaxHeaderBytes: 1 << 20, // 1 MB
		}

		go func() {
			slog.Info("now handling requests", "listen_addr", httpServer.Addr)
			if err = httpServer.ListenAndServe(); err != nil {
				helpers.FatalError("server", "error", err.Error())
			}
		}()
	})
	return err
}
