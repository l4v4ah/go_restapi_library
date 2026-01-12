package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	htppHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) HTTPServer {
	return HTTPServer{
		htppHandlers: httpHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/library").Methods("POST").HandlerFunc(s.htppHandlers.HandleCreateNewBook)
	router.Path("/library/{book}").Methods("PATCH").HandlerFunc(s.htppHandlers.HandleCompleteBook)
	router.Path("/library/{book}").Methods("GET").HandlerFunc(s.htppHandlers.HandleGetBook)
	router.Path("/library").Methods("GET").Queries("author", "").HandlerFunc(s.htppHandlers.HandleGetAuthorBooks)
	router.Path("/library").Methods("GET").Queries("complete", "").HandlerFunc(s.htppHandlers.HandleGetCompletedBooks)
	router.Path("/library").Methods("GET").HandlerFunc(s.htppHandlers.HandleGetAllBooks)
	router.Path("/library/{book}").Methods("DELETE").HandlerFunc(s.htppHandlers.HandleDeleteBook)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
