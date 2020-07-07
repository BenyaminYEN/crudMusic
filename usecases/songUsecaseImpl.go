package usecases

import (
	"crudMusic/models"
	"crudMusic/repositories"
	"crudMusic/validation"
)

type SongUsecaseImpl struct {
	songRepo repositories.SongRepository
}

func (s SongUsecaseImpl) GetSongs() ([]*models.Song, error) {
	songs, err := s.songRepo.GetSongs()
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (s SongUsecaseImpl) GetSongByGenre(Genre string) ([]*models.Song, error) {
	songs, err := s.songRepo.GetSongByGenre(Genre)
	validation.ValidateInputNotNil(songs)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (s SongUsecaseImpl) GetSongByArtist(Artist string) ([]*models.Song, error) {
	songs, err := s.songRepo.GetSongByArtist(Artist)
	validation.ValidateInputNotNil(songs)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (s SongUsecaseImpl) GetSongByAlbum(Album string) ([]*models.Song, error) {
	songs, err := s.songRepo.GetSongByAlbum(Album)
	validation.ValidateInputNotNil(songs)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func (s SongUsecaseImpl) GetSongByReleaseYear(Year int) ([]*models.Song, error) {
	songs, err := s.songRepo.GetSongByReleaseYear(Year)
	validation.ValidateInputNumber(songs)
	if err != nil {
		return nil, err
	}
	return songs, nil
}

func InitSongUsecase(songRepo repositories.SongRepository) SongUsecase {
	return &SongUsecaseImpl{songRepo}
}
