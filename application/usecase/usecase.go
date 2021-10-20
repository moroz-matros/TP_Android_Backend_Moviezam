package usecase

import (
	"github.com/moroz-matros/TP_Android_Backend_Moviezam/application/models"
	"github.com/moroz-matros/TP_Android_Backend_Moviezam/application/repository"
)

type Usecase struct {
	Repo *repository.Repo
}

func CreateUsecase(repo *repository.Repo) *Usecase {
	return &Usecase{Repo: repo}
}

func (uc *Usecase) SearchSong(name string) ([]models.Song, error) {
	return uc.Repo.SearchSong(name)
}

func (uc *Usecase) SearchArtist(name string) ([]models.Artist, error) {
	return  uc.Repo.SearchArtist(name)
}

func (uc *Usecase) GetSongsByArtistId(id int) ([]models.Song, error){
	return uc.Repo.GetSongsByArtistId(id)
}

func (uc *Usecase) GetSongsByFilmId(id int) ([]models.Song, error){
	return uc.Repo.GetSongsByFilmId(id)
}

func (uc *Usecase) SearchFilm(name string) ([]models.FilmJSON, error) {
	filmsSQL, err := uc.Repo.SearchFilm(name)
	if err != nil {
		return []models.FilmJSON{}, err
	}

	var filmsJSON []models.FilmJSON
	for _, elem := range filmsSQL {
		filmsJSON = append(filmsJSON, models.ConvertFilm(elem))
	}
	return filmsJSON, nil
}

func (uc *Usecase) GetSimilarFilms(id int) ([]models.FilmJSON, error) {
	filmsSQL, err := uc.Repo.GetSimilarFilms(id)
	if err != nil {
		return []models.FilmJSON{}, err
	}

	var filmsJSON []models.FilmJSON
	for _, elem := range filmsSQL {
		filmsJSON = append(filmsJSON, models.ConvertFilm(elem))
	}
	return filmsJSON, nil
}