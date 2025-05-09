package plex

import (
	"fmt"
	"os"
	"path"
	"poster-setter/internal/config"
	"poster-setter/internal/logging"
	"poster-setter/internal/mediux"
	"poster-setter/internal/modals"
	"poster-setter/internal/utils"
)

func DownloadAndUpdatePosters(plex modals.MediaItem, file modals.PosterFile) logging.ErrorLog {

	if !config.Global.SaveImageNextToContent {
		logErr := UpdateSetOnly(plex, file)
		if logErr.Err != nil {
			return logErr
		}
		return logging.ErrorLog{}
	}

	// Check if the temporary folder has the image
	// If it does, we don't need to download it again
	// If it doesn't, we need to download it
	// The image is saved in the temp-images/mediux/full folder with the file ID as the name
	formatDate := file.Modified.Format("20060102")
	fileName := fmt.Sprintf("%s_%s.jpg", file.ID, formatDate)
	filePath := path.Join(mediux.MediuxFullTempImageFolder, fileName)
	exists := utils.CheckIfImageExists(filePath)
	var imageData []byte
	if !exists {
		// Check if the temporary folder exists
		logErr := utils.CheckFolderExists(mediux.MediuxFullTempImageFolder)
		if logErr.Err != nil {
			return logErr
		}
		// Download the image from Mediux
		imageData, _, logErr = mediux.FetchImage(file.ID, formatDate, true)
		if logErr.Err != nil {
			return logErr
		}
		// Add the image to the temporary folder
		err := os.WriteFile(filePath, imageData, 0644)
		if err != nil {
			return logging.ErrorLog{Err: err, Log: logging.Log{Message: fmt.Sprintf("Failed to write image to %s: %v", filePath, err)}}
		}
		logging.LOG.Trace(fmt.Sprintf("Image %s downloaded and saved to temporary folder", file.ID))
	} else {
		// Read the contents of the file
		var err error
		imageData, err = os.ReadFile(filePath)
		if err != nil {
			return logging.ErrorLog{Err: err, Log: logging.Log{Message: fmt.Sprintf("Failed to read image from %s: %v", filePath, err)}}
		}
	}

	// Check the Plex Item type
	// If it is a movie or show, handle the poster/backdrop/seasonPoster/titlecard accordingly

	newFilePath := ""
	newFileName := ""

	if plex.Type == "movie" {
		// Handle movie-specific logic
		newFilePath = path.Dir(plex.Movie.File.Path)
		if file.Type == "poster" {
			newFileName = "poster.jpg"
		} else if file.Type == "backdrop" {
			newFileName = "backdrop.jpg"
		}
	} else if plex.Type == "show" {
		// Handle show-specific logic
		newFilePath = plex.Series.Location
		if file.Type == "poster" {
			newFileName = "poster.jpg"
		} else if file.Type == "backdrop" {
			newFileName = "backdrop.jpg"
		} else if file.Type == "seasonPoster" {
			newFilePath = path.Join(newFilePath, fmt.Sprintf("Season %s", utils.Get2DigitNumber(int64(file.Season.Number))))
			newFileName = fmt.Sprintf("Season%s.jpg", utils.Get2DigitNumber(int64(file.Season.Number)))
		} else if file.Type == "titlecard" {
			// For titlecards, get the file path from Plex
			episodePath := getEpisodePathFromPlex(plex, file)
			if episodePath != "" {
				newFilePath = path.Dir(episodePath)
				newFileName = path.Base(episodePath)
				newFileName = newFileName[:len(newFileName)-len(path.Ext(newFileName))] + ".jpg"
			} else {
				return logging.ErrorLog{Err: fmt.Errorf("episode path is empty"), Log: logging.Log{Message: "No episode path found for titlecard"}}
			}
		}
	}

	// Ensure the directory exists
	err := os.MkdirAll(newFilePath, os.ModePerm)
	if err != nil {
		return logging.ErrorLog{Err: err, Log: logging.Log{Message: fmt.Sprintf("Failed to create directory %s", newFilePath)}}
	}

	// Create the new file
	newFile, err := os.Create(path.Join(newFilePath, newFileName))
	if err != nil {
		return logging.ErrorLog{Err: err, Log: logging.Log{Message: fmt.Sprintf("Failed to create file %s", newFileName)}}
	}
	defer newFile.Close()

	// Write the image data to the file
	_, err = newFile.Write(imageData)
	if err != nil {
		return logging.ErrorLog{Err: err, Log: logging.Log{Message: fmt.Sprintf("Failed to write image data to file %s", newFileName)}}
	}

	// If cacheImages is False, delete the image from the temporary folder
	if !config.Global.CacheImages {
		err := os.Remove(filePath)
		if err != nil {
			logging.LOG.Error(fmt.Sprintf("Failed to delete image from temporary folder: %v", err))
		}
	}

	// Determine the itemRatingKey
	itemRatingKey := getItemRatingKey(plex, file)
	if itemRatingKey == "" {
		logging.LOG.Error(fmt.Sprintf("Item rating key is empty for '%s'", plex.Title))
		return logging.ErrorLog{Err: fmt.Errorf("item rating key is empty"), Log: logging.Log{Message: fmt.Sprintf("Item rating key is empty for '%s'", plex.Title)}}
	}
	refreshPlexItem(itemRatingKey)
	posterKey, logErr := getPosters(itemRatingKey)
	if logErr.Err != nil {
		return logErr
	}
	setPoster(itemRatingKey, posterKey, file.Type)

	return logging.ErrorLog{}
}
