package main

import (
	"context"
	"github.com/fenpaws/MELE-Mod-Downloader/internal"
	"github.com/fenpaws/MELE-Mod-Downloader/internal/nxm"
	"github.com/fenpaws/MELE-Mod-Downloader/internal/utils"
	"github.com/fenpaws/MELE-Mod-Downloader/pkg/nexusmods"
	log "github.com/sirupsen/logrus"
	url2 "net/url"
	"os"
)

func init() {
	// Set up the logger based on the configuration
	utils.SetupLogger("TRACE", "PLAIN")

}

// TODO: Register mxn:// handler,
// TODO: Handel callbacks in background
// TODO: UI?
// TODO: configuration
// TODO: save files to proper location

func main() {
	var ctx = context.Background()

	log.WithContext(ctx).Info("Startup GO:LE (Mass Effect: Legendary Mod Downloader)")

	// 1. Load the Collection
	biq2, err := internal.LoadAndParseBiq2("./LE1 Mod Collection (by Synth).biq2")
	if err != nil {
		log.WithError(err).Fatal("Could not load LE1 Mod Collections")
	}

	// 2. Open the mod in the users browser to get the callback
	log.WithContext(ctx).WithField("link", biq2.Mods[0].Downloadlink).Debug("Mod download link")
	nxmDowloadLink := biq2.Mods[0].Downloadlink + "&nmm=1"

	err = utils.OpenURL(nxmDowloadLink)
	if err != nil {
		log.WithError(err).Fatal("Could not open download link")
	}

	// 3. Handel the nxm callback and extract the relevant fields
	url, err := url2.Parse("nxm://masseffectlegendaryedition/mods/661/files/6665?key=i1O1YbxMRsFLcKh_V4p_OA&expires=1737208657&user_id=8748009")
	if err != nil {
		log.WithError(err).Fatal("Could not parse URL")
	}

	nxmInfo := nxm.HandleNxmURL(url)
	log.WithField("modId", nxmInfo.ModID).WithField("fileId", nxmInfo.FileID).WithField("key", nxmInfo.Key).WithField("expires", nxmInfo.Expires).Debug("Extracted NXM info")

	// 4. Generate Download link based on callback
	nexusClient := nexusmods.NewNexusModsClient(os.Getenv("nexusApiKey"))
	link, err := nexusClient.GenerateDownloadLink("masseffectlegendaryedition", nxmInfo.ModID, nxmInfo.FileID, nxmInfo.Key, nxmInfo.Expires)
	if err != nil {
		log.WithError(err).Fatal("Could not generate download link")
	}
	log.WithField("link", link.URI).Debug("Generated link")

	// 5. Add do download Queue

	// 6. Download based on Queue

}
