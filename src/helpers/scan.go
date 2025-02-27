package helpers

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/wahyuhadi/supply-chain/models"
	"github.com/wahyuhadi/supply-chain/services"
)

type URLInfo struct {
	Token string
	URL   string
}

func Scan(config models.Optional) {
	// publicKey, keyError := ssh.NewPublicKeysFromFile("git", "/root/ssh/key", "test123!")
	// if keyError != nil {
	// 	fmt.Println(keyError)
	// }
	if config.Private {
		repo, _ := extractTokenAndURL(config.RepoURI)
		projects_repo := repo.URL
		repos := projects_repo[strings.LastIndex(projects_repo, "/")+1:]
		repos = fmt.Sprintf("/tmp/%s", repos)
		log.Println(fmt.Sprintf("[+] Clone repo %s to folder %s ", projects_repo, repos))
		_, err := git.PlainClone(repos, false, &git.CloneOptions{
			URL: projects_repo,
			Auth: &http.BasicAuth{
				Username: "oauth2",
				Password: repo.Token,
			},
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
	} else {
		projects_repo := config.RepoURI
		repos := projects_repo[strings.LastIndex(projects_repo, "/")+1:]
		repos = fmt.Sprintf("/tmp/%s", repos)
		log.Println(fmt.Sprintf("[+] Clone repo %s to folder %s ", projects_repo, repos, "this is public repos"))
		_, err := git.PlainClone(repos, false, &git.CloneOptions{
			URL:               config.RepoURI,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
			ReferenceName:     plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", config.Branch)),
			Auth: &http.BasicAuth{
				Username: "oauth2",
				Password: "",
			},
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

}

// Struct untuk menyimpan hasil token dan URL

// Function untuk memproses URL dan mengembalikan token dan URL yang bersih dalam bentuk struct
func extractTokenAndURL(inputURL string) (URLInfo, error) {
	// Parsing URL
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return URLInfo{}, err
	}

	// Mendapatkan token (username) yang ada di bagian otentikasi URL
	token := parsedURL.User.Username()

	// Menyusun URL tanpa bagian otentikasi
	parsedURL.User = nil
	urlWithoutAuth := parsedURL.String()

	// Mengembalikan struct URLInfo yang berisi token dan URL bersih
	return URLInfo{
		Token: token,
		URL:   urlWithoutAuth,
	}, nil
}
