package repositories

import (
	"crudMusic/models"
	"database/sql"
)

//membuat struct keseluruhan dari DB Song
type ArtistRepoImpl struct {
	db *sql.DB
}

func (s *ArtistRepoImpl) GetArtists() ([]*models.Artist, error) {
	var artists []*models.Artist
	query := "SELECT * FROM artist"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		artist := models.Artist{}
		err := rows.Scan(&artist.ArtistID, &artist.ArtistName)

		if err != nil {
			return nil, err
		}

		artists = append(artists, &artist)

	}

	return artists, nil
}

func InitArtistRepositoryImpl(db *sql.DB) ArtistRepository {
	return &ArtistRepoImpl{db}
}
