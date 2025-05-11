package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type ClockworkApi struct {
	vipe   *viper.Viper
	Router *mux.Router
}

func NewClockworkApi(vipe *viper.Viper) *ClockworkApi {
	r := mux.NewRouter()

	api := &ClockworkApi{
		vipe:   vipe,
		Router: r,
	}

	api.addRoutes()

	// Apply middleware after routes are registered
	api.Router.Use(api.authMiddleware)

	return api
}

func (c *ClockworkApi) Serve() {
	port := c.vipe.GetString("server.port")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      c.Router,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Starting server on port %s\n", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed: %v", err)
	}
}
