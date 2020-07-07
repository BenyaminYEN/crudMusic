package repositories

import (
	"crudMusic/models"
	"database/sql"
)

//membuat struct keseluruhan dari DB Song
type SongRepoImpl struct {
	db *sql.DB
}

func (s *SongRepoImpl) GetSongs() ([]*models.Song, error) {
	var songs []*models.Song
	query := "SELECT * FROM song"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		song := models.Song{}
		err := rows.Scan(&song.SongID, &song.Title, &song.AlbumID, &song.Genre, &song.Artist, &song.Album, &song.ReleaseYear)

		if err != nil {
			return nil, err
		}

		songs = append(songs, &song)

	}

	return songs, nil
}

func (s *SongRepoImpl) GetSongByGenre(Genre string) ([]*models.Song, error) {
	var songs []*models.Song
	query := "SELECT * FROM song WHERE genre = ?"
	rows, err := s.db.Query(query, Genre)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		song := models.Song{}
		err := rows.Scan(&song.SongID, &song.Title, &song.AlbumID, &song.Genre, &song.Artist, &song.Album, &song.ReleaseYear)

		if err != nil {
			return nil, err
		}

		songs = append(songs, &song)

	}

	return songs, nil
}

func (s *SongRepoImpl) GetSongByArtist(Artist string) ([]*models.Song, error) {
	var songs []*models.Song
	query := "SELECT * FROM song WHERE artist = ?"
	rows, err := s.db.Query(query, Artist)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		song := models.Song{}
		err := rows.Scan(&song.SongID, &song.Title, &song.AlbumID, &song.Genre, &song.Artist, &song.Album, &song.ReleaseYear)

		if err != nil {
			return nil, err
		}

		songs = append(songs, &song)

	}

	return songs, nil
}

func (s *SongRepoImpl) GetSongByAlbum(Album string) ([]*models.Song, error) {
	var songs []*models.Song
	query := "SELECT * FROM song WHERE album = ?"
	rows, err := s.db.Query(query, Album)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		song := models.Song{}
		err := rows.Scan(&song.SongID, &song.Title, &song.AlbumID, &song.Genre, &song.Artist, &song.Album, &song.ReleaseYear)

		if err != nil {
			return nil, err
		}

		songs = append(songs, &song)

	}

	return songs, nil
}

func (s *SongRepoImpl) GetSongByReleaseYear(Year int) ([]*models.Song, error) {
	var songs []*models.Song
	query := "SELECT * FROM song WHERE release_year = ?"
	rows, err := s.db.Query(query, Year)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		song := models.Song{}
		err := rows.Scan(&song.SongID, &song.Title, &song.AlbumID, &song.Genre, &song.Artist, &song.Album, &song.ReleaseYear)

		if err != nil {
			return nil, err
		}

		songs = append(songs, &song)

	}

	return songs, nil
}

func InitSongRepositoryImpl(db *sql.DB) SongRepository {
	return &SongRepoImpl{db}
}
