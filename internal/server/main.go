package server

import (
	"fmt"
	"github.com/chriss-de/ssshare/internal/helpers"
	"log/slog"
	"net/http"
	"sync"

	localMiddleware "github.com/chriss-de/ssshare/internal/middleware"
	restV1 "github.com/chriss-de/ssshare/internal/rest/v1"

	"github.com/chriss-de/grouter/v1"
	"github.com/chriss-de/mux-middlewares/middlewares"
	"github.com/spf13/viper"
)

var (
	singletonOnce      sync.Once
	httpServer         *http.Server
	urlPathSharePrefix = "/s"
)

func Initialize() (err error) {
	singletonOnce.Do(func() {

		// http router
		router := grouter.NewRouter(viper.GetString("server.baseUrl"),
			middlewares.RealIP,
			middlewares.Logging,
			localMiddleware.Recovery,
		)

		// healthz
		router.Get("/healthz").DoFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
		})

		router.Get(fmt.Sprintf("%s/{groupID}/{shareID}", urlPathSharePrefix)).DoFunc(getFile)
		router.Get(fmt.Sprintf("%s/{groupID}/{shareID}/", urlPathSharePrefix)).DoFunc(getFile)

		// REST API
		restV1.RegisterEndpoints(router.AddSubRouter("/api/v1"))

		// start http server
		httpServer = &http.Server{
			Addr:           ":8080",
			Handler:        router.GetServeMux(),
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
