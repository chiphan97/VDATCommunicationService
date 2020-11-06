package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "gitlab.com/vdat/mcsvc/chat/docs"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/dchat"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/groups"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/userdetail"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/utils"
	"log"
	"net"
	"net/http"
	"os"
	_ "path"
	"path/filepath"
	"time"
)

// spaHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// @title Swagger Chat server API
// @version 0.1
// @description This is swagger for chat server.
// @description local:	  http://localhost:5000/.
// @description staging:    http://vdat-mcsvc-chat-staging.vdatlab.com/.
// @description production: https://vdat-mcsvc-chat.vdatlab.com/.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @query.collection.format multi
// @Schemes http https
// @host localhost:5000
// @BasePath /api/v1
// @query.collection.format multi
func main() {
	//go metrics()
	database.Connect()

	//readfile
	utils.CheckFileSocketId()

	//start broker
	go dchat.Wsbroker.Run()

	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", CheckHelthHandlr).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/chat/{idgroup}", auth.AuthenMiddleJWT(dchat.ChatHandlr)).Methods(http.MethodGet, http.MethodOptions)
	// handler
	r.HandleFunc("/message/{socketId}", dchat.ChatHandlr).Methods(http.MethodGet, http.MethodOptions)
	//api
	groups.RegisterGroupApi(r)
	userdetail.RegisterUserApi(r)

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// Handler web app
	spa := spaHandler{staticPath: "public", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	r.Use(mux.CORSMethodMiddleware(r))

	srv := &http.Server{
		Handler: r,
		Addr:    ":5000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("server starting")
	log.Fatal(srv.ListenAndServe())

}

func metrics() {
	// The debug listener mounts the http.DefaultServeMux, and serves up
	// stuff like the Prometheus metrics route, the Go debug and profiling
	// routes, and so on.
	debugListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("transport", "debug/HTTP", "during", "Listen", "err", err)
		os.Exit(1)
	}

	defer debugListener.Close()

	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	fmt.Println(http.Serve(debugListener, http.DefaultServeMux))
}
