package models

type Song struct {
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
}

type FilmSQL struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	NameStub     string `json:"name_stub"`
	AirDateStart int64   `json:"air_date_start"`
	AirDateEnd   int64 `json:"air_date_end"`
	Image        string `json:"image"`
	ReleaseDate  int64 `json:"release_date"`
	TypeFilm     string `json:"type_film"`
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
