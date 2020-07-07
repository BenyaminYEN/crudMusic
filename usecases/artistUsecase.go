package usecases

import "crudMusic/models"

type ArtistUsecase interface {
	GetArtists() ([]*models.Artist, error)
}
