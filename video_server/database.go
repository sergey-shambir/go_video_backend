package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUser     = "videoservice"
	mysqlPassword = "1234"
	mysqlHost     = ""
)

// VideoRow - represents one `video` SQL table row.
type VideoRow struct {
	key       string
	title     string
	url       string
	duration  int
	thumbnail string
}

// TODO: extract interface VideoRepository and implement it to keep `*sql.DB` as filed

func openMysqlConn() (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@%s/workshop_video_server", mysqlUser, mysqlPassword, mysqlHost))
}

func dbRegisterVideo(db *sql.DB, row VideoRow) error {
	q := `INSERT INTO video SET video_key = ?, title = ?,  url = ?`
	rows, err := db.Query(q, row.key, row.title, row.url)
	if err != nil {
		rows.Close()
	}
	return err
}

func dbGetVideoList(db *sql.DB) ([]VideoRow, error) {
	// TODO: no need to ask 'url'
	q := `SELECT video_key, title, url, thumbnail_url, duration FROM video`
	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	return scanVideoRows(rows)
}

func dbVideoInfo(db *sql.DB, key string) (*VideoRow, error) {
	// TODO: no need to ask 'video_key'
	q := `SELECT video_key, title, url, thumbnail_url, duration FROM video WHERE video_key = ?`
	rows, err := db.Query(q, key)
	videos, err := scanVideoRows(rows)
	if err != nil {
		return nil, err
	}
	if len(videos) != 1 {
		return nil, errors.New("video with key '" + key + "' not found")
	}
	return &videos[0], nil
}

func scanVideoRows(rows *sql.Rows) ([]VideoRow, error) {
	var videos []VideoRow
	for rows.Next() {
		var video VideoRow
		err := rows.Scan(&video.key, &video.title, &video.url, &video.thumbnail, &video.duration)
		if err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}
	return videos, nil
}
