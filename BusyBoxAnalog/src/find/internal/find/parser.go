package find

import (
	"fmt"
	"goday02/src/find/internal/input"
	"io/fs"
	"os"
	"path/filepath"
)

func Find(cfg input.CLIcfg) {
	if _, err := os.ReadDir(cfg.Path); err != nil {
		return
	}
	walkFunc := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if cfg.FileFlag {
			infoFile, err := os.Lstat(path)
			if err != nil {
				return err
			}
			if infoFile.Mode()&os.ModeSymlink == 0 && !d.IsDir() {
				if !cfg.ExtFlag {
					fmt.Println(path)
				} else if cfg.Ext == filepath.Ext(path) {
					fmt.Println(path)
				}
			}
		}
		if cfg.DirFlag {
			if d.IsDir() {
				fmt.Println(path)
			}
		}
		if cfg.SymFlag {
			infoFile, err := os.Lstat(path)
			if err != nil {
				return err
			}
			if infoFile.Mode()&os.ModeSymlink != 0 {
				target, err := os.Readlink(path)
				if err != nil {
					return err
				}
				if _, err := os.Stat(path); os.IsNotExist(err) {
					fmt.Println(path + " -> [broken]")
				} else {
					fmt.Println(path + " -> " + target)
				}
			}
		}
		return nil
	}
	filepath.WalkDir(cfg.Path, walkFunc)
}
