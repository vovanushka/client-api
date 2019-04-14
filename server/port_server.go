package server

import (
	"net/http"

	"github.com/vovanushka/client-api/reader"

	"google.golang.org/grpc/codes"

	"github.com/gorilla/mux"
	"github.com/vovanushka/client-api/service"
	"google.golang.org/grpc/status"
)

type portRouter struct {
	portService *service.PortService
}

func NewPortRouter(p *service.PortService, router *mux.Router) *mux.Router {
	portRouter := portRouter{p}

	router.HandleFunc("/{portID}", portRouter.getPort).Methods("GET")
	router.HandleFunc("/upload", portRouter.uploadPorts).Methods("PUT")

	return router
}

func (p *portRouter) getPort(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	portID := vars["portID"]
	if portID == "" {
		Error(w, http.StatusBadRequest, "bad port id")
		return
	}

	port, err := p.portService.Get(portID)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			Error(w, http.StatusInternalServerError, "server error")
		} else {
			if st.Code() == codes.NotFound {
				Error(w, http.StatusNotFound, "port was not found")
			} else {
				Error(w, http.StatusInternalServerError, st.Message())
			}
		}
		return
	}

	Json(w, http.StatusOK, port)
}

func (p *portRouter) uploadPorts(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		Error(w, http.StatusBadRequest, "file undefined")
		return
	}
	defer file.Close()
	reader, err := reader.NewJSONReader(file, p.portService)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	err = reader.ReadStream()
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	Json(w, http.StatusOK, nil)
}
