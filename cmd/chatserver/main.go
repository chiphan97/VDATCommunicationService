package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/auth"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/database"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/dchat"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/groups"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/userdetail"
	"gitlab.com/vdat/mcsvc/chat/pkg/service/utils"
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
	//go metrics()

	//start broker
	go dchat.Wsbroker.Run()

	database.Connect()

	//readfile
	utils.CheckFileSocketId()

	r := mux.NewRouter()

	// handler
	//r.HandleFunc("/user-online", handler.UserOnlineHandler)
	r.HandleFunc("/", serveHome)
	//r.HandleFunc("/chat/{idgroup}",auth.AuthenMiddleJWT(dchat.ChatHandlr))
	r.HandleFunc("/chat/{idgroup}", auth.AuthenMiddleJWT(dchat.ChatHandlr)).Methods(http.MethodGet, http.MethodOptions)
	//api
	groups.RegisterGroupApi(r)
	userdetail.RegisterUserApi(r)

	r.Use(mux.CORSMethodMiddleware(r))

	fmt.Println("starting")
	//fmt.Println(len(utils.ArraySocketId))
	//write file
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//go func() {
	//	for sig := range c {
	//		log.Printf("captured %v, stopping profiler and exiting..", sig)
	//		fmt.Println(utils.ArraySocketId)
	//		err := utils.WriteLines(utils.ArraySocketId,"socketid.data")
	//		if err !=nil{
	//			log.Fatal(err)
	//		}
	//		pprof.StopCPUProfile()
	//		os.Exit(1)
	//	}
	//}()

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
