package cmd

import (
	"context"
	"github.com/fenpaws/MELE-Mod-Downloader/internal"
	v1 "github.com/fenpaws/MELE-Mod-Downloader/internal/api"
	"github.com/fenpaws/MELE-Mod-Downloader/internal/nxm"
	"github.com/fenpaws/MELE-Mod-Downloader/internal/utils"
	"github.com/fenpaws/MELE-Mod-Downloader/pkg/models"
	"github.com/fenpaws/MELE-Mod-Downloader/pkg/nexusmods"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the GOLE server and process mod collections",
	Long: `This command starts the GOLE (Mass Effect: Legendary Edition Mod Downloader) server, 
loads and parses the specified mod collection, and processes the mods for download.`,

	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// load and parse the mod collection
		biq2, err := internal.LoadAndParseBiq2(args[0])
		if err != nil {
			log.WithError(err).Fatal("Could not load mod collections")
		}

		// run the main http server
		runServer(ctx, biq2)
	},
}

func runServer(ctx context.Context, biq2 *models.ModCollection) {
	log.WithContext(ctx).Info("Startup GO:LE (Mass Effect: Legendary Mod Downloader)")

	modChannel := make(chan string)

	ginRouter := gin.Default()

	apiRouter := v1.NewAPIRouter(ginRouter.Group("/api/v1"), modChannel)
	apiRouter.InitializeRoutes()

	server := &http.Server{
		Addr:    ":8081",
		Handler: ginRouter,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.WithContext(ctx).WithError(err).Fatal("failed to start server")
		}
	}()

	err := execute(ctx, biq2, modChannel)
	if err != nil {
		return
	}

	time.Sleep(1 * time.Minute)
}

func execute(ctx context.Context, biq2 *models.ModCollection, modChannel chan string) error {

	modDownloadList := make([]models.ModDownloadInfo, 0)

	// Process all mods in the pack
	for _, mod := range biq2.Mods {
		log.WithContext(ctx).WithField("link", mod.Downloadlink).Debug("Mod download link")
		nexusFileDownloadLink := mod.Downloadlink + "&nmm=1"

		// 1. Open base mod URL in browser
		err := utils.OpenURL(nexusFileDownloadLink)
		if err != nil {
			log.WithContext(ctx).WithError(err).Error("Could not open download link")
		}

		// 2. Read the callback from nexus mods
		mxnUrl := <-modChannel

		// 3. Convert callback URL to propper URL
		parsedURL, err := url.Parse(mxnUrl)
		if err != nil {
			log.WithError(err).Error("Could not parse URL")
		}

		// 4. Extract relevant information to generate the download link
		nxmInfo := nxm.HandleNxmURL(parsedURL)
		log.WithField("modId", nxmInfo.ModID).WithField("fileId", nxmInfo.FileID).WithField("key", nxmInfo.Key).WithField("expires", nxmInfo.Expires).Debug("Extracted NXM info")

		// 5. Get mod file link from nexus mods
		nexusClient := nexusmods.NewNexusModsClient(os.Getenv("nexusApiKey"))
		link, err := nexusClient.GenerateDownloadLink("masseffectlegendaryedition", nxmInfo.ModID, nxmInfo.FileID, nxmInfo.Key, nxmInfo.Expires)
		if err != nil {
			log.WithError(err).Error("Could not generate download link")
		}
		log.WithField("link", link.URI).Debug("Mod file link")

		// 6. Add download files to download slice
		modDownloadList = append(modDownloadList, models.ModDownloadInfo{
			ModName: mod.Modname,
			URL:     link.URI,
		})
	}

	log.WithContext(ctx).WithField("found_mods", len(modDownloadList)).Info("Successfully found mods")

	var wg sync.WaitGroup
	sem := make(chan struct{}, 3)
	for _, modDownload := range modDownloadList {
		wg.Add(1)
		sem <- struct{}{}
		go func(url models.ModDownloadInfo) {
			defer func() { <-sem }()
			internal.DownloadFile(ctx, modDownload.URL, modDownload.ModName, "./mods", &wg)
		}(modDownload)
	}

	wg.Wait()

	return nil
}
