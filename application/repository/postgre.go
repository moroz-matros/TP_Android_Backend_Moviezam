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


func (r *Repo) GetSongsByFilmId(id int) ([]models.SongSQL, error){
	var songs []models.SongSQL
	err := pgxscan.Select(context.Background(), r.Pool, &songs,
		"SELECT * from sounds WHERE id IN " +
		"(SELECT id_sound FROM film_sound WHERE id_film = $1)", id)

	if errors.As(err, &sql.ErrNoRows) {
		return []models.SongSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.SongSQL{}, err
	}

	return songs, nil
}

func (r *Repo) GetArtistNameBySongId(id int) (string, error){
	var name string
	err := r.Pool.QueryRow(context.Background(),
		"SELECT a.name from artists as a, sound_artist as sa " +
		"WHERE sa.id_sound = $1 AND sa.id_artist = a.id", id).Scan(
			&name)

	if errors.As(err, &sql.ErrNoRows){
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return name, nil
}

func (r *Repo) GetAlbumNameBySongId(id int) (string, error){
	var name string
	err := r.Pool.QueryRow(context.Background(),
		"SELECT a.name from albums as a, album_sound as sa " +
			"WHERE sa.id_sound = $1 AND sa.id_album = a.id", id).Scan(&name)

	if errors.As(err, &sql.ErrNoRows){
		return "", nil
	}

	if err != nil {
		return "",  err
	}

	return name, nil
}

func (r *Repo) GetSongsByArtistId(id int) ([]models.SongSQL, error){
	var songs []models.SongSQL
	err := pgxscan.Select(context.Background(), r.Pool, &songs,
		"SELECT * FROM sounds WHERE id IN " +
		"(SELECT DISTINCT id_sound FROM sound_artist WHERE id_artist = $1)", id)

	if errors.As(err, &sql.ErrNoRows) {
		return []models.SongSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.SongSQL{}, err
	}

	return songs, nil
}

func (r *Repo) SearchArtist(name string, page int) ([]models.ArtistSQL, error) {
	var artists []models.ArtistSQL
	err := pgxscan.Select(context.Background(), r.Pool, &artists,
		"SELECT * FROM artists WHERE LOWER(name) " +
		"LIKE '%' || $1 || '%' ORDER BY id LIMIT 20 OFFSET $2", strings.ToLower(name), (page-1)*10)

	if errors.As(err, &sql.ErrNoRows) {
		return []models.ArtistSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.ArtistSQL{}, err
	}

	return artists, nil
}

func (r *Repo) SearchFilm(name string, page int) ([]models.FilmSQL, error) {
	var films []models.FilmSQL
	err := pgxscan.Select(context.Background(), r.Pool, &films,
		"SELECT * FROM films WHERE LOWER(name) LIKE '%' || $1 || '%'" +
		" ORDER BY id LIMIT 20 OFFSET $2", strings.ToLower(name), (page-1)*10)

	if errors.As(err, &sql.ErrNoRows) {
		return []models.FilmSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.FilmSQL{}, err
	}

	return films, nil
}

func (r *Repo) SearchSong(name string, page int) ([]models.SongSQL, error) {
	var songs []models.SongSQL
	err := pgxscan.Select(context.Background(), r.Pool, &songs,
		"SELECT * FROM sounds WHERE LOWER(name) LIKE '%' || $1 || '%' " +
			" ORDER BY id LIMIT 20 OFFSET $2", strings.ToLower(name), (page-1)*10)
	if errors.As(err, &sql.ErrNoRows) {
		return []models.SongSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.SongSQL{}, err
	}

	return songs, nil
}

func (r *Repo) GetSimilarFilms(id int, page int) ([]models.FilmSQL, error) {
	var films []models.FilmSQL
	err := pgxscan.Select(context.Background(), r.Pool, &films,
		"SELECT * FROM films WHERE id IN" +
		"(SELECT DISTINCT (film_id) FROM similars WHERE similar_film_id = $1)" +
		" ORDER BY id LIMIT 20 OFFSET $2", id, (page-1)*10)

	if errors.As(err, &sql.ErrNoRows) {
		return []models.FilmSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.FilmSQL{}, err
	}

	return films, nil
}

func (r *Repo) GetFilmsBySongId(id int) ([]models.FilmSQL, error) {
	var films []models.FilmSQL
	err := pgxscan.Select(context.Background(), r.Pool, &films,
		"SELECT * FROM films WHERE id IN (SELECT DISTINCT id_film from film_sound " +
		"WHERE id_sound = $1)", id)
	if errors.As(err, &sql.ErrNoRows) {
		return []models.FilmSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return []models.FilmSQL{}, err
	}


	return films, nil
}

func (r *Repo) GetArtistById(id int) (models.ArtistSQL, error) {
	var artists []models.ArtistSQL
	err := pgxscan.Select(context.Background(), r.Pool, &artists,
		"SELECT * FROM artists WHERE id = $1", id)

	if errors.As(err, &sql.ErrNoRows) {
		return models.ArtistSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return models.ArtistSQL{}, err
	}

	return artists[0], nil
}

func (r *Repo) GetFilmById(id int) (models.FilmSQL, error) {
	var films []models.FilmSQL
	err := pgxscan.Select(context.Background(), r.Pool, &films,
		"SELECT * FROM films WHERE id = $1", id)
	if errors.As(err, &sql.ErrNoRows) {
		return models.FilmSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return models.FilmSQL{}, err
	}


	return films[0], nil
}

func (r *Repo) GetArtistBySongId(id int) (models.ArtistSQL, error) {
	var artists []models.ArtistSQL
	err := pgxscan.Select(context.Background(), r.Pool, &artists,
		"SELECT * FROM artists WHERE id IN " +
		"(SELECT DISTINCT id_artist FROM sound_artist WHERE id_sound = $1)", id)

	if errors.As(err, &sql.ErrNoRows) {
		return models.ArtistSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return models.ArtistSQL{}, err
	}

	return artists[0], nil
}

func (r *Repo) GetSimilarFilmsNoPage(id int) ([]models.FilmSQL, error) {
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

func (r *Repo) GetSongById(id int) (models.SongSQL, error){
	var songs []models.SongSQL
	err := pgxscan.Select(context.Background(), r.Pool, &songs,
		"SELECT * FROM sounds WHERE id = $1", id)

	if errors.As(err, &sql.ErrNoRows) {
		return models.SongSQL{}, nil
	}

	if err != nil {
		log.Println(err)
		return models.SongSQL{}, err
	}

	return songs[0], nil
}