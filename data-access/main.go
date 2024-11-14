package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to the database!")

	albums, err := albumsByArtist("John Coltrane")

	if err != nil {
		log.Fatal("err")
	}

	fmt.Printf("Albums found: %v\n", albums)

	album, err := albumById(2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album found by id: %v\n", album)

	id, err := addAlbum(Album{
		Title:  "American Idiot",
		Artist: "Green Day",
		Price:  49.99,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Album added with id: %v\n", id)
}

func albumsByArtist(artistName string) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", artistName)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", artistName, err)
	}
	defer rows.Close()
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", artistName, err)
		}
		albums = append(albums, album)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", artistName, err)
	}
	return albums, nil
}

func albumById(albumId int64) (Album, error) {
	var album Album
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", albumId)

	row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)

	if err := row.Err(); err != nil {
		return album, fmt.Errorf("albumById %q: %v", albumId, err)
	}
	return album, nil
}

func addAlbum(album Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album(title, artist, price) VALUES ( ?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum %v", err)
	}
	id, err := result.LastInsertId()

	if err != nil {
		return 0, fmt.Errorf("addAlbum :v", err)
	}
	return id, nil
}

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}
