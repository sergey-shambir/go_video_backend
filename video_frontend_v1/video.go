package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/segmentio/ksuid"
)

// VideoListItem - item for the list video API.
type VideoListItem struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	Duration  time.Duration `json:"duration"`
	Thumbnail string        `json:"thumbnail"`
}

type VideoList []VideoListItem

type VideoInfo struct {
	VideoListItem
	URL string `json:"url"`
}

func getList(c APIContext) error {
	list := &VideoList{
		VideoListItem{
			ID:        "d290",
			Name:      "Black Retrospective Woman",
			Duration:  time.Duration(15) * time.Second,
			Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		},
	}
	return c.WriteJSON(list)
}

func getVideo(c APIContext) error {
	info := &VideoInfo{
		VideoListItem: VideoListItem{
			ID:        "d290",
			Name:      "Black Retrospective Woman",
			Duration:  time.Duration(15) * time.Second,
			Thumbnail: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg",
		},
		URL: "/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4",
	}
	return c.WriteJSON(info)
}

func uploadVideo(w http.ResponseWriter, r *http.Request) error {
	srcFile, header, err := r.FormFile("file[]")
	if err != nil {
		return err
	}
	// contentType := header.Header.Get("Content-Type")
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

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
