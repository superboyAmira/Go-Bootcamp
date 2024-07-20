package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"goday02/src/rotate/input"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	paths, archive, err := input.GetFile()
	if err != nil {
		return
	}
	for _, path := range paths {
		archiveFile(path, archive)
	}
}

func archiveFile(path string, archive *string) {
	fileOpened, err := os.Open(path)
	if err != nil {
		return
	}

	fileToArchive, err := os.Stat(path)
	if err != nil {
		return
	}
	tarName := fmt.Sprintf(strings.Split(path, ".")[0]+"_%d.tar.gz", fileToArchive.ModTime().Unix())

	var touchedFile *os.File
	if archive != nil {
		if _, err := os.Stat(*archive); os.IsNotExist(err) {
			if err := os.MkdirAll(*archive, os.ModePerm); err != nil {
				return
			}
		}
		touchedFilePath := filepath.Join(*archive, tarName)
		touchedFile, err = os.Create(touchedFilePath)
		if err != nil {
			return
		}
	} else {
		touchedFile, err = os.Create(tarName)
		if err != nil {
			return
		}
	}
	defer touchedFile.Close()
	gzipWriter := gzip.NewWriter(touchedFile)
	defer gzipWriter.Close()
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()
	header, err := tar.FileInfoHeader(fileToArchive, fileToArchive.Name())
	if err != nil {
		return
	}
	if err := tarWriter.WriteHeader(header); err != nil {
		return
	}
	if _, err := io.Copy(tarWriter, fileOpened); err != nil {
		return
	}
}
