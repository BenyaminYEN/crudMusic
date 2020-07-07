package models

type Song struct {
	SongID      int    `json:"song_id"`
	Title       string `json:"title"`
	AlbumID     int    `json:"album_id"`
	Genre       string `json:"genre"`
	Artist      string `json:"artist"`
	Album       string `json:"album"`
	ReleaseYear int    `json:"release_year"`
}
