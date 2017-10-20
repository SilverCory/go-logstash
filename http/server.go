package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"github.com/SilverCory/go-logstash/log"
)

type Server struct {
	authKey    string
	rootRouter *mux.Router
	logRoute   *mux.Route
	logger     *log.Logger
	LogCallback func(path, data string)
}

func New(authKey string, logger *log.Logger) (s *Server) {

	s = &Server{
		rootRouter: mux.NewRouter(),
		authKey:    authKey,
		logger:     logger,
	}

	s.logRoute = s.rootRouter.HandleFunc("/log/{path:.*}", s.HandleLog).Methods("POST")
	return

}

func (s *Server) HandleLog(writer http.ResponseWriter, request *http.Request) {

	if request.Header.Get("auth") != s.authKey {
		writer.WriteHeader(http.StatusForbidden)
		return
	}

	vars := mux.Vars(request)
	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprint(err)))
		return
	}

	err = s.logger.Log(vars["path"], bytes)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprint(err)))
		return
	}

	if s.LogCallback != nil {
		go s.LogCallback(vars["path"], string(bytes))
	}

}

func (s *Server) Open() {
	http.Handle("/", s.rootRouter)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println("Fatal err:", err)
	}
}
