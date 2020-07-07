package controller

import (
	"crudMusic/usecases"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SongHandler struct {
	songUsecase usecases.SongUsecase
}

func SongController(r *mux.Router, service usecases.SongUsecase) {
	songHandler := SongHandler{service}

	r.HandleFunc("/songs", songHandler.GetSongs).Methods(http.MethodGet)
	r.HandleFunc("/song/genre/{genre}", songHandler.GetSongByGenre).Methods(http.MethodGet)
	r.HandleFunc("/song/artist/{artist}", songHandler.GetSongByArtist).Methods(http.MethodGet)
	r.HandleFunc("/song/album/{album}", songHandler.GetSongByAlbum).Methods(http.MethodGet)
	r.HandleFunc("/song/release_year/{year}", songHandler.GetSongByReleaseYear).Methods(http.MethodGet)
}

func (s *SongHandler) GetSongs(w http.ResponseWriter, r *http.Request) {
	songs, err := s.songUsecase.GetSongs()
	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	byteOfsongs, err := json.Marshal(songs)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfsongs))
	w.Write([]byte("Data successfully found"))
}

func (s *SongHandler) GetSongByGenre(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strGenre := vars["genre"]

	songs, err := s.songUsecase.GetSongByGenre(strGenre)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfSongs, err := json.Marshal(songs)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfSongs))
}

func (s *SongHandler) GetSongByArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strArtist := vars["artist"]

	songs, err := s.songUsecase.GetSongByArtist(strArtist)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfSongs, err := json.Marshal(songs)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfSongs))
}

func (s *SongHandler) GetSongByAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strAlbum := vars["album"]

	songs, err := s.songUsecase.GetSongByAlbum(strAlbum)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfSongs, err := json.Marshal(songs)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfSongs))
}

func (s *SongHandler) GetSongByReleaseYear(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strYear := vars["year"]

	year, _ := strconv.Atoi(strYear)

	songs, err := s.songUsecase.GetSongByReleaseYear(year)
	if err != nil {
		w.Write([]byte("Data Not Found!"))
	}

	byteOfSongs, err := json.Marshal(songs)

	if err != nil {
		w.Write([]byte("Data not found!"))
	}

	w.Write([]byte(byteOfSongs))
}
