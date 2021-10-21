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

func ConvertFilmToCard(film FilmSQL) FilmCard {
	return FilmCard{
		Id:           film.Id,
		Name:         film.Name,
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

func ConvertSongToCard(song SongSQL, artist string, album string) SongCard {
	return SongCard{
		Id:     song.Id,
		Name:   song.Name,
		Artist: artist,
		Album: album,
		ExternalArtUrl: song.ExternalArtUrl.String,
	}
}

func ConvertSongToShazam(song SongSQL, artist string, album string, films []FilmCard) SongShazam {
	return SongShazam{
		Id:             song.Id,
		Name:           song.Name,
		NameStub:       song.NameStub,
		Artist:         artist,
		AlbumName:      album,
		ExternalArtUrl: song.ExternalArtUrl.String,
		Amazon:         song.Amazon.String,
		AppleMusic:     song.AppleMusic.String,
		Itunes:         song.Itunes.String,
		Spotify:        song.Spotify.String,
		Youtube:        song.Youtube.String,
		Films:          films,
	}
}

func ConvertArtist(artist ArtistSQL) Artist {
	return Artist{
		Id:            artist.Id,
		Name:          artist.Name,
		NameStub:      artist.Name,
		Image:         artist.Image,
		Biography:     artist.Biography.String,
		UrlOfficial:   artist.UrlOfficial.String,
		UrlWikipedia:  artist.UrlWikipedia.String,
		UrlYoutube:    artist.UrlYoutube.String,
		UrlSpotify:    artist.UrlSpotify.String,
		UrlApplemusic: artist.UrlApplemusic.String,
		UrlItunes:     artist.UrlItunes.String,
		UrlAmazon:     artist.UrlAmazon.String,
	}
}

func ConvertArtistToCard(artist ArtistSQL) ArtistCard {
	return ArtistCard{
		Id:            artist.Id,
		Name:          artist.Name,
		Image:         artist.Image,
	}
}