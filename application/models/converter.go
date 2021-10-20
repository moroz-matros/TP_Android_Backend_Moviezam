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