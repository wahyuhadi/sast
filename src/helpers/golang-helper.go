package helpers

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/wahyuhadi/supply-chain/models"
	"github.com/wahyuhadi/supply-chain/services"
)

// helper for lang go

var (
	folder_name = "golang-s-chain-attack"
)

func GoLang(config models.Optional) {
	if config.Conf == "" {
		scanOnRootPath(config)
	}

	if config.ConfScan != nil {
		if config.ConfScan.ScanType == "dir" {
			scanDirByConfig(config, config.ConfScan.Projects)
		}

		if config.ConfScan.ScanType == "repo" {
			scanRepoByConfig(config, config.ConfScan.Projects)
		}

	}
}

func scanRepoByConfig(config models.Optional, repos []string) {
	cwd, _ := os.Getwd()

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

		log.Println("[+] Scanning project directory in ", projects_repo)
		log.Println("[+] Golang deteted")
		log.Println("[+] Run go mod tidy")
		cmdt := exec.Command("go", "mod", "tidy")
		cmdt.Run()
		log.Println("[+] Create Vendor")
		cmd := exec.Command("go", "mod", "vendor")
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			log.Println("[!] Error when running go mod vendor")
			continue
		}

		os.Rename("vendor", folder_name)
		services.Walk(folder_name, config)
		err = os.RemoveAll(folder_name)
		if err != nil { // handle error
			continue
		}
		log.Println("[+] Deleted vendor folder ")
		os.Chdir(cwd)
		os.RemoveAll(repos)

	}
}

func scanDirByConfig(config models.Optional, dir []string) {
	for _, projects_dir := range dir {
		services.Scanning_root_path(config)

		os.Chdir(projects_dir)
		log.Println("[+] Scanning project directory in ", projects_dir)
		log.Println("[+] Golang deteted")
		log.Println("[+] Run go mod tidy")
		cmdt := exec.Command("go", "mod", "tidy")
		err := cmdt.Run()
		if err != nil {
			log.Println("[!] Error when running go mod tidy")
			continue
		}
		log.Println("[+] Create Vendor")
		cmd := exec.Command("go", "mod", "vendor")
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			log.Println("[!] Error when running go mod vendor")
			continue
		}

		os.Rename("vendor", folder_name)
		services.Walk(folder_name, config)
		err = os.RemoveAll(folder_name)
		if err != nil { // handle error
			continue
		}
		log.Println("[+] Deleted vendor folder ")

	}
}

func scanOnRootPath(config models.Optional) {
	log.Println("[+] Golang detected")

	services.Scanning_root_path(config)
	if _, err := os.Stat(folder_name); os.IsNotExist(err) {
		// path/to/whatever does not exist
		log.Println("[+] Run go mod tidy")

		cmdt := exec.Command("go", "mod", "tidy")
		cmdt.Stdout = os.Stdout
		cmdt.Stderr = os.Stderr
		err := cmdt.Run()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("[+] Create Vendor")
		cmd := exec.Command("go", "mod", "vendor")
		var out bytes.Buffer
		cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			log.Fatal(err.Error())
		}

		os.Rename("vendor", folder_name)
	}

	services.Walk(folder_name, config)
	err := os.RemoveAll(folder_name)
	if err != nil { // handle error
		log.Fatal(err)
	}
	log.Println("[+] Deleted vendor folder ")
}
