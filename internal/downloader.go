package internal

import (
	"context"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func DownloadFile(ctx context.Context, url, modName, downloadLocation string, wg *sync.WaitGroup) {
	log.WithContext(ctx).WithField("mod_name", modName).Info("Downloading mod")

	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Failed to download file")
		return
	}
	defer resp.Body.Close()

	fileName := filepath.Base(modName + ".7zip")
	filePath := filepath.Join(downloadLocation, fileName)
	out, err := os.Create(filePath)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Failed to create file")
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Failed to save file")
		return
	}

	log.WithContext(ctx).WithField("file_location", filePath).Info("Successfully downloaded file")
}
