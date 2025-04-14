package download

import (
	"fmt"
	"poster-setter/internal/database"
	"poster-setter/internal/logging"
	"poster-setter/internal/mediux"
	"poster-setter/internal/modals"
	"poster-setter/internal/plex"
)

func CheckForUpdatesToPosters() {

	items, logErr := database.GetAllItemsFromDatabase()
	if logErr.Err != nil {
		logging.LOG.ErrorWithLog(logErr)
		return
	}

	for _, item := range items {
		if item.AutoDownload {
			logging.LOG.Debug(fmt.Sprintf("Checking for updates to posters for '%s'", item.Plex.Title))
			var updatedSet modals.PosterSet
			var logErr logging.ErrorLog
			if item.Set.Type == "movieset" || item.Set.Type == "collection" || item.Set.Type == "showset" {
				updatedSet, logErr = mediux.FetchSetByID(item.Set.Type, item.Set.ID)
			} else {
				logging.LOG.Error(fmt.Sprintf("Set for '%s' is not a valid type: %s", item.Plex.Title, item.Set.Type))
			}
			if logErr.Err != nil {
				logging.LOG.ErrorWithLog(logErr)
				continue
			}

			updated := compareLastUpdateToUpdateSetDateUpdated(item.LastUpdate, updatedSet.DateUpdated)
			if updated {
				logging.LOG.Info(fmt.Sprintf("Posters for '%s' have been updated. Downloading new posters...", item.Plex.Title))
				// Download the new posters and update Plex

				item.Set.Files = plex.FilterAndSortFiles(updatedSet.Files, item.SelectedTypes)
				for _, file := range item.Set.Files {
					logErr := plex.DownloadAndUpdateSet(item.Plex, file)
					if logErr.Err != nil {
						logging.LOG.ErrorWithLog(logErr)
						continue
					}
				}
				// Update the item in the database with the new info
				logErr = database.UpdateAutoDownloadItem(modals.ClientMessage{
					Plex:          item.Plex,
					Set:           updatedSet,
					AutoDownload:  item.AutoDownload,
					SelectedTypes: item.SelectedTypes,
				})
				if logErr.Err != nil {
					logging.LOG.ErrorWithLog(logErr)
					continue
				}
			}
		}
	}
}
