package controller

import (
	"crudMusic/usecases"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ArtistHandler struct {
	artistUsecase usecases.ArtistUsecase
}

func ArtistController(r *mux.Router, service usecases.ArtistUsecase) {
	artistHandler := ArtistHandler{service}

	r.HandleFunc("/artists", artistHandler.GetArtists).Methods(http.MethodGet)
}

func (s *ArtistHandler) GetArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := s.artistUsecase.GetArtists()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfArtists, err := json.Marshal(artists)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfArtists))
	w.Write([]byte("Data successfully found"))
}
