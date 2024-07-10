package main

import (
	backupgenerator "goday01/internal/backup_generator"
	"goday01/internal/comparator"
	"goday01/internal/converter"
	"goday01/internal/input"
	"log/slog"
	"os"
)

func main() {
	log := setUpLogger()

	fileInfo := input.ParseFile()

	if fileInfo.FileType_f != "" {
		converter.Convert(fileInfo, log)
	}
	if fileInfo.FileType_new != "" && fileInfo.FileType_old != "" {
		comparator.Compare(fileInfo, log)
	}
	if fileInfo.Snapshot != "" {
		backupgenerator.MakeSnapshot(fileInfo.Snapshot, log)
	}
	if fileInfo.Path_backup_new != "" && fileInfo.Path_backup_old != "" {

	}
}

func setUpLogger() *slog.Logger {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	return log
}
