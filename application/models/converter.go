package models

import "time"

func ConvertFilm(film FilmSQL) FilmJSON {
	return FilmJSON{
		Id:           film.Id,
		Name:         film.Name,
		NameStub:     film.NameStub,
		AirDateStart: time.Unix(film.AirDateStart, 0).String(),
		AirDateEnd:   time.Unix(film.AirDateEnd, 0).String(),
		Image:        film.Image,
		ReleaseDate:  film.ReleaseDate.String(),
	}

}

func ConvertSong(song SongSQL) SongJSON {
	return SongJSON{
		Id:           song.Id,
		Name:         song.Name,
		NameStub:     song.NameStub,
		AlbumName:      song.AlbumName,
		ExternalArtUrl: song.ExternalArtUrl.String,
		Amazon:         song.Amazon.String,
		AppleMusic:     song.AppleMusic.String,
		Itunes:         song.Itunes.String,
		Spotify:        song.Spotify.String,
		Youtube:        song.Youtube.String,
	}

}