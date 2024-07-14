package input

import (
	"flag"
	"path/filepath"
)

type CLIcfg struct {
	Path_f          string
	FileType_f      string
	Path_old        string
	Path_new        string
	FileType_old    string
	FileType_new    string
	Path_backup_old string
	Path_backup_new string
	Snapshot        string
}

/*
 * Only one mode of operation is supported(take a look to str №27)
 */
func ParseFile() (fileInfo CLIcfg) {
	path := flag.String("f", "", "path to file")
	old_path := flag.String("old", "", "path to previus file (task 'to compare') or new snapshot path(task 3)")
	new_path := flag.String("new", "", "path to new file (task 'to compare') or old snapshot path(task 3)")
	snapshot := flag.String("snap-name", "", "file name to make snapshot")
	flag.CommandLine.Init("flag settings didn`t stop executing a app", flag.ContinueOnError)

	flag.Parse()
	fileInfo.Path_f = *path
	fileInfo.FileType_f = filepath.Ext(*path)
	if filepath.Ext(*new_path) == ".txt" && filepath.Ext(*old_path) == ".txt" {
		fileInfo.Path_backup_old = *old_path
		fileInfo.Path_backup_new = *new_path
	} else {
		fileInfo.Path_old = *old_path
		fileInfo.Path_new = *new_path
		fileInfo.Path_backup_old = ""
		fileInfo.Path_backup_new = ""
		fileInfo.FileType_new = filepath.Ext(*new_path)
		fileInfo.FileType_old = filepath.Ext(*old_path)
	}
	fileInfo.Snapshot = *snapshot

	return
}
