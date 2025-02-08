package repository

import (
	"API-RESTful/models"
	"database/sql"
)

type AlbumRepository struct {
	DB *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{DB: db}
}
func (r *AlbumRepository) GetAlbums() ([]models.Album, error) {
	rows, err := r.DB.Query("SELECT id, title, artist, price FROM albums")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var albums []models.Album
	for rows.Next() {
		var album models.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return albums, nil
}
func (r *AlbumRepository) AddAlbum(album models.Album) error {
	_, err := r.DB.Exec("INSERT INTO albums (id, title, artist, price) VALUES (?,?,?,?)", album.ID, album.Title, album.Artist, album.Price)
	return err
}
