package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/wahyuhadi/supply-chain/models"
	"github.com/wahyuhadi/supply-chain/services"
)

func PHP(config models.Optional) {
	if config.Conf == "" {
		ScanPhPRootDir(config)
	}

	if config.ConfScan != nil {
		// if config.ConfScan.ScanType == "dir" {
		// 	scanDirByConfig(config, config.ConfScan.Projects)
		// }

		if config.ConfScan.ScanType == "repo" {
			phpScanRepoByConfig(config, config.ConfScan.Projects)
		}

	}
}
func phpScanRepoByConfig(config models.Optional, repos []string) {

	for _, projects_repo := range repos {
		repos := projects_repo[strings.LastIndex(projects_repo, "/")+1:]
		log.Println(fmt.Sprintf("[+] Clone repo %s to folder %s ", projects_repo, repos))

		_, err := git.PlainClone(repos, false, &git.CloneOptions{
			URL: projects_repo,
		})

		if err != nil {
			log.Println("[!] Error when clone repo ", projects_repo)
			fmt.Println(err)
			continue
		}

		os.Chdir(repos)
		services.Scanning_root_path(config)
		ExtractResult(config)

	}
}

func ScanPhPRootDir(conf models.Optional) {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	dirs := strings.Split(dir, "/")
	conf.Dir = dirs[len(dirs)-1]
	services.Scanning_root_path(conf)
	ExtractResult(conf)
}
