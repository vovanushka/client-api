package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vovanushka/client-api/service"
)

type Server struct {
	router *mux.Router
}

func NewServer(p *service.PortService) *Server {
	s := Server{router: mux.NewRouter()}
	NewPortRouter(p, s.newSubrouter("/port"))
	return &s
}

func (s *Server) Start(port string) {
	log.Println("Listening on port " + port)

	if err := http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(s.router)); err != nil {
		log.Fatal("http.ListenAndServe: ", err)
	}
}

func (s *Server) newSubrouter(path string) *mux.Router {
	return s.router.PathPrefix(path).Subrouter()
}
