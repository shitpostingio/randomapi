package rest

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/shitpostingio/randomapi/backstore"
)

//Setup initialize the http server
func Setup(bindAddress string, bs *backstore.Backstore, allowedOrigins []string) *http.Server {

	ri := NewInterface(bs)

	router := mux.NewRouter()

	origins := handlers.AllowedOrigins(allowedOrigins)
	meths := handlers.AllowedMethods([]string{http.MethodGet, http.MethodOptions})
	heads := handlers.AllowedHeaders([]string{"x-user-platform", "x-user-id", "Content-Type", "content-type", "Origin"})

	router.HandleFunc("/random", ri.Random).Methods(http.MethodGet, http.MethodOptions)

	srv := &http.Server{
		Handler: handlers.CORS(meths, origins, heads)(router),
		Addr:    bindAddress,

		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	return srv
}
