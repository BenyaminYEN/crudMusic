package usecases

import "crudMusic/models"

type SongUsecase interface {
	GetSongs() ([]*models.Song, error)
	GetSongByGenre(Genre string) ([]*models.Song, error)
	GetSongByArtist(Artist string) ([]*models.Song, error)
	GetSongByAlbum(Album string) ([]*models.Song, error)
	GetSongByReleaseYear(Year int) ([]*models.Song, error)
}
