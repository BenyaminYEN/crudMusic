package repositories

import "crudMusic/models"

type ArtistRepository interface {
	GetArtists() ([]*models.Artist, error)
}
