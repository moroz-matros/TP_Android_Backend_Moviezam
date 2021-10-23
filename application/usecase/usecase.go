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

func (uc *Usecase) SearchSong(name string, page int) ([]models.SongCard, error) {
	songs, err := uc.Repo.SearchSong(name, page)
	if err != nil {
		return []models.SongCard{}, err
	}

	var songCards []models.SongCard
	for _, elem := range songs{
		artistName, err := uc.Repo.GetArtistNameBySongId(elem.Id)
		if err != nil {
			return []models.SongCard{}, err
		}
		albumName, err := uc.Repo.GetAlbumNameBySongId(elem.Id)
		if err != nil {
			return []models.SongCard{}, err
		}
		songCards = append(songCards, models.ConvertSongToCard(elem, artistName,
			albumName))
	}

	return songCards, nil
}

func (uc *Usecase) SearchArtist(name string, page int) ([]models.ArtistCard, error) {
	artistsSQL, err := uc.Repo.SearchArtist(name, page)
	if err != nil {
		return []models.ArtistCard{}, err
	}

	var artistCards []models.ArtistCard
	for _, elem := range artistsSQL {
		artistCards = append(artistCards, models.ConvertArtistToCard(elem))
	}

	return artistCards, nil
}

func (uc *Usecase) GetSongsByArtistId(id int) ([]models.SongSQL, error){
	return uc.Repo.GetSongsByArtistId(id)
}

func (uc *Usecase) GetArtistById(id int) (models.Artist, error){
	artist, err := uc.Repo.GetArtistById(id)
	if err != nil {
		return models.Artist{}, err
	}

	a := models.ConvertArtist(artist)
	songs, err := uc.Repo.GetSongsByArtistId(id)
	if err != nil {
		return models.Artist{}, err
	}

	var songCards []models.SongCard
	for _, elem := range songs{
		artistName, err := uc.Repo.GetArtistNameBySongId(elem.Id)
		if err != nil {
			return models.Artist{}, err
		}
		albumName, err := uc.Repo.GetAlbumNameBySongId(elem.Id)
		if err != nil {
			return models.Artist{}, err
		}
		songCards = append(songCards, models.ConvertSongToCard(elem, artistName,
			albumName))
	}
	a.Songs = songCards

	return a, nil
}

func (uc *Usecase) SearchFilm(name string, page int) ([]models.FilmCard, error) {
	filmsSQL, err := uc.Repo.SearchFilm(name, page)
	if err != nil {
		return []models.FilmCard{}, err
	}

	var filmCards []models.FilmCard
	for _, elem := range filmsSQL {
		filmCards = append(filmCards, models.ConvertFilmToCard(elem))
	}
	return filmCards, nil
}

func (uc *Usecase) GetSimilarFilms(id int, page int) ([]models.FilmCard, error) {
	filmsSQL, err := uc.Repo.GetSimilarFilms(id, page)
	if err != nil {
		return []models.FilmCard{}, err
	}

	var filmCards []models.FilmCard
	for _, elem := range filmsSQL {
		filmCards = append(filmCards, models.ConvertFilmToCard(elem))
	}
	return filmCards, nil
}

func (uc *Usecase) ShazamSong(name string) (models.SongShazam, error){
	songs, err := uc.Repo.SearchSong(name, 1)
	if err != nil {
		return models.SongShazam{}, err
	}
	song := songs[0]

	films, err :=  uc.Repo.GetFilmsBySongId(song.Id)
	var filmCards []models.FilmCard
	for _, elem := range films {
		filmCards = append(filmCards, models.ConvertFilmToCard(elem))
	}
	artistName, err := uc.Repo.GetArtistNameBySongId(song.Id)
	if err != nil {
		return models.SongShazam{}, err
	}
	albumName, err := uc.Repo.GetAlbumNameBySongId(song.Id)
	if err != nil {
		return models.SongShazam{}, err
	}

	return models.ConvertSongToShazam(song, artistName, albumName, filmCards), nil
}

func (uc *Usecase) GetFilmById(id int) (models.Film, error){
	film, err := uc.Repo.GetFilmById(id)
	if err != nil {
		return models.Film{}, err
	}

	songs, err := uc.Repo.GetSongsByFilmId(id)
	if err != nil {
		return models.Film{}, err
	}

	var songCards []models.SongCard
	var artistCards []models.ArtistCard

	for _, elem := range songs{
		artist, err := uc.Repo.GetArtistBySongId(elem.Id)
		if err != nil {
			return models.Film{}, err
		}
		toAdd := true
		for _, e := range artistCards {
			if artist.Id == e.Id {
				toAdd = false
				break
			}
		}
		if toAdd {
			artistCards = append(artistCards, models.ConvertArtistToCard(artist))
		}
		albumName, err := uc.Repo.GetAlbumNameBySongId(elem.Id)
		if err != nil {
			return models.Film{}, err
		}
		songCards = append(songCards, models.ConvertSongToCard(elem, artist.Name,
			albumName))
	}



	similar, err := uc.Repo.GetSimilarFilmsNoPage(film.Id)
	if err != nil {
		return models.Film{}, err
	}
	var similarCards []models.FilmCard
	for _, elem := range similar {
		similarCards = append(similarCards, models.ConvertFilmToCard(elem))
	}
	f := models.ConvertFilm(film, artistCards, songCards, similarCards)

	return f, nil
}