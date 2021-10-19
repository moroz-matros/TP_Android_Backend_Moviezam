package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/moroz-matros/TP_Android_Backend_Moviezam/application/models"
	"github.com/moroz-matros/TP_Android_Backend_Moviezam/application/usecase"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	Uc *usecase.Usecase
}

func CreateHandler(m *mux.Router, uc *usecase.Usecase) *Handler {
	handler := &Handler{Uc: uc}

	m.HandleFunc("/find_film", handler.SearchFilm)
	// /find_film?search=Harry+Potter

	m.HandleFunc("/find_song", handler.SearchSong)
	// /songs?search=just+tonight

	m.HandleFunc("/find_artist", handler.SearchArtist)
	// /find_artist?search=sia

	m.HandleFunc("/artist", handler.GetArtistBySongId)
	// /artist?song_id=1
	m.HandleFunc("/songs", handler.GetSongs)
	// /songs?artist_id=1&film_id=1

	return handler
}


func (h *Handler) GetFilmsBySongId(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) GetSimilarFilmsBySongName(w http.ResponseWriter, r *http.Request){

}

func (h *Handler) GetArtistBySongId(w http.ResponseWriter, r *http.Request){
	songId := r.URL.Query().Get("song_id")

}

func (h *Handler) GetSongs(w http.ResponseWriter, r *http.Request){
	artistId, err := strconv.Atoi(r.URL.Query().Get("artist_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if artistId != 0 {
		songs, err := h.Uc.GetSongsByArtistId(artistId)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		body, _ := json.Marshal(songs)
		w.WriteHeader(http.StatusOK)
		w.Write(body)

		return
	}

	filmId, err := strconv.Atoi(r.URL.Query().Get("film_id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if filmId != 0 {
		songs, err := h.Uc.GetSongsByFilmId(artistId)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		body, _ := json.Marshal(songs)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		return
	}
}

func (h *Handler) SearchFilm(w http.ResponseWriter, r *http.Request){
	searchString := r.URL.Query().Get("search")

	films, err := h.Uc.SearchFilm(searchString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(films)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) SearchSong(w http.ResponseWriter, r *http.Request){
	searchString := r.URL.Query().Get("search")

	songs, err := h.Uc.SearchSong(searchString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(songs)

	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func (h *Handler) SearchArtist(w http.ResponseWriter, r *http.Request){
	searchString := r.URL.Query().Get("search")

	artists, err := h.Uc.SearchArtist(searchString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(artists)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) GetSongByName(w http.ResponseWriter, r *http.Request){
	searchString := r.URL.Query().Get("name")

	song := models.Song{Id: 11, Name: searchString}

	body, err := json.Marshal(song)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) GetSongsByFilm(w http.ResponseWriter, r *http.Request){
	filmId := r.URL.Query().Get("film_id")

}
