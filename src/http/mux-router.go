package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type muxRouter struct {
}

var (
	muxDispatcher = mux.NewRouter()
)


func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri,f).Methods("GET")
}

func (m *muxRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri,f).Methods("POST")
}

func (m *muxRouter) SERVE(port string) {
	fmt.Printf("Mux Run ON PORT %v", port)
	http.ListenAndServe(port,muxDispatcher)
}
