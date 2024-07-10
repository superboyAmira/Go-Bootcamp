package backupgenerator

import (
	"log/slog"
	"os"
	"path/filepath"
)

func MakeSnapshot(snapPath string, log *slog.Logger) {
	var snapDir string = "../.."
	file, err := os.Create(snapPath)
	if err != nil {
		log.Debug(err.Error())
		return
	}
	defer file.Close()

	err = filepath.Walk(snapDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			_, err := file.WriteString(path + "\n")
			if err != nil {
				log.Debug("failed to write to snapshot file")
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Debug("error walking the path")
		return
	}
}