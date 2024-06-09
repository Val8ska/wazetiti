package common

import (
	"log"
	"os"
	"path/filepath"
)

func WalkThrough(dir string) (listFiles []string) {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == ".json" {
				listFiles = append(listFiles, path)
			}
			return nil
		})
	if err != nil {
		log.Println("[WAZE realtime] filepath walk", err)
	}
	return listFiles
}
