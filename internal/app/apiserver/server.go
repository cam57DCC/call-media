package apiserver

import (
	"fmt"
	"github.com/cam57DCC/call-media/internal/app/controller"
	"net/http"
)

type Server struct {
	Router *http.ServeMux
}

func newServer() *Server {
	s := &Server{
		Router: http.NewServeMux(),
	}

	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	s.Router.ServeHTTP(responseWriter, request)
}

func (s *Server) configureRouter() {

	s.Router.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})
	s.Router.HandleFunc("/request/add", controller.AddRequest)
	s.Router.HandleFunc("/request/repeat", controller.RequestRepeat)
	s.Router.HandleFunc("/request/count", controller.GetCountRequests)
	fmt.Println("Server run!")

}
