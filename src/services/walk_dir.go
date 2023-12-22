package services

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/wahyuhadi/supply-chain/models"
)

func Walk(root string, conf models.Optional) {
	maxDepth := 2
	rootDir := root
	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && strings.Count(path, string(os.PathSeparator)) == maxDepth {
			log.Println("\n[+] Cheking vendor package ", path)
			scanning(path, conf)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
}
