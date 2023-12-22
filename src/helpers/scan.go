package helpers

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/wahyuhadi/supply-chain/models"
	"github.com/wahyuhadi/supply-chain/services"
)

func Scan(config models.Optional) {
	publicKey, keyError := ssh.NewPublicKeysFromFile("git", "/root/ssh/key", "test123!")
	if keyError != nil {
		fmt.Println(keyError)
	}
	projects_repo := config.RepoURI
	repos := projects_repo[strings.LastIndex(projects_repo, "/")+1:]
	repos = fmt.Sprintf("/tmp/%s", repos)
	log.Println(fmt.Sprintf("[+] Clone repo %s to folder %s ", projects_repo, repos))
	_, err := git.PlainClone(repos, false, &git.CloneOptions{
		URL:           projects_repo,
		Auth:          publicKey,
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", config.Branch)),
	})

	if err != nil {
		log.Println("[!] Error when clone repo ", projects_repo)
		fmt.Println(err)
		os.Exit(1)
	}

	os.Chdir(repos)
	services.Scanning_root_path(config)
	ExtractResult(config)
}
