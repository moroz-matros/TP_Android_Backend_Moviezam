package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/pgxpool"
	"github.com/moroz-matros/TP_Android_Backend_Moviezam/application/models"
	"log"
	"strings"
)

type Repo struct {
	Pool *pgxpool.Pool
}

func CreateRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{Pool: pool}
}


func (r *Repo) GetSongByName(name string){

}

func (r *Repo) GetSongsByFilmId(id int) ([]models.SongSQL, error){
	var songs []models.SongSQL
	err := pgxscan.Select(context.Background(), r.Pool, &songs,
		"SELECT * from sounds " +
		"JOIN film_sound ON id_film = $1", id)

	if errors.As(err, &sql.ErrNoRows) {
		return []models.SongSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.SongSQL{}, err
	}

	return songs, nil
}

func (r *Repo) GetArtistBySongId(id uint) (models.Artist, error){
	var artist models.Artist
	err := r.Pool.QueryRow(context.Background(),
		"SELECT * from artists " +
		"JOIN sound_artist on id_sound = $1", id).Scan(
			&artist.Id, &artist.Image, &artist.Name, &artist.Biography,
			&artist.NameStub, &artist.UrlAmazon, &artist.UrlApplemusic,
			&artist.UrlItunes, &artist.UrlOfficial, &artist.UrlSpotify, &artist.UrlWikipedia,
			&artist.UrlYoutube)

	if errors.As(err, &sql.ErrNoRows){
		return models.Artist{}, nil
	}

	if err != nil {
		return models.Artist{}, err
	}

	return artist, nil
}

func (r *Repo) GetSongsByArtistId(id int) ([]models.SongSQL, error){
	var songs []models.SongSQL
	err := pgxscan.Select(context.Background(), r.Pool, &songs,
		"SELECT * FROM songs JOIN sound_artist ON artist_id = $1", id)

	if errors.As(err, &sql.ErrNoRows) {
		return []models.SongSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.SongSQL{}, err
	}

	return songs, nil
}

func (r *Repo) SearchArtist(name string) ([]models.Artist, error) {
	var artists []models.Artist
	err := pgxscan.Select(context.Background(), r.Pool, &artists,
		"SELECT * FROM artists WHERE LOWER(name) LIKE '%' || $1 || '%'", strings.ToLower(name))

	if errors.As(err, &sql.ErrNoRows) {
		return []models.Artist{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.Artist{}, err
	}

	return artists, nil
}

func (r *Repo) SearchFilm(name string) ([]models.FilmSQL, error) {
	var films []models.FilmSQL
	err := pgxscan.Select(context.Background(), r.Pool, &films,
		"SELECT * FROM films WHERE LOWER(name) LIKE '%' || $1 || '%'", strings.ToLower(name))

	if errors.As(err, &sql.ErrNoRows) {
		return []models.FilmSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.FilmSQL{}, err
	}

	return films, nil
}

func (r *Repo) SearchSong(name string) ([]models.SongSQL, error) {
	var songs []models.SongSQL
	err := pgxscan.Select(context.Background(), r.Pool, &songs,
		"SELECT * FROM sounds WHERE LOWER(name) LIKE '%' || $1 || '%'", strings.ToLower(name))
	if errors.As(err, &sql.ErrNoRows) {
		return []models.SongSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.SongSQL{}, err
	}

	return songs, nil
}

func (r *Repo) GetSimilarFilms(id int) ([]models.FilmSQL, error) {
	var films []models.FilmSQL
	err := pgxscan.Select(context.Background(), r.Pool, &films,
		"SELECT * FROM films WHERE id IN" +
		"(SELECT DISTINCT (film_id) FROM similars WHERE similar_film_id = $1)", id)

	if errors.As(err, &sql.ErrNoRows) {
		return []models.FilmSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.FilmSQL{}, err
	}

	return films, nil
}