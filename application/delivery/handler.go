package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
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
	// /find_film?search=Harry+Potter&page=1

	m.HandleFunc("/film", handler.GetFilmById)
	// /film?id=1

	m.HandleFunc("/similar_films", handler.GetSimilarFilms)
	// /similar_films?film_id=1&page=1

	m.HandleFunc("/find_song", handler.SearchSong)
	// /songs?search=just+tonight&page=1

	m.HandleFunc("/song", handler.GetSongById)
	// /song?id=1

	m.HandleFunc("/find_artist", handler.SearchArtist)
	// /find_artist?search=sia

	m.HandleFunc("/artist", handler.GetArtistById)
	// /artist?id=1

	m.HandleFunc("/moviezam", handler.ShazamSong)
	// /moviezam?search=lala

	return handler
}

func (h *Handler) GetSimilarFilms(w http.ResponseWriter, r *http.Request) {
	filmId := r.URL.Query().Get("film_id")
	id, err := strconv.Atoi(filmId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	page := r.URL.Query().Get("page")
	var p int
	if page == "" {
		p = 1
	} else {
		p, err = strconv.Atoi(page)
		if err != nil || p < 0 {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if p == 0 {
			p = 1
		}
	}

	films, err := h.Uc.GetSimilarFilms(id, p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(films)
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) SearchFilm(w http.ResponseWriter, r *http.Request) {
	searchString := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	var p int
	var err error
	if page == "" {
		p = 1
	} else {
		p, err = strconv.Atoi(page)
		if err != nil || p < 0 {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if p == 0 {
			p = 1
		}
	}

	films, err := h.Uc.SearchFilm(searchString, p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(films)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) SearchSong(w http.ResponseWriter, r *http.Request) {
	searchString := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	var p int
	var err error
	if page == "" {
		p = 1
	} else {
		p, err = strconv.Atoi(page)

		if err != nil || p < 0 {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if p == 0 {
			p = 1
		}
	}
	songs, err := h.Uc.SearchSong(searchString, p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(songs)

	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func (h *Handler) SearchArtist(w http.ResponseWriter, r *http.Request) {
	searchString := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	var p int
	var err error
	if page == "" {
		p = 1
	} else {
		p, err = strconv.Atoi(page)

		if err != nil || p < 0 {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if p == 0 {
			p = 1
		}
	}

	artists, err := h.Uc.SearchArtist(searchString, p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(artists)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) GetSongsByFilm(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("film_id")

}

func (h *Handler) ShazamSong(w http.ResponseWriter, r *http.Request) {
	searchString := r.URL.Query().Get("search")
	song, err := h.Uc.ShazamSong(searchString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(song)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) GetArtistById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	artist, err := h.Uc.GetArtistById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(artist)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) GetFilmById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	film, err := h.Uc.GetFilmById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(film)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handler) GetSongById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	song, err := h.Uc.GetSongById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, _ := json.Marshal(song)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
