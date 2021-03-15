package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	js "github.com/neel229/linktree-clone/internal/serializer"
	"github.com/neel229/linktree-clone/internal/shortener"
)

type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

type handler struct {
	redirectService shortener.RedirectService
}

func NewHandler(redirectService shortener.RedirectService) RedirectHandler {
	return &handler{
		redirectService: redirectService,
	}
}

func (h *handler) serializer(contentType string) shortener.RedirectSerialzer {
	return &js.Redirect{}
}

func (h *handler) Get(rw http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	redirect, err := h.redirectService.Find(code)
	if err != nil {
		log.Fatal(err)
		http.Error(rw, "error fetching redirect data", http.StatusBadRequest)
		return
	}
	json.NewEncoder(rw).Encode(&redirect.Handles)
}

func (h *handler) Post(rw http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		http.Error(rw, "error reading data from the URL", http.StatusInternalServerError)
		return
	}

	redirect, err := h.serializer(contentType).Decode(data)
	if err != nil {
		log.Fatal(err)
		http.Error(rw, "error serializing data", http.StatusInternalServerError)
	}

	err = h.redirectService.Store(redirect)
	if err != nil {
		log.Fatal(err)
		http.Error(rw, "error storing data", http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(&data)
}
