package httpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gevorg-tsat/link-shortener/config"
	"github.com/gevorg-tsat/link-shortener/internal/errors"
	"github.com/gevorg-tsat/link-shortener/internal/server"
	pb "github.com/gevorg-tsat/link-shortener/internal/shortener_v1"
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPServer struct {
	S *http.Server
}

func New(grpcServer *server.ShortenerServer, cfg *config.Config) *HTTPServer {
	return &HTTPServer{S: &http.Server{
		Addr:    fmt.Sprintf("%v:%v", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler: handlers(grpcServer, cfg),
	}}
}

func handlers(shortenerServer *server.ShortenerServer, cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/{identifier}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			params := mux.Vars(r)
			identifier := params["identifier"]
			shortUrl := fmt.Sprintf("http://%v:%v/%v", cfg.HTTP.Host, cfg.HTTP.Port, identifier)
			originalLink, err := shortenerServer.Get(context.Background(), &pb.ShortLink{Url: shortUrl})
			if err != nil {
				errors.WriteResponse(w, err)
				return
			}
			http.Redirect(w, r, originalLink.Url, http.StatusSeeOther)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			query := r.URL.Query()
			if !query.Has("url") {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("url query required"))
				return
			}
			shortLink, err := shortenerServer.Post(context.Background(), &pb.OriginalLink{Url: query["url"][0]})
			if err != nil {
				errors.WriteResponse(w, err)
				return
			}
			shortLink.Url = fmt.Sprintf("http://%v:%v/%v", cfg.HTTP.Host, cfg.HTTP.Port, shortLink.Url)
			b, err := json.Marshal(shortLink)
			if err != nil {
				errors.WriteResponse(w, errors.InternalServerError)
				return
			}
			w.Write(b)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}).Methods(http.MethodPost)

	return router
}
