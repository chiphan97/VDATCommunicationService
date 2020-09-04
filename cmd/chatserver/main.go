package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/vdat/mcsvc/chat/pkg/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/groups"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/useronline"
	"net"
	"net/http"
	"os"
)

// https://github.com/gorilla/websocket/blob/master/examples/chat/main.go
func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func main() {
	go metrics()
	r := mux.NewRouter()
	fmt.Println("starting")

	database.Connect()

	r.HandleFunc("/", serveHome)
	//http.HandleFunc("/test", service.TestHandler)
	//http.HandleFunc("/test", handler.TestHandler)
	//r.HandleFunc("/user-online", handler.UserOnlineHandler)
	//http.Handle("/", http.FileServer(http.Dir(".")))

	//api
	groups.RegisterGroupApi(r)
	useronline.RegisterUserOnlineApi(r)

	r.Use(mux.CORSMethodMiddleware(r))
	//corsObj := handlers.AllowedOrigins([]string{"*"})
	err := http.ListenAndServe(":5000", r)
	if err != nil {
		panic("Error: " + err.Error())
	}
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
