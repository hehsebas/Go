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
	rows, err := r.DB.Query("SELECT * FROM albums")
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
func (r *AlbumRepository) DeleteAlbum(id int) error {
	_, err := r.DB.Exec("DELETE FROM albums WHERE id = ?", id)
	return err
}
func (r *AlbumRepository) ModifyAlbum(album models.Album, id int) error {
	_, err := r.DB.Exec("UPDATE albums SET id=?, title = ?, artist = ?, price = ? WHERE id = ?", album.ID, album.Title, album.Artist, album.Price, id)
	return err
}
