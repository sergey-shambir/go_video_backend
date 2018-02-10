package main

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"github.com/segmentio/ksuid"
)

const (
	videoUrlPrefix = "content/"
)

// VideoListItem - item for the list video API.
// field Duration - duration in seconds
type VideoListItem struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Duration  int    `json:"duration"`
	Thumbnail string `json:"thumbnail"`
}

type VideoList []VideoListItem

type VideoInfo struct {
	VideoListItem
	URL string `json:"url"`
}

func getList(c APIContext) error {
	db, err := openMysqlConn()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := dbGetVideoList(db)
	if err != nil {
		return err
	}

	var result VideoList
	for _, row := range rows {
		video := VideoListItem{
			ID:        row.key,
			Name:      row.title,
			Duration:  row.duration,
			Thumbnail: row.thumbnail,
		}
		result = append(result, video)
	}
	return c.WriteJSON(result)
}

func getVideo(c APIContext) error {
	id := c.Vars()["id"]

	db, err := openMysqlConn()
	if err != nil {
		return err
	}
	defer db.Close()

	row, err := dbVideoInfo(db, id)
	if err != nil {
		return err
	}

	info := &VideoInfo{
		VideoListItem: VideoListItem{
			ID:        row.key,
			Name:      row.title,
			Duration:  row.duration,
			Thumbnail: row.thumbnail,
		},
		URL: row.url,
	}
	return c.WriteJSON(info)
}

func uploadVideo(w http.ResponseWriter, r *http.Request) error {
	srcFile, header, err := r.FormFile("file[]")
	if err != nil {
		return err
	}
	// TODO: check contentType := header.Header.Get("Content-Type")
	filename := header.Filename

	suid := ksuid.New()
	dstDir := suid.String()
	dstFilePath := path.Join(dstDir, "index"+path.Ext(filename))

	logrus.WithFields(logrus.Fields{
		"filename":    filename,
		"dstFilePath": dstFilePath,
	}).Info("uploadVideo")

	err = os.MkdirAll(dstDir, 0777)
	if err != nil {
		return err
	}

	dstFile, err := os.OpenFile(dstFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	db, err := openMysqlConn()
	if err != nil {
		return err
	}
	defer db.Close()

	row := VideoRow{
		key:   suid.String(),
		title: path.Base(filename),
		url:   videoUrlPrefix + dstFilePath,
	}
	err = dbRegisterVideo(db, row)
	if err != nil {
		return err
	}

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
