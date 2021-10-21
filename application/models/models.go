package models

import (
	"database/sql"
	"time"
)

type SongJSON struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	NameStub       string `json:"name_stub"`
	AlbumName      string `json:"album_name"`
	ExternalArtUrl string `json:"external_art_url"`
	Amazon         string `json:"amazon"`
	AppleMusic     string `json:"apple_music"`
	Itunes         string `json:"itunes"`
	Spotify        string `json:"spotify"`
	Youtube        string `json:"youtube"`
}

type SongSQL struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	NameStub       string         `json:"name_stub"`
	AlbumName      string         `json:"album_name"`
	ExternalArtUrl sql.NullString `json:"external_art_url"`
	Amazon         sql.NullString `json:"amazon"`
	AppleMusic     sql.NullString `json:"apple_music"`
	Itunes         sql.NullString `json:"itunes"`
	Spotify        sql.NullString `json:"spotify"`
	Youtube        sql.NullString `json:"youtube"`
}

type Artist struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	NameStub      string `json:"name_stub"`
	Image         string `json:"image"`
	Biography     string `json:"biography"`
	UrlOfficial   string `json:"url_official"`
	UrlWikipedia  string `json:"url_wikipedia"`
	UrlYoutube    string `json:"url_youtube"`
	UrlSpotify    string `json:"url_spotify"`
	UrlApplemusic string `json:"url_applemusic"`
	UrlItunes     string `json:"url_itunes"`
	UrlAmazon     string `json:"url_amazon"`
	Songs []SongCard `json:"songs"`
}

type ArtistSQL struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	NameStub      string `json:"name_stub"`
	Image         string `json:"image"`
	Biography     sql.NullString `json:"biography"`
	UrlOfficial   sql.NullString `json:"url_official"`
	UrlWikipedia  sql.NullString `json:"url_wikipedia"`
	UrlYoutube    sql.NullString `json:"url_youtube"`
	UrlSpotify    sql.NullString `json:"url_spotify"`
	UrlApplemusic sql.NullString `json:"url_applemusic"`
	UrlItunes     sql.NullString `json:"url_itunes"`
	UrlAmazon     sql.NullString `json:"url_amazon"`
}

type SongCard struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Artist string `json:"artist"`
	Album  string `json:"album"`
	ExternalArtUrl string `json:"external_art_url"`
}

type FilmSQL struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	NameStub     string    `json:"name_stub"`
	AirDateStart int64     `json:"air_date_start"`
	AirDateEnd   int64     `json:"air_date_end"`
	Image        string    `json:"image"`
	ReleaseDate  time.Time `json:"release_date"`
	TypeFilm     string    `json:"type_film"`
}

type FilmJSON struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	NameStub     string `json:"name_stub"`
	AirDateStart string `json:"air_date_start"`
	AirDateEnd   string `json:"air_date_end"`
	Image        string `json:"image"`
	ReleaseDate  string `json:"release_date"`
}

type FilmCard struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	Image       string `json:"image"`
}

type ArtistCard struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
}

type SongShazam struct {
	Id             int        `json:"id"`
	Name           string     `json:"name"`
	NameStub       string     `json:"name_stub"`
	Artist         string     `json:"artist"`
	AlbumName      string     `json:"album_name"`
	ExternalArtUrl string     `json:"external_art_url"`
	Amazon         string     `json:"amazon"`
	AppleMusic     string     `json:"apple_music"`
	Itunes         string     `json:"itunes"`
	Spotify        string     `json:"spotify"`
	Youtube        string     `json:"youtube"`
	Films          []FilmCard `json:"films"`
}
