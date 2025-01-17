package internal

import (
	"context"
	"fmt"
	"github.com/fenpaws/MELE-Mod-Downloader/internal/utils"
	"github.com/schollz/progressbar/v3"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func DownloadFile(ctx context.Context, url, modName, downloadLocation string, wg *sync.WaitGroup, mpBar *utils.MultiProgressBar) {

	// TODO: Check if file exist
	// TODO: MD5 hsah check
	// TODO: better downloader?

	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Failed to create request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.WithContext(ctx).WithError(err).Error("Failed to download file")
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

	bar := progressbar.NewOptions64(
		resp.ContentLength,
		progressbar.OptionSetDescription(fmt.Sprintf("Downloading %-20s", modName)),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(10),
		progressbar.OptionThrottle(250*time.Millisecond),
		progressbar.OptionShowCount(),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetRenderBlankState(false),
	)
	_ = mpBar.Add(bar)

	io.Copy(io.MultiWriter(out, bar), resp.Body)
}
