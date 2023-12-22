package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/wahyuhadi/supply-chain/models"
)

var (
	URIRule = "https://xsec.thefalcons.site/sast/"
)

func scanning(path string, conf models.Optional) {
	cwd, _ := os.Getwd()
	os.Chdir(path)
	cmd := exec.Command("semgrep", "--config", conf.Rules, conf.Mode)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
	os.Chdir(cwd)
}

func Scanning_root_path(conf models.Optional) {
	cwd, _ := os.Getwd()
	rules := fmt.Sprintf("%s%s", URIRule, conf.Lang)
	log.Println(fmt.Sprintf("[+] Scanning on root path before vendor in dir %s ", cwd))
	cmd := exec.Command("semgrep", "--config", rules, conf.Mode, "--output", conf.FileSaved, "--json")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
