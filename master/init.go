package master

import (
	"crudMusic/controller"
	"crudMusic/repositories"
	"crudMusic/usecases"
	"database/sql"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, db *sql.DB) {
	songRepo := repositories.InitSongRepositoryImpl(db)
	songUseCase := usecases.InitSongUsecase(songRepo)

	artistRepo := repositories.InitArtistRepositoryImpl(db)
	artistUseCase := usecases.InitArtistUsecase(artistRepo)

	controller.SongController(r, songUseCase)
	controller.ArtistController(r, artistUseCase)
}
