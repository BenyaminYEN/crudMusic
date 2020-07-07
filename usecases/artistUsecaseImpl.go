package usecases

import (
	"crudMusic/models"
	"crudMusic/repositories"
)

type ArtistUsecaseImpl struct {
	artistRepo repositories.ArtistRepository
}

func (s ArtistUsecaseImpl) GetArtists() ([]*models.Artist, error) {
	artists, err := s.artistRepo.GetArtists()
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func InitArtistUsecase(artistRepo repositories.ArtistRepository) ArtistUsecase {
	return &ArtistUsecaseImpl{artistRepo}
}
