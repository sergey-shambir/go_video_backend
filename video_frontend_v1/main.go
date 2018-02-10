package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/sirupsen/logrus"
)

const (
	defaultServerURL = "localhost:8081"
	defaultWorkdir   = "../../src/workshops2018/web/content"
)

func main() {
	// TODO: don't make chdir and load workdir from config / env
	err := chdirWorkdir()
	if err != nil {
		panic(err)
	}

	killChan := getKillSignalChan()

	logFile := openFileLogger()
	defer logFile.Close()
	server := startServer(defaultServerURL)
	waitForKillSignal(killChan)
	server.Shutdown(context.Background())
}

func chdirWorkdir() error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	workdir := path.Join(ex, defaultWorkdir)
	return os.Chdir(workdir)
}

func getKillSignalChan() chan os.Signal {
	killChan := make(chan os.Signal, 1)
	signal.Notify(killChan, os.Kill, os.Interrupt, syscall.SIGTERM)
	return killChan
}

func waitForKillSignal(killChan <-chan os.Signal) {
	killSignal := <-killChan
	switch killSignal {
	case os.Interrupt:
		logrus.Info("got SIGINT, shutting down...")
	case syscall.SIGTERM:
		logrus.Info("got SIGTERM, shutting down...")
	}
}

func startServer(serverURL string) *http.Server {
	logrus.WithFields(logrus.Fields{"url": serverURL}).Info("starting server")
	server := &http.Server{Addr: serverURL, Handler: newRouter()}
	go func() {
		logrus.Fatal(server.ListenAndServe())
	}()

	return server
}
